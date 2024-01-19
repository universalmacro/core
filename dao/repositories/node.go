package repositories

import (
	"github.com/universalmacro/common/dao"
	single "github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/singleton"
)

var nodeRepository = single.NewSingleton[NodeRepository](func() *NodeRepository {
	return &NodeRepository{
		dao.NewRepository[entities.Node](singleton.GetDBInstance()),
	}
}, single.Eager)

func GetNodeRepository() *NodeRepository {
	return nodeRepository.Get()
}

type NodeRepository struct {
	*dao.Repository[entities.Node]
}

type NodeConfigRepository struct {
	*dao.Repository[entities.NodeConfig]
}

var nodeConfigRepository = single.NewSingleton[NodeConfigRepository](func() *NodeConfigRepository {
	return &NodeConfigRepository{
		dao.NewRepository[entities.NodeConfig](singleton.GetDBInstance()),
	}
}, single.Eager)

func GetNodeConfigRepository() *NodeConfigRepository {
	return nodeConfigRepository.Get()
}
