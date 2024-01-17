package repositories

import (
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/entities"
)

var nodeRepository = singleton.NewSingleton[NodeRepository](func() *NodeRepository {
	return &NodeRepository{
		dao.NewRepository[entities.Node](),
	}
}, singleton.Eager)

func GetNodeRepository() *NodeRepository {
	return nodeRepository.Get()
}

type NodeRepository struct {
	*dao.Repository[entities.Node]
}

type NodeConfigRepository struct {
	*dao.Repository[entities.NodeConfig]
}

var nodeConfigRepository = singleton.NewSingleton[NodeConfigRepository](func() *NodeConfigRepository {
	return &NodeConfigRepository{
		dao.NewRepository[entities.NodeConfig](),
	}
}, singleton.Eager)

func GetNodeConfigRepository() *NodeConfigRepository {
	return nodeConfigRepository.Get()
}
