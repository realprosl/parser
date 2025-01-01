package ts

import (
	"strings"
	"tml-ng/parser/model"
)

  type Import struct{
    dependece []string
    path string
  }

  // constructor
  func NewImport()*Import{
    return &Import{}
  }

  // implement Import
  func (i *Import) GetDependece()[]string{
    return i.dependece
  }

  // return struct cloned
  func (i *Import) Clone()model.Import{
    return &Import{path:i.path, dependece:i.dependece}
  }

  // clear any properties in struct
  func (i *Import) Reset(){
    i.path = ""
    i.dependece = []string{}
  }

  // getter path
  func (i *Import) GetPath()string{
    return i.path
  }

  // setter dependece
  func (i *Import) Dependece(source... string){
    i.dependece = source
  }

  // setter path
  func (i *Import) Path(path string){
    i.path = path
  }

  // implement NodeType
  func (i *Import) Type()string{
    return "Import"
  }

  // validate if is import instruction in source string
  func (i *Import) Contain(instruction string)bool{
    return strings.HasPrefix(instruction, "import ")
  }

  // parse code string in ast custom
  func (i *Import) Parse(code string){

    // dividir instruccion en path y Dependece
    div := strings.Split(code, " from ")
    if len(div) > 1{
      dependece, path := strings.Replace(div[0],"import ","", 1), div[1]

      i.path = strings.TrimSpace(path)

      if strings.Contains(dependece, "{") && strings.Contains(dependece,"}"){
        i.dependece = strings.Split(strings.Replace(strings.Replace(dependece, "{","",1),"}","",1),",")
      }else{
        i.dependece = append(i.dependece, strings.TrimSpace(dependece))
      }

    }
  }
