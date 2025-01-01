package model

type Node struct{
  position Position
  children []*Node
  typeNode NodeType
}

func NewNode(typeNode NodeType, position Position )*Node{
  return &Node{
    position:position,
    typeNode:typeNode,
  }
}

func (n *Node) AddChildren(children *Node)*Node{
  n.children = append(n.children, children)
  return n
}

func (n *Node) Find(query string, fn func(*Node)){
  if n.Type() == query { fn(n) }else{
    if len(n.children) > 0 {
      for _, child := range n.children{
        child.Find(query, fn)
      }
    }
  }
}

func (n *Node) FindImports(fn func(Import)){
  n.Find("Import", func(node *Node){ fn(node.Unwrap().(Import)) })
}

func (n *Node) FinsClassDeclaration(fn func(ClassDeclaration)){
  n.Find("ClassDeclaration", func(node *Node){ fn(node.Unwrap().(ClassDeclaration)) })
}


func (n *Node) Unwrap() interface{}{
  return n.typeNode
}

func (n *Node) Type() string{
  return n.typeNode.Type()
}

func (n *Node) Position()Position{
  return n.position
}


