package model

type Program struct{
  Name string
}

func (p *Program) Type()string{
  return "Program"
}
