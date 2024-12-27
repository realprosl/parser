package model


type Position struct{
  Start int
  End int
}

func NoPosition()Position{
  return Position{}
}
