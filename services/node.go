package services

import (
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/dao/repositories"
	"github.com/universalmacro/core/services/models"
)

func newNodeService() *NodeService {
	return &NodeService{
		adminRepository: repositories.GetNodeRepository(),
	}
}

var nodeService = singleton.NewSingleton(newNodeService, singleton.Eager)

func GetNodeService() *NodeService {
	return nodeService.Get()
}

type NodeService struct {
	adminRepository *repositories.NodeRepository
}

func (s *NodeService) CreateNode(name, description string) *models.Node {
	entity := &entities.Node{Name: name, Description: description}
	s.adminRepository.Create(entity)
	node := models.NewNode(entity)
	return node
}
