package services

import (
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/dao/repositories"
	"github.com/universalmacro/core/services/models"
)

func newNodeService() *NodeService {
	return &NodeService{
		nodeRepository: repositories.GetNodeRepository(),
	}
}

var nodeService = singleton.NewSingleton(newNodeService, singleton.Eager)

func GetNodeService() *NodeService {
	return nodeService.Get()
}

type NodeService struct {
	nodeRepository *repositories.NodeRepository
}

func (s *NodeService) CreateNode(name, description string) *models.Node {
	entity := &entities.Node{Name: name, Description: description}
	node := models.NewNode(entity)
	node.UpdateSecurityKey()
	s.nodeRepository.Create(node.Entity())
	return node
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
