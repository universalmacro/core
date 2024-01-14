package models

import (
	"time"

	"github.com/universalmacro/core/dao/entities"
)

func NewNode(entitiy *entities.Node) *Node {
	return &Node{entitiy: entitiy}
}

type Node struct {
	entitiy *entities.Node
}

func (n *Node) ID() uint {
	return n.entitiy.ID
}

func (n *Node) Name() string {
	return n.entitiy.Name
}

func (n *Node) Description() string {
	return n.entitiy.Description
}

func (n *Node) CreatedAt() time.Time {
	return n.entitiy.CreatedAt
}

func (n *Node) UpdatedAt() time.Time {
	return n.entitiy.UpdatedAt
}
