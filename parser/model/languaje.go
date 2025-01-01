package model

type Language interface{
  SplitInstructions(code string)[]string
  Import(string)Import
  ClassDeclaration(string)ClassDeclaration
  Contain(string,string)bool
}

type NodeType interface{
  Type()string
  Contain(string)bool
  Parse(string)
}

type Import interface{
  NodeType
  GetDependece()[]string
  Dependece(source... string)
  GetPath()string
  Path(string)
  Clone()Import
  Reset()
}

type ClassDeclaration interface{
  NodeType
  GetName()string
  GetImplements()[]string
  GetExtends()string
}

