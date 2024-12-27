package strings

import "strings"


/* BUILDER */

type Builder struct{
  strings.Builder
}

func (b *Builder) Dump(slice *[]string){
  if len(strings.ReplaceAll(b.String(), "\n", "")) > 0 {
    *slice = append(*slice, b.String())
  } 
  b.Reset()
}
