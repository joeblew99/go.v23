package java

import (
	"bytes"
	"log"
	"path"

	"veyron.io/veyron/veyron2/vdl/compile"
	"veyron.io/veyron/veyron2/vdl/vdlutil"
)

const serverWrapperTmpl = `// This file was auto-generated by the veyron vdl tool.
// Source(s):  {{ .Source }}
package {{ .PackagePath }};

public final class {{ .ServiceName }}ServiceWrapper {

    private final {{ .FullServiceName }}Service service;

{{/* Define fields to hold each of the embedded service wrappers*/}}
{{ range $embed := .Embeds }}
    {{/* e.g. private final com.somepackage.gen_impl.ArithStub stubArith; */}}
    private final {{ $embed.WrapperClassName }} {{ $embed.LocalWrapperVarName }};
    {{ end }}

    public {{ .ServiceName }}ServiceWrapper(final {{ .FullServiceName }}Service service) {
        this.service = service;
        {{/* Initialize the embeded service wrappers */}}
        {{ range $embed := .Embeds }}
        this.{{ $embed.LocalWrapperVarName }} = new {{ $embed.WrapperClassName }}(service);
        {{ end }}
    }

    /**
     * Returns all tags associated with the provided method or null if the method isn't implemented
     * by this service.
     */
    public java.lang.Object[] getMethodTags(final com.veyron2.ipc.ServerCall call, final java.lang.String method) throws com.veyron2.ipc.VeyronException {
        {{ range $methodName, $tags := .MethodTags }}
        if ("{{ $methodName }}".equals(method)) {
            return new java.lang.Object[] {
                {{ range $tag := $tags }} {{ $tag }}, {{ end }}
            };
        }
        {{ end }}
        {{ range $embed := .Embeds }}
        try {
            return this.{{ $embed.LocalWrapperVarName }}.getMethodTags(call, method);
        } catch (com.veyron2.ipc.VeyronException e) {}  // method not found.
        {{ end }}
        throw new com.veyron2.ipc.VeyronException("method: " + method + " not found");
    }

     {{/* Iterate over methods defined directly in the body of this service */}}
    {{ range $method := .Methods }}
    public {{ $method.RetType }} {{ $method.Name }}(final com.veyron2.ipc.ServerCall call{{ $method.DeclarationArgs }}) throws com.veyron2.ipc.VeyronException {
        {{ if $method.IsStreaming }}
        final com.veyron2.vdl.Stream<{{ $method.SendType }}, {{ $method.RecvType }}> stream = new com.veyron2.vdl.Stream<{{ $method.SendType }}, {{ $method.RecvType }}>() {
            @Override
            public void send({{ $method.SendType }} item) throws com.veyron2.ipc.VeyronException {
                call.send(item);
            }
            @Override
            public {{ $method.RecvType }} recv() throws java.io.EOFException, com.veyron2.ipc.VeyronException {
                final com.google.common.reflect.TypeToken<?> type = new com.google.common.reflect.TypeToken< {{ $method.RecvType }} >() {
                    private static final long serialVersionUID = 1L;
                };
                final java.lang.Object result = call.recv(type);
                try {
                    return ({{ $method.RecvType }})result;
                } catch (java.lang.ClassCastException e) {
                    throw new com.veyron2.ipc.VeyronException("Unexpected result type: " + result.getClass().getCanonicalName());
                }
            }
        };
        {{ end }} {{/* end if $method.IsStreaming */}}
        {{ if $method.Returns }} return {{ end }} this.service.{{ $method.Name }}( call {{ $method.CallingArgs }} {{ if $method.IsStreaming }} ,stream {{ end }} );
    }
{{end}}

{{/* Iterate over methods from embeded services and generate code to delegate the work */}}
{{ range $eMethod := .EmbedMethods }}
    public {{ $eMethod.RetType }} {{ $eMethod.Name }}(final com.veyron2.ipc.ServerCall call{{ $eMethod.DeclarationArgs }}) throws com.veyron2.ipc.VeyronException {
        {{/* e.g. return this.stubArith.cosine(call, [args], options) */}}
        {{ if $eMethod.Returns }}return{{ end }}  this.{{ $eMethod.LocalWrapperVarName }}.{{ $eMethod.Name }}(call{{ $eMethod.CallingArgs }});
    }
{{ end }} {{/* end range .EmbedMethods */}}

}
`

