package repositories

import (
	"github.com/universalmacro/common/dao"
	single "github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/singleton"
)

var GetNodeRepository = single.EagerSingleton(func() *NodeRepository {
	return &NodeRepository{
		dao.NewRepository[entities.Node](singleton.GetDBInstance()),
	}
})

type NodeRepository struct {
	*dao.Repository[entities.Node]
}

var GetNodeConfigRepository = single.EagerSingleton(func() *NodeConfigRepository {
	return &NodeConfigRepository{
		dao.NewRepository[entities.NodeConfig](singleton.GetDBInstance()),
	}
})

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
