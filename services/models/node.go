package models

import "github.com/universalmacro/core/dao/entities"

func NewNode(entitiy *entities.Node) *Node {
	return &Node{entitiy: entitiy}
}

type Node struct {
	entitiy *entities.Node
}

func (n *Node) ID() uint {
	return n.entitiy.ID
}
