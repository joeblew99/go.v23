// Yacc grammar file for the veyron IDL langage.
//
// We use syntax similar to, but not identical to, Go.  For more information
// about the Go language see: http://golang.org/ref/spec
//
// Similar to Go, the formal grammar uses semicolons ';' as terminators, but
// idiomatic usage may omit most semicolons using the following rules:
//   1) During the tokenization phase, semicolons are always auto-inserted at
//      the end of each line after certain tokens.  This is implemented in
//      the lexer via the autoSemi function.
//   2) Semicolons may be omitted before a closing ')' or '}'.  This is
//      implemented via the osemi rule below.
// For details see http://golang.org/ref/spec#Semicolons
//
// The differences between IDL and Go syntax:
//   * IDL specifies APIs via types, interfaces and error ids.  You don't
//     specify implementation.
//   * IDL doesn't consider an interface to be a type; you can't use an
//     interface as the type of a function argument.
//     TODO(toddw): Work out the semantics and change this.
//   * IDL numeric constant definitions are always typed; it would be annoying
//     to support untyped constants across different languages.  There is no
//     iota identifier.
//   * IDL methods may be optionally tagged with a list of const values.
//   * IDL allows error ids to be defined that work across address spaces and
//     languages.
//
// To generate the grammar.go source file containing the parser, run the
// following command.  It'll also generate a grammar.debug file containing a
// list of all states produced for the parser, and some stats.
//   go tool yacc -o grammar.go -v grammar.debug grammar.y && go fmt grammar.go

////////////////////////////////////////////////////////////////////////
// Declarations section.
%{
// This grammar.go file was auto-generated by yacc from grammar.y.

package parse

import (
  "math/big"
  "strings"
)

type strPos struct {
  str string
  pos Pos
}

type intPos struct {
  int *big.Int
  pos Pos
}

type ratPos struct {
  rat *big.Rat
  pos Pos
}

type imagPos struct {
  imag *BigImag
  pos  Pos
}

// typeListToStrList converts a slice of Type to a slice of strPos.  Each type
// must be a TypeNamed with an empty PackageName, otherwise errors are reported,
// and ok=false is returned.
func typeListToStrList(yylex yyLexer, typeList []Type) (strList []strPos, ok bool) {
  ok = true
  for _, t := range typeList {
    var tn *TypeNamed
    if tn, ok = t.(*TypeNamed); !ok {
      lexPosErrorf(yylex, t.Pos(), "%s invalid (expected one or more variable names)", t.String())
      return
    }
    if strings.ContainsRune(tn.Name, '.') {
      ok = false
      lexPosErrorf(yylex, t.Pos(), "%s invalid (expected one or more variable names).", tn.Name)
      return
    }
    strList = append(strList, strPos{tn.Name, tn.P})
  }
  return
}

// ensureNonEmptyToken reports an error if tok is empty.
func ensureNonEmptyToken(yylex yyLexer, tok strPos, errMsg string) {
  if len(tok.str) == 0 {
    lexPosErrorf(yylex, tok.pos, errMsg)
  }
}
%}

// This union is turned into the struct type yySymType.  Most symbols include
// positional information; this is necessary since Go yacc doesn't support
// passing positional information, so we need to track it ourselves.
%union {
  pos        Pos
  strpos     strPos
  intpos     intPos
  ratpos     ratPos
  imagpos    imagPos
  typeexpr   Type
  typeexprs  []Type
  fields     []*Field
  iface      *Interface
  constexpr  ConstExpr
  constexprs []ConstExpr
  complit    *ConstCompositeLit
  kvlit      KVLit
  kvlits     []KVLit
}

// Terminal tokens.  We leave single-char tokens as-is using their ascii code as
// their id, to make the grammar more readable; multi-char tokens get their own
// id.  The start* tokens are dummy tokens to kick off the parse.
%token            startImportsOnly startFullFile
%token <pos>      ';' ':' ',' '.' '(' ')' '[' ']' '{' '}' '<' '>' '='
%token <pos>      '!' '+' '-' '*' '/' '%' '|' '&' '^'
%token <pos>      tOROR tANDAND tLE tGE tNE tEQEQ tLSH tRSH
%token <pos>      tPACKAGE tIMPORT tTYPE tMAP tSTRUCT tINTERFACE tSTREAM
%token <pos>      tCONST tTRUE tFALSE tERRORID
%token <strpos>   tIDENT tSTRLIT
%token <intpos>   tINTLIT
%token <ratpos>   tRATLIT
%token <imagpos>  tIMAGLIT

%type <strpos>     nameref
%type <typeexpr>   type
%type <typeexprs>  type_list streamargs
%type <fields>     field_spec_list field_spec named_arg_list inargs outargs
%type <iface>      iface_item_list iface_item
%type <constexpr>  expr unary_expr operand v_lit
%type <constexprs> tags tag_list
%type <complit>    comp_lit
%type <kvlit>      kv_lit
%type <kvlits>     kv_lit_list

