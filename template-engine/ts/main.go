package ts

import (
	"strings"
        str "tml-ng/gox/strings"
        "tml-ng/gox/bools"
)

type _Languaje struct{}

func New()*_Languaje{
  return &_Languaje{}
}


func (l *_Languaje)  HasKeyword()bool{
  return true
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
func (l *_Languaje) ValidateBlockSintax(name, code string)bool{
  return true
}

/////////////////////////////////////////////////////////////////7

func isCurrentOpenWrapper(str string)bool{
  return str == "'" || str == "`" || str == `"` || str == "[" || str == "(" || str == "{"
}

func isCurrentCloseWrapper(str string)bool{
  return str == "'" || str == "`" || str == `"` || str == "]" || str == ")" || str == "}"
}


func FindFirstRegion(open string, closed string, source string)(string){
  var match, buffer strings.Builder
  var insideWrapper, insideContext bool

  openWrapper := func(char rune)bool{
    return (strings.HasSuffix(buffer.String(), open) || string(char) == open) && !insideContext && !insideWrapper
  }
  closeWrapper := func(char rune)bool{
    return ( strings.HasSuffix(buffer.String(), closed) || string(char) == closed) && !insideContext && insideWrapper
  }
  openContext := func(char rune)bool{
    return isCurrentOpenWrapper(string(char)) && !insideContext && insideWrapper
  }
  closeContext := func(char rune)bool{
    return isCurrentCloseWrapper(string(char)) && insideContext && insideWrapper
  }

  for _, char := range source{

    switch true{

      case openWrapper(char):
        insideWrapper = true
        buffer.Reset()
        break

      case openContext(char):
        insideContext = true
        match.WriteRune(char)
        buffer.Reset()
        break

      case closeContext(char):
        insideContext = false
        match.WriteRune(char)
        buffer.Reset()
        break

      case closeWrapper(char):
        insideWrapper = false
        buffer.Reset()
        return match.String()

      case insideWrapper:
        match.WriteRune(char)
        break

      default:
        buffer.WriteRune(char)

    }

  }
  return match.String()
}


func FindAllRegions(open, closed, src string)(results []string){
  var match, buffer strings.Builder
  var insideWrapper, insideContext bool

  openWrapper := func(char rune)bool{
    return (strings.HasSuffix(buffer.String(), open) || string(char) == open) && !insideContext && !insideWrapper
  }
  closeWrapper := func(char rune)bool{
    return ( strings.HasSuffix(buffer.String(), closed) || string(char) == closed) && !insideContext && insideWrapper
  }
  openContext := func(char rune)bool{
    return isCurrentOpenWrapper(string(char)) && !insideContext && insideWrapper
  }
  closeContext := func(char rune)bool{
    return isCurrentCloseWrapper(string(char)) && insideContext && insideWrapper
  }

  for _, char := range src{

    switch true{

      case openWrapper(char):
        insideWrapper = true
        buffer.Reset()
        break

      case openContext(char):
        insideContext = true
        match.WriteRune(char)
        buffer.Reset()
        break

      case closeContext(char):
        insideContext = false
        match.WriteRune(char)
        buffer.Reset()
        break

      case closeWrapper(char):
        insideWrapper = false
        buffer.Reset()
        results = append(results, match.String())
        match.Reset()

      case insideWrapper:
        match.WriteRune(char)
        break

      default:
        buffer.WriteRune(char)

    }

  }
  return
}











