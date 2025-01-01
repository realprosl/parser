package ts

import (
	"strings"
	"tml-ng/gox/bools"
	str "tml-ng/gox/strings"
	"tml-ng/parser/model"
)

type _Languaje struct{
  typeNodes map[string] model.NodeType
}

func New()*_Languaje{
  return &_Languaje{
    typeNodes : map[string]model.NodeType{
      "Import": &Import{},
      "ClassDeclaration": &ClassDeclaration{},
    },
  }
}

func (l *_Languaje) TypeNode(name string)model.NodeType{
  value, ok := l.typeNodes[name]
  if ok {
    return value
  }
  return nil
}

func (l *_Languaje) Import(source string) model.Import{
  imp := &Import{}
  imp.Parse(source)
  return imp
}

func (l *_Languaje) ClassDeclaration(source string) model.ClassDeclaration{
  cls := &ClassDeclaration{}
  cls.Parse(source)
  return cls
}

func (l *_Languaje) Contain(name, source string)bool{
  return l.TypeNode(name).Contain(source)
}

func (l *_Languaje) SplitInstructions(code string)(instructions []string){
  var insideString bool
  var buffer str.Builder
  context := 0

  for _, char := range code{
    switch true{

      // <};> fuera de string
      case strings.HasSuffix(buffer.String(), "};") && !insideString && context == 0:
        buffer.WriteRune(char)
        buffer.Dump(&instructions)

       // <;> fuera de string
      case strings.HasSuffix(buffer.String(), ";") && !strings.HasSuffix(buffer.String(), "};") && !insideString:
        buffer.WriteRune(char)
        buffer.Dump(&instructions)

      // <}\n> fuera de string y no seguido de <,>
      case strings.HasSuffix(buffer.String(), "}\n") && !insideString && context == 0:
        buffer.WriteRune(char)
        buffer.Dump(&instructions)

      // <\n> fuera de string y no precedido por <,>
      case char == '\n' && !strings.HasSuffix(buffer.String(), ",\n") && !insideString && context == 0:
        buffer.WriteRune(char)
        buffer.Dump(&instructions)

      case char == '"':
        bools.Toggle(&insideString)
        buffer.WriteRune(char)

      case char == '}' && !insideString:
        buffer.WriteRune(char)
        context--

      case char == '{' && !insideString:
        buffer.WriteRune(char)
        context++

      default:
        buffer.WriteRune(char)
    }
  }
  return

}

