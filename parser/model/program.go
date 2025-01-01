package model

type Program struct{
  Name string
}

func (p *Program) Type()string{
  return "Program"
}

func (p *Program) Contain(source string)bool{
  return false
}

func (p *Program) Parse(source string){}