type serviceWrapperMethod struct {
	CallingArgs     string
	DeclarationArgs string
	IsStreaming     bool
	Name            string
	RecvType        string
	RetType         string
	Returns         bool
	SendType        string
}

type serviceWrapperEmbedMethod struct {
	CallingArgs         string
	DeclarationArgs     string
	LocalWrapperVarName string
	Name                string
	RetType             string
	Returns             bool
}

type serviceWrapperEmbed struct {
	LocalWrapperVarName string
	WrapperClassName    string
}

func processServiceWrapperMethod(iface *compile.Interface, method *compile.Method, env *compile.Env) serviceWrapperMethod {
	return serviceWrapperMethod{
		CallingArgs:     javaCallingArgStr(method.InArgs, true),
		DeclarationArgs: javaDeclarationArgStr(method.InArgs, env, true),
		IsStreaming:     isStreamingMethod(method),
		Name:            vdlutil.ToCamelCase(method.Name),
		RecvType:        javaType(method.OutStream, true, env),
		RetType:         clientInterfaceOutArg(iface, method, true, env),
		Returns:         len(method.OutArgs) >= 2,
		SendType:        javaType(method.InStream, true, env),
	}
}

func processServiceWrapperEmbedMethod(iface *compile.Interface, embedMethod *compile.Method, env *compile.Env) serviceWrapperEmbedMethod {
	return serviceWrapperEmbedMethod{
		CallingArgs:         javaCallingArgStr(embedMethod.InArgs, true),
		DeclarationArgs:     javaDeclarationArgStr(embedMethod.InArgs, env, true),
		LocalWrapperVarName: vdlutil.ToCamelCase(iface.Name) + "Wrapper",
		Name:                vdlutil.ToCamelCase(embedMethod.Name),
		RetType:             clientInterfaceOutArg(iface, embedMethod, true, env),
		Returns:             len(embedMethod.OutArgs) >= 2,
	}
}

// genJavaServiceWrapperFile generates a java file containing a service wrapper for the specified
// interface.
func genJavaServiceWrapperFile(iface *compile.Interface, env *compile.Env) JavaFileInfo {
	embeds := []serviceWrapperEmbed{}
	for _, embed := range allEmbeddedIfaces(iface) {
		embeds = append(embeds, serviceWrapperEmbed{
			WrapperClassName:    javaPath(javaGenPkgPath(path.Join(embed.File.Package.Path, javaGenImplDir, embed.Name+"ServiceWrapper"))),
			LocalWrapperVarName: vdlutil.ToCamelCase(embed.Name) + "Wrapper",
		})
	}
	methodTags := make(map[string][]string)
	// Add generated methods to the tag map:
	methodTags["getMethodTags"] = []string{}
	// Copy method tags off of the interface.
	for _, method := range iface.Methods {
		tags := make([]string, len(method.Tags))
		for i, tagVal := range method.Tags {
			tags[i] = javaConstVal(tagVal, env)
		}
		methodTags[vdlutil.ToCamelCase(method.Name)] = tags
	}
	embedMethods := []serviceWrapperEmbedMethod{}
	for _, embedMao := range dedupedEmbeddedMethodAndOrigins(iface) {
		embedMethods = append(embedMethods, processServiceWrapperEmbedMethod(embedMao.Origin, embedMao.Method, env))
	}
	methods := make([]serviceWrapperMethod, len(iface.Methods))
	for i, method := range iface.Methods {
		methods[i] = processServiceWrapperMethod(iface, method, env)
	}

	data := struct {
		EmbedMethods    []serviceWrapperEmbedMethod
		Embeds          []serviceWrapperEmbed
		FullServiceName string
		Methods         []serviceWrapperMethod
		MethodTags      map[string][]string
		PackagePath     string
		ServiceName     string
		Source          string
	}{
		EmbedMethods:    embedMethods,
		Embeds:          embeds,
		FullServiceName: javaPath(interfaceFullyQualifiedName(iface)),
		Methods:         methods,
		MethodTags:      methodTags,
		PackagePath:     javaPath(path.Join(javaGenPkgPath(iface.File.Package.Path), javaGenImplDir)),
		ServiceName:     iface.Name,
		Source:          iface.File.BaseName,
	}
	var buf bytes.Buffer
	err := parseTmpl("service wrapper", serverWrapperTmpl).Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute service wrapper template: %v", err)
	}
	return JavaFileInfo{
		Dir:  javaGenImplDir,
		Name: iface.Name + "ServiceWrapper.java",
		Data: buf.Bytes(),
	}
}
