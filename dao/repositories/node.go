package repositories

import (
	"github.com/universalmacro/common/dao"
	single "github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/singleton"
)

var nodeRepository = single.SingletonFactory[NodeRepository](func() *NodeRepository {
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

var nodeConfigRepository = single.SingletonFactory[NodeConfigRepository](func() *NodeConfigRepository {
	return &NodeConfigRepository{
		dao.NewRepository[entities.NodeConfig](singleton.GetDBInstance()),
	}
}, single.Eager)

func GetNodeConfigRepository() *NodeConfigRepository {
	return nodeConfigRepository.Get()
}

type NodeConfigRepository struct {
	*dao.Repository[entities.NodeConfig]
}

func (r *NodeConfigRepository) GetByFronendDomain(domain string) (*entities.NodeConfig, error) {
	var nodeConfig entities.NodeConfig
	if err := r.DB.Where("frontend_domain = ?", domain).Find(&nodeConfig).Error; err != nil {
		return nil, err
	}
	return &nodeConfig, nil
}