// There are 5 precedence levels for operators, all left-associative, just like
// Go.  Lines are listed in order of increasing precedence.
%left tOROR
%left tANDAND
%left '<' '>' tLE tGE tNE tEQEQ
%left '+' '-' '|' '^'
%left '*' '/' '%' '&' tLSH tRSH

%left notPackage

%start start

%%
////////////////////////////////////////////////////////////////////////
// Rules section.

start:
  startImportsOnly package imports gen_eof
| startFullFile    package imports decls

// Dummy rule to terminate the parse after the imports, regardless of whether
// there are any decls.  Decls always start with either the tTYPE, tCONST or
// tERRORID tokens, and the rule handles all cases - either there's no trailing
// text (the empty case, which would have resulted in EOF anyways), or there's
// one or more decls, where we need to force an EOF.
gen_eof:
  // Empty.
  { lexGenEOF(yylex) }
| tTYPE
  { lexGenEOF(yylex) }
| tCONST
  { lexGenEOF(yylex) }
| tERRORID
  { lexGenEOF(yylex) }

// PACKAGE
package:
  %prec notPackage
  { lexPosErrorf(yylex, Pos{}, "file must start with package statement") }
| tPACKAGE tIDENT ';'
  { lexIDLFile(yylex).PackageDef = NamePos{Name:$2.str, Pos:$2.pos} }

// IMPORTS
imports:
  // Empty.
| imports import ';'

import:
  tIMPORT '(' ')'
| tIMPORT '(' import_spec_list osemi ')'
| tIMPORT import_spec

import_spec_list:
  import_spec
| import_spec_list ';' import_spec

import_spec:
  tSTRLIT
  {
    imps := &lexIDLFile(yylex).Imports
    *imps = append(*imps, &Import{Path:$1.str, NamePos:NamePos{Pos:$1.pos}})
  }
| tIDENT tSTRLIT
  {
    imps := &lexIDLFile(yylex).Imports
    *imps = append(*imps, &Import{Path:$2.str, NamePos:NamePos{Name:$1.str, Pos:$1.pos}})
  }

// DECLARATIONS
decls:
  // Empty.
| decls decl ';'

decl:
  tTYPE '(' ')'
| tTYPE '(' type_spec_list osemi ')'
| tTYPE type_spec
| tTYPE interface_spec
| tCONST '(' ')'
| tCONST '(' const_spec_list osemi ')'
| tCONST const_spec
| tERRORID '(' ')'
| tERRORID '(' errorid_spec_list osemi ')'
| tERRORID errorid_spec

// DATA TYPE DECLARATIONS
type_spec_list:
  type_spec
| type_spec_list ';' type_spec

type_spec:
  tIDENT type
  {
    tds := &lexIDLFile(yylex).TypeDefs
    *tds = append(*tds, &TypeDef{Type:$2, NamePos:NamePos{Name:$1.str, Pos:$1.pos}})
  }

type:
  nameref
  { $$ = &TypeNamed{Name:$1.str, P:$1.pos} }
| '[' tINTLIT ']' type
  { $$ = &TypeArray{Len:int($2.int.Int64()), Elem:$4, P:$1} }
| '[' ']' type
  { $$ = &TypeList{Elem:$3, P:$1} }
| tMAP '[' type ']' type
  { $$ = &TypeMap{Key:$3, Elem:$5, P:$1} }
| tSTRUCT '{' field_spec_list osemi '}'
  { $$ = &TypeStruct{Fields:$3, P:$1} }
| tSTRUCT '{' '}'
  { $$ = &TypeStruct{P:$1} }

field_spec_list:
  field_spec
  { $$ = $1 }
| field_spec_list ';' field_spec
  { $$ = append($1, $3...) }

