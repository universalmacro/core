package models

import (
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/utils"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/services/models"
)

func AdminConvertor(admin models.Admin) Admin {
	return Admin{
		ID:          utils.UintToString(admin.ID()),
		Account:     admin.Account(),
		Role:        admin.Role(),
		PhoneNumber: PhoneNumberConvertor(admin.PhoneNumber()),
		CreatedAt:   admin.CreatedAt().Unix(),
		UpdatedAt:   admin.UpdatedAt().Unix(),
	}
}

func PhoneNumberConvertor(phoneNumber *models.PhoneNumber) *PhoneNumber {
	if phoneNumber == nil {
		return nil
	}
	if phoneNumber.CountryCode == "" || phoneNumber.Number == "" {
		return nil
	}
	return &PhoneNumber{
		CountryCode: phoneNumber.CountryCode,
		Number:      phoneNumber.Number,
	}
}

func NodeConvertor(node *models.Node) *Node {
	if node == nil {
		return nil
	}
	return &Node{
		ID:          utils.UintToString(node.ID()),
		Name:        node.Name(),
		SecurityKey: node.SecurityKey(),
		Description: node.Description(),
		CreatedAt:   node.CreatedAt().Unix(),
		UpdatedAt:   node.UpdatedAt().Unix(),
	}
}

func AdminListConvertor(admins dao.List[models.Admin]) dao.List[Admin] {
	var adminList dao.List[Admin]
	var items []Admin = make([]Admin, 0)
	for _, admin := range admins.Items {
		items = append(items, AdminConvertor(admin))
	}
	adminList.Items = items
	adminList.Pagination = admins.Pagination
	return adminList
}

func NodeListConvertor(nodes dao.List[models.Node]) dao.List[Node] {
	var nodeList dao.List[Node]
	var items []Node = make([]Node, 0)
	for _, node := range nodes.Items {
		items = append(items, *NodeConvertor(&node))
	}
	nodeList.Items = items
	nodeList.Pagination = nodes.Pagination
	return nodeList
}

func NodeConfigConvertor(nodeConfig *entities.NodeConfig) *NodeConfig {
	if nodeConfig == nil {
		return nil
	}
	return &NodeConfig{
		Api:      nodeConfig.Api,
		Server:   nodeConfig.Server,
		Database: nodeConfig.Database,
		Redis:    nodeConfig.Redis,
	}
}
