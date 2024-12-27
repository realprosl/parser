package html

import(
  "strings"
)

func ChangeInterpolingWrapper(src string) string{

  if strings.Contains(src, "{{") && strings.Contains(src, "}}"){
    src = strings.ReplaceAll(src, "{{", "${")
    src = strings.ReplaceAll(src, "}}", "}")
  }
  return src
}