// The field_spec rule is intended to capture the following patterns:
//    var type
//    var0, var1, var2 type
// where var* refers to a variable name, and type refers to a type.  Each var
// is expressed as an identifier.  An oddity here is that we use a type_list to
// capture the list of variables rather than using a list of IDENTS.  This means
// the grammar accepts invalid constructions, and we must validate afterwards.
//
// We do this to avoid a LALR reduce/reduce conflict with function arguments.
// The problem is exhibited by the in-args of these two functions, where func1
// has three args respectively named A, B, C all of type t1, and func2 has three
// args with name and type t2, t3 and t4 respectively.  The func1 style is
// captured by field_spec in named_arg_list, while the func2 style is captured
// by type_list in args.
//   func1(A, B, C t1)
//   func2(t2, t3, t4)
//
// If we used an ident_list to capture "A, B, C" in func1, but used a type_list
// to capture "t2, t3, t4" in func2, we'd have a reduce/reduce conflict since
// yacc cannot determine whether to reduce as an ident_list or as a type_list;
// we don't know until we've reached token t1 in func1, or token ')' in func2.
//
// The fix can be considered both beautiful and a huge hack.  To avoid the
// conflict we force both forms to use type_list to capture both "A, B, C" and
// "t2, t3, t4".  This avoids the conflict since we're now always reducing via
// type_list, but allows invalid constructions like "[]int, []int []int".  So we
// validate in the action and throw errors.
//
// An alternate fix would have been to remove the IDENT case from the type rule,
// use ident_list to capture both cases, and manually "expand" the grammar to
// distinguish the cases appropriately.  That would ensure we don't allow
// constructions like "int, int int" in the grammar itself, but would lead to a
// much more complicated grammar.  As a bonus, with the type_list solution we
// can give better error messages.
field_spec:
  type_list type
  {
    if names, ok := typeListToStrList(yylex, $1); ok {
      for _, n := range names {
        $$ = append($$, &Field{Type:$2, NamePos:NamePos{Name:n.str, Pos:n.pos}})
      }
    } else {
      lexPosErrorf(yylex, $2.Pos(), "perhaps you forgot a comma before %q?.", $2.String())
    }
  }

// INTERFACE DECLARATIONS
interface_spec:
  tIDENT tINTERFACE '{' '}'
  {
    ifs := &lexIDLFile(yylex).Interfaces
    *ifs = append(*ifs, &Interface{NamePos:NamePos{Name:$1.str, Pos:$1.pos}})
  }
| tIDENT tINTERFACE '{' iface_item_list osemi '}'
  {
    $4.Name, $4.Pos = $1.str, $1.pos
    ifs := &lexIDLFile(yylex).Interfaces
    *ifs = append(*ifs, $4)
  }

iface_item_list:
  iface_item
  { $$ = $1 }
| iface_item_list ';' iface_item
  {
    $1.Embeds = append($1.Embeds, $3.Embeds...)
    $1.Methods = append($1.Methods, $3.Methods...)
    $$ = $1
  }

iface_item:
  tIDENT inargs streamargs outargs tags
  { $$ = &Interface{Methods: []*Method{{InArgs:$2, InStream:$3[0], OutStream:$3[1], OutArgs:$4, Tags:$5, NamePos:NamePos{Name:$1.str, Pos:$1.pos}}}} }
| nameref
  { $$ = &Interface{Embeds: []*NamePos{{Name:$1.str, Pos:$1.pos}}} }

inargs:
  '(' ')'
  { $$ = nil }
| '(' named_arg_list ocomma ')'
  { $$ = $2 }
| '(' type_list ocomma ')'
  // Just like Go, we allow a list of types without variable names.  See the
  // field_spec rule for a workaround to avoid a reduce/reduce conflict.
  {
    for _, t := range $2 {
      $$ = append($$, &Field{Type:t, NamePos:NamePos{Pos:t.Pos()}})
    }
  }

// The named_arg_list rule is just like the field_spec_list, but uses comma ','
// as a delimiter rather than semicolon ';'.
named_arg_list:
  field_spec
  { $$ = $1 }
| named_arg_list ',' field_spec
  { $$ = append($1, $3...) }

type_list:
  type
  { $$ = []Type{$1} }
| type_list ',' type
  { $$ = append($1, $3) }

// The outargs accept everything regular inargs accept, and are also allowed to
// be empty or contain a single non-parenthesized type without an arg name.
outargs:
  // Empty.
  { $$ = nil }
| type
  { $$ = []*Field{{Type:$1, NamePos:NamePos{Pos:$1.Pos()}}} }
| inargs
  { $$ = $1 }

streamargs:
  // Empty.
  { $$ = []Type{nil, nil} }
| tSTREAM '<' '>'
  { $$ = []Type{nil, nil} }
| tSTREAM '<' type '>'
  { $$ = []Type{$3, nil} }
| tSTREAM '<' type ',' type '>'
  { $$ = []Type{$3, $5} }

tags:
  // Empty.
  { $$ = nil }
| '{' '}'
  { $$ = nil }
| '{' tag_list ocomma '}'
  { $$ = $2 }

tag_list:
  expr
  { $$ = []ConstExpr{$1} }
| tag_list ',' expr
  { $$ = append($1, $3) }

// CONST DEFINITIONS
const_spec_list:
  const_spec
| const_spec_list ';' const_spec

const_spec:
  tIDENT '=' expr
  {
    cds := &lexIDLFile(yylex).ConstDefs
    *cds = append(*cds, &ConstDef{Expr:$3, NamePos:NamePos{Name:$1.str, Pos:$1.pos}})
  }

expr:
  unary_expr
  { $$ = $1 }
| expr tOROR expr
  { $$ = &ConstBinaryOp{"||", $1, $3, $2} }
