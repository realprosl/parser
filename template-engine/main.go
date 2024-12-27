package main

import (
        "tml-ng/ts"
        "tml-ng/parser/model"
)

var ts_source = `

import DefaultImp from "./html/main.go"
import {imp1, imp2} from "./ts/main.go"
import * as nameImp from "./main.go"
import "./ts/main.go"

const object = {
  name:"Alberto",
  age:43
}

class Input{
  index = "nada"
  number = 3445

  onMount(){
    return this.index
  }
}

class Algo extends Mierda{
  constructor(){}
  static name = "alberto"
}

`

func main(){
  code := model.NewParser(ts_source, ts.New())
  code.Parser()
}

