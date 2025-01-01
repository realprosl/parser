package main

import (
	"fmt"
	"tml-ng/parser/model"
	"tml-ng/parser/model/ts"
)

var ts_source = `

import DefaultImp from "./html/main.go"
import {alberto, paco} from "./casa/madre/h.go"
import {react} from "/tuputacasa.vom"

export class Input implements Mount{
  index = "nada"
  number = 3445

  onMount(){
    return this.index
  }
}

`

func main(){
  ast := model.NewParser(ts_source, ts.New()).Parse()

  ast.FindImports(func(node model.Import){
    fmt.Println(node.GetDependece())
  })

  ast.FinsClassDeclaration(func(node model.ClassDeclaration){
    fmt.Println(node.GetName())
  })

}

