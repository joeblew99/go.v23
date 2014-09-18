package java

import (
	"bytes"
	"fmt"
	"log"
	"path"

	"veyron.io/veyron/veyron2/vdl/compile"
	"veyron.io/veyron/veyron2/vdl/vdlutil"
)

const clientInterfaceTmpl = `// This file was auto-generated by the veyron vdl tool.
// Source: {{ .Source }}
package {{ .PackagePath }};

{{ .ServiceDoc }}
public interface {{ .ServiceName }} {{ .Extends }} {
{{ range $method := .Methods }}
    {{/* If this method has multiple return arguments, generate the class. */}}
    {{ if $method.IsMultipleRet }}
    public static class {{ $method.UppercaseMethodName }}Out {
        {{ range $retArg := $method.RetArgs }}
        public {{ $retArg.Type }} {{ $retArg.Name }};
        {{ end }}
    }
    {{ end }}

    {{/* Generate the method signature. */}}
    {{ $method.Doc }}
    public {{ $method.RetType }} {{ $method.Name }}(final com.veyron2.ipc.Context context{{ $method.Args }}) throws com.veyron2.ipc.VeyronException;
    public {{ $method.RetType }} {{ $method.Name }}(final com.veyron2.ipc.Context context{{ $method.Args }}, final com.veyron2.Options veyronOpts) throws com.veyron2.ipc.VeyronException;
{{ end }}
}
`

type clientInterfaceArg struct {
	Type string
	Name string
}

type clientInterfaceMethod struct {
	Name                string
	Doc                 string
	RetType             string
	Args                string
	UppercaseMethodName string
	RetArgs             []clientInterfaceArg
	IsMultipleRet       bool
}

func clientInterfaceNonStreamingOutArg(iface *compile.Interface, method *compile.Method, useClass bool, env *compile.Env) string {
	switch len(method.OutArgs) {
	case 0:
		panic("Unexpected to have 0 out args in an interface method")
	case 1:
		// "void" or "Void"
		return javaType(nil, useClass, env)
	case 2:
		return javaType(method.OutArgs[0].Type, useClass, env)
	default:
		return javaPath(path.Join(interfaceFullyQualifiedName(iface), method.Name+"Out"))
	}
}

func clientInterfaceOutArg(iface *compile.Interface, method *compile.Method, isService bool, env *compile.Env) string {
	if isStreamingMethod(method) && !isService {
		return fmt.Sprintf("com.veyron2.vdl.ClientStream<%s,%s, %s>", javaType(method.InStream, true, env), javaType(method.OutStream, true, env), clientInterfaceNonStreamingOutArg(iface, method, true, env))
	}
	return clientInterfaceNonStreamingOutArg(iface, method, false, env)
}

func processClientInterfaceMethod(iface *compile.Interface, method *compile.Method, env *compile.Env) clientInterfaceMethod {
	retArgs := make([]clientInterfaceArg, len(method.OutArgs)-1)
	// Include all return args except for error.
	for i := 0; i < len(method.OutArgs)-1; i++ {
		retArgs[i].Name = vdlutil.ToCamelCase(method.OutArgs[i].Name)
		retArgs[i].Type = javaType(method.OutArgs[i].Type, false, env)
	}
	return clientInterfaceMethod{
		Name:                vdlutil.ToCamelCase(method.Name),
		Doc:                 method.Doc,
		RetType:             clientInterfaceOutArg(iface, method, false, env),
		UppercaseMethodName: method.Name,
		Args:                javaDeclarationArgStr(method.InArgs, env, true),
		RetArgs:             retArgs,
		IsMultipleRet:       len(retArgs) > 1,
	}
}

// genJavaClientInterfaceFile generates the Java interface file for the provided
// interface.
func genJavaClientInterfaceFile(iface *compile.Interface, env *compile.Env) JavaFileInfo {
	methods := make([]clientInterfaceMethod, len(iface.Methods))
	for i, method := range iface.Methods {
		methods[i] = processClientInterfaceMethod(iface, method, env)
	}
	data := struct {
		Extends     string
		Methods     []clientInterfaceMethod
		PackagePath string
		ServiceDoc  string
		ServiceName string
		Source      string
	}{
		Extends:     javaExtendsStr(iface.Embeds, ""),
		Methods:     methods,
		PackagePath: javaPath(javaGenPkgPath(iface.File.Package.Path)),
		ServiceDoc:  javaDoc(iface.Doc),
		ServiceName: iface.Name,
		Source:      iface.File.BaseName,
	}
	var buf bytes.Buffer
	err := parseTmpl("client interface", clientInterfaceTmpl).Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute struct template: %v", err)
	}
	return JavaFileInfo{
		Name: iface.Name + ".java",
		Data: buf.Bytes(),
	}
}
