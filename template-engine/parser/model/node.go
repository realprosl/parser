package model

type NodeType interface{
  Type()string
}


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

func (n *Node) AddChildren(children... *Node)*Node{
  n.children = append(n.children, children...)
  return n
}

// TODO: por hacer
func (n *Node) Find(query string, fn func(*Node)){}

// TODO: por hacer
func (n *Node) FindAll(query string , fn func(*Node)){}

func (n *Node) Unwrap() interface{}{
  return n.typeNode
}

func (n *Node) Type() string{
  return n.typeNode.Type()
}

func (n *Node) Position()Position{
  return n.position
}

func (n *Node) End()int{
  return n.position.End
}

func (n *Node) Start()int{
  return n.position.Start
}


