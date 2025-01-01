package model


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

func (p *Parser) Parse()*Node{

  program := NewNode(&Program{}, NoPosition()) 

  for _, item := range p.language.SplitInstructions(p.code){

    switch true{

      case p.language.Contain("Import", item): 
        program.AddChildren(NewNode(p.language.Import(item), NoPosition()))

      case p.language.Contain("ClassDeclaration", item):
        program.AddChildren(NewNode(p.language.ClassDeclaration(item), NoPosition()))
    }
  }
  return program
}
