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
