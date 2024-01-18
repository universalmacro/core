package models

import (
	"time"

	"github.com/universalmacro/common/utils/random"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/dao/repositories"
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

func (n *Node) Config() *entities.NodeConfig {
	nodeConfig, _ := repositories.GetNodeConfigRepository().FindOne("node_id = ?", n.ID())
	return nodeConfig
}

func (n *Node) UpdateConfig(
	api *entities.ApiConfig,
	server *entities.ServerConfig,
	database *entities.DBConfig,
	redis *entities.RedisConfig,
) *entities.NodeConfig {
	nodeConfigRepository := repositories.GetNodeConfigRepository()
	nodeConfig, _ := nodeConfigRepository.FindOne("node_id = ?", n.ID())
	if api != nil {
		nodeConfig.Api = api
	}
	if server != nil {
		nodeConfig.Server = server
	}
	if database != nil {
		nodeConfig.Database = database
	}
	if redis != nil {
		nodeConfig.Redis = redis
	}
	nodeConfigRepository.Update(nodeConfig)
	return nodeConfig
}

func (n *Node) GetDatabaseConfig() *entities.DBConfig {
	nodeConfigRepository := repositories.GetNodeConfigRepository()
	nodeConfig, _ := nodeConfigRepository.FindOne("node_id = ?", n.ID())
	return nodeConfig.Database
}

func (n *Node) UpdateDatabaseConfig(dbConfig *entities.DBConfig) *entities.DBConfig {
	nodeConfigRepository := repositories.GetNodeConfigRepository()
	nodeConfig, _ := nodeConfigRepository.FindOne("node_id = ?", n.ID())
	nodeConfig.Database = dbConfig
	nodeConfigRepository.Update(nodeConfig)
	return nodeConfig.Database
}

func (n *Node) GetRedisConfig() *entities.RedisConfig {
	nodeConfig, _ := repositories.GetNodeConfigRepository().FindOne("node_id = ?", n.ID())
	return nodeConfig.Redis
}

func (n *Node) UpdateRedisConfig(redisConfig *entities.RedisConfig) *entities.RedisConfig {
	nodeConfigRepository := repositories.GetNodeConfigRepository()
	nodeConfig, _ := nodeConfigRepository.FindOne("node_id = ?", n.ID())
	nodeConfig.Redis = redisConfig
	nodeConfigRepository.Update(nodeConfig)
	return nodeConfig.Redis
}
