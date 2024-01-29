package services

import (
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/common/snowflake"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/dao/repositories"
	"github.com/universalmacro/core/services/models"
	"gorm.io/gorm"
)

func newNodeService() *NodeService {
	return &NodeService{
		nodeRepository:       repositories.GetNodeRepository(),
		nodeConfigRepository: repositories.GetNodeConfigRepository(),
	}
}

var nodeService = singleton.NewSingleton(newNodeService, singleton.Eager)

func GetNodeService() *NodeService {
	return nodeService.Get()
}

type NodeService struct {
	nodeRepository       *repositories.NodeRepository
	nodeConfigRepository *repositories.NodeConfigRepository
}

func (s *NodeService) CreateNode(name, description string) *models.Node {
	var idGenerator = snowflake.NewIdGenertor(1)
	entity := &entities.Node{Name: name, Description: description}
	node := models.NewNode(entity)
	node.UpdateSecurityKey()
	s.nodeRepository.Create(node.Entity())
	repositories.GetNodeConfigRepository().Create(
		&entities.NodeConfig{Model: gorm.Model{ID: idGenerator.Uint()},
			NodeID: node.ID()})
	return node
}

func (s *NodeService) GetNode(id uint) *models.Node {
	entity, _ := s.nodeRepository.GetById(id)
	if entity == nil {
		return nil
	}
	return models.NewNode(entity)
}

func (s *NodeService) GetNodeByFrontendDomain(domain string) *models.Node {
	config, _ := s.nodeConfigRepository.GetByFronendDomain(domain)
	if config == nil {
		return nil
	}
	return s.GetNode(config.NodeID)
}

func (s *NodeService) ListNode(index, limit int64) dao.List[models.Node] {
	if limit == 0 {
		limit = 1
	}
	nodeList, _ := s.nodeRepository.Pagination(index, limit)
	var nodes []models.Node
	for index := range nodeList.Items {
		nodes = append(nodes, *models.NewNode(&nodeList.Items[index]))
	}
	return dao.List[models.Node]{Items: nodes, Pagination: nodeList.Pagination}
}

func (s *NodeService) DeleteNode(id uint) {
	node, _ := s.nodeRepository.GetById(id)
	nodeConfig, _ := s.nodeConfigRepository.FindOne("node_id = ?", id)
	s.nodeConfigRepository.Delete(nodeConfig)
	s.nodeRepository.Delete(node)
}
