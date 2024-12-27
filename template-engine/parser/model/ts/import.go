package ts

  type ImportDefault struct{
    Dependece []string
    Path string
  }
  func (i *ImportDefault) Type()string{
    return "ImportDefault"
  }



  type ImportAsName struct{
    Name string
    Path string
  }
  func (i *ImportAsName) Type()string{
    return "ImportAsName"
  }



  type ImportEvaluated struct{
    Path string
  }
  func (i *ImportEvaluated) Type()string{
    return "ImportEvaluated"
  }
