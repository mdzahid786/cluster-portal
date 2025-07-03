package storage

import "github.com/mdzahid786/cluster-portal/internal/types"

type Storage interface {
	CreateCluster(name string, servers int) (int64, error)
	GetClusterByID(id int64) (types.Cluster, error)
	GetClusters() ([]types.Cluster, error)
	UpdateCluster(id int, servers int) (int, error)
}