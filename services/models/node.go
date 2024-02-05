package models

import (
	"fmt"
	"time"

	"github.com/universalmacro/common/snowflake"
	"github.com/universalmacro/common/utils/random"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/dao/repositories"
	"gorm.io/gorm"
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
	if nodeConfig == nil {
		nodeConfig = &entities.NodeConfig{Model: gorm.Model{ID: snowflake.NewIdGenertor(0).Uint()},
			NodeID: n.ID()}
		repositories.GetNodeConfigRepository().Create(nodeConfig)
	}
	return nodeConfig
}

func (n *Node) UpdateConfig(
	frontendDomain *string,
	api *entities.ApiConfig,
	server *entities.ServerConfig,
	database *entities.DBConfig,
	redis *entities.RedisConfig,
	tencentCloud *entities.TencentCloudConfig,
) *entities.NodeConfig {
	nodeConfig := n.Config()
	if frontendDomain != nil {
		nodeConfig.FrontendDomain = frontendDomain
	}
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
	if tencentCloud != nil {
		nodeConfig.TencentCloudConfig = tencentCloud
	}
	repositories.GetNodeConfigRepository().Update(nodeConfig)
	return nodeConfig
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

func (n *Node) CreateMerchant(account, password string) {
	fmt.Println(n.Config().Api)
}

func (n *Node) ListMerchants(index, limit int64) {

}
