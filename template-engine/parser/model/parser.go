package model

import "fmt"

type Parser struct{
  code string
  language Language
}

func NewParser(code string, lang Language)*Parser{
  return &Parser{
    code: code,
    language : lang,
  }
}

func (p *Parser) Parser(){
  for index, item := range p.language.SplitInstructions(p.code){
    fmt.Println("<<",index,">>")
    fmt.Println(item)
  }
}