| expr tANDAND expr
  { $$ = &ConstBinaryOp{"&&", $1, $3, $2} }
| expr '<' expr
  { $$ = &ConstBinaryOp{"<", $1, $3, $2} }
| expr '>' expr
  { $$ = &ConstBinaryOp{">", $1, $3, $2} }
| expr tLE expr
  { $$ = &ConstBinaryOp{"<=", $1, $3, $2} }
| expr tGE expr
  { $$ = &ConstBinaryOp{">=", $1, $3, $2} }
| expr tNE expr
  { $$ = &ConstBinaryOp{"!=", $1, $3, $2} }
| expr tEQEQ expr
  { $$ = &ConstBinaryOp{"==", $1, $3, $2} }
| expr '+' expr
  { $$ = &ConstBinaryOp{"+", $1, $3, $2} }
| expr '-' expr
  { $$ = &ConstBinaryOp{"-", $1, $3, $2} }
| expr '*' expr
  { $$ = &ConstBinaryOp{"*", $1, $3, $2} }
| expr '/' expr
  { $$ = &ConstBinaryOp{"/", $1, $3, $2} }
| expr '%' expr
  { $$ = &ConstBinaryOp{"%", $1, $3, $2} }
| expr '|' expr
  { $$ = &ConstBinaryOp{"|", $1, $3, $2} }
| expr '&' expr
  { $$ = &ConstBinaryOp{"&", $1, $3, $2} }
| expr '^' expr
  { $$ = &ConstBinaryOp{"^", $1, $3, $2} }
| expr tLSH expr
  { $$ = &ConstBinaryOp{"<<", $1, $3, $2} }
| expr tRSH expr
  { $$ = &ConstBinaryOp{">>", $1, $3, $2} }

unary_expr:
  operand
  { $$ = $1 }
| '!' unary_expr
  { $$ = &ConstUnaryOp{"!", $2, $1} }
| '+' unary_expr
  { $$ = &ConstUnaryOp{"+", $2, $1} }
| '-' unary_expr
  { $$ = &ConstUnaryOp{"-", $2, $1} }
| '^' unary_expr
  { $$ = &ConstUnaryOp{"^", $2, $1} }
| type '(' expr ')'
  { $$ = &ConstTypeConv{$1, $3, $1.Pos()} }
// TODO(bprosnitz) Add .real() and .imag() for complex.

operand:
  tTRUE
  { $$ = &ConstLit{true, $1} }
| tFALSE
  { $$ = &ConstLit{false, $1} }
| tSTRLIT
  { $$ = &ConstLit{$1.str, $1.pos} }
| tINTLIT
  { $$ = &ConstLit{$1.int, $1.pos} }
| tRATLIT
  { $$ = &ConstLit{$1.rat, $1.pos} }
| tIMAGLIT
  { $$ = &ConstLit{$1.imag, $1.pos} }
| nameref
  { $$ = &ConstNamed{$1.str, $1.pos} }
| type comp_lit
  { $$ = &ConstCompositeLit{$1, $2.KVList, $1.Pos()} }
| '(' expr ')'
  { $$ = $2 }

comp_lit:
  '{' '}'
  { $$ = &ConstCompositeLit{nil, nil, $1} }
| '{' kv_lit_list ocomma '}'
  { $$ = &ConstCompositeLit{nil, $2, $1} }

kv_lit_list:
  kv_lit
  { $$ = []KVLit{$1} }
| kv_lit_list ',' kv_lit
  { $$ = append($1, $3) }

kv_lit:
  v_lit
  { $$ = KVLit{Value:$1} }
| expr ':' v_lit
  { $$ = KVLit{Key:$1, Value:$3} }

v_lit:
  comp_lit
  { $$ = $1 }
| expr
  { $$ = $1 }

// ERROR IDS
errorid_spec_list:
  errorid_spec
| errorid_spec_list ';' errorid_spec

errorid_spec:
  tIDENT
  {
    eds := &lexIDLFile(yylex).ErrorIDs
    *eds = append(*eds, &ErrorID{NamePos:NamePos{Name:$1.str, Pos:$1.pos}})
  }
| tIDENT '=' tSTRLIT
  {
    ensureNonEmptyToken(yylex, $3, "error id must be non-empty if specified")
    eds := &lexIDLFile(yylex).ErrorIDs
    *eds = append(*eds, &ErrorID{ID:$3.str, NamePos:NamePos{Name:$1.str, Pos:$1.pos}})
  }

// MISC TOKENS

// nameref describes a named reference to another type, interface or const.
nameref:
  tIDENT
  { $$ = $1 }
| tIDENT '.' tIDENT
  { $$ = strPos{$1.str+"."+$3.str, $1.pos} }

// Optional semicolon
osemi:
  // Empty.
| ';'

// Optional comma
ocomma:
  // Empty.
| ','
