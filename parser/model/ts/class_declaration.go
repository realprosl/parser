package ts

import (
	"strings"
)

type ClassDeclaration struct{
  name string
  implements []string
  extends string
  private bool
}

/* ASSETS */

func sliceFirst(src []string)string{
  if len(src) > 0 {
    return src[0]
  }
  return ""
}

func getName(source string)string{

  return sliceFirst(
    strings.Split(
      strings.TrimSpace(
        strings.Replace(
          strings.Replace(
            sliceFirst(strings.Split(source, "\n")),
            "export ", "", 1,
          ),
        "class ", "", 1,
        ),
      ),
      " ",
    ))
}

/* IMPLEMENTS NODETYPE */

 func (c *ClassDeclaration) Type()string{
   return "ClassDeclaration"
 }

 func (c *ClassDeclaration) Contain(code string)bool{
   return ( strings.HasPrefix(code, "export ") || strings.HasPrefix(code, "export class ") ) && ( strings.HasSuffix(code, "}") )
 }

 func (c *ClassDeclaration) Parse(code string){
   c.name = getName(code)
 }

/* IMPLEMENTS CLASSDECLARATION */

func (c *ClassDeclaration) GetName()string{
  return c.name
}

func (c *ClassDeclaration) GetImplements()[]string{
  return c.implements
}

func (c *ClassDeclaration) GetExtends()string{
  return c.extends
}

func (c *ClassDeclaration) IsPrivate()bool{
  return c.private
}


