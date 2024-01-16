package models

import (
	"time"

	"github.com/universalmacro/common/utils/random"
	"github.com/universalmacro/core/dao/entities"
)

func NewNode(entity *entities.Node) *Node {
	return &Node{entity: entity}
}

type Node struct {
	entity *entities.Node
}

func (n *Node) ID() uint {
	return n.entity.ID
}

func (n *Node) Name() string {
	return n.entity.Name
}

func (n *Node) Description() string {
	return n.entity.Description
}

func (n *Node) CreatedAt() time.Time {
	return n.entity.CreatedAt
}

func (n *Node) UpdatedAt() time.Time {
	return n.entity.UpdatedAt
}

func (n *Node) SecurityKey() string {
	return n.entity.SecurityKey
}

func (n *Node) UpdateSecurityKey() string {
	n.entity.SecurityKey = random.RandomString(64)
	return n.entity.SecurityKey
}

func (n *Node) Entity() *entities.Node {
	return n.entity
}
