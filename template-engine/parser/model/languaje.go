package model

type Language interface{
  HasKeyword()bool
  SplitInstructions(code string)[]string
  ValidateBlockSintax(name, code string)bool
}

