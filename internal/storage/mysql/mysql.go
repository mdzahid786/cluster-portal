package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mdzahid786/cluster-portal/internal/config"
	"github.com/mdzahid786/cluster-portal/internal/types"
)

type Mysql struct {
	Db *sql.DB
}

func New(conf *config.Config)(*Mysql, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.Username, conf.Password, conf.DbHost, conf.Dbname))
	if err != nil {
		log.Fatal("Count not conneted to the database ", err.Error())
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS clusters(
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(200),
		servers INT
	);`)
	if err!=nil {
		return nil, err
	}
	return &Mysql{
		Db: db,
	}, nil
}

func(m *Mysql) CreateCluster(name string, servers int) (int64, error) {
	stm, err := m.Db.Prepare("INSERT INTO clusters(name, servers) VALUES(?, ?)")
	if err!=nil {
		return 0,err
	}
	defer stm.Close()
	result, err :=stm.Exec(name, servers)
	if err!=nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err!=nil {
		return 0, err
	}
	return lastId, nil
}

func(m *Mysql) GetClusterByID(id int64)(types.Cluster, error) {
	stmt, err := m.Db.Prepare("SELECT id, name, servers FROM clusters WHERE id=? LIMIT 0,1")
	if err!=nil {
		return types.Cluster{},err
	}
	defer stmt.Close()
	var cluster types.Cluster
	err = stmt.QueryRow(id).Scan(&cluster.Id, &cluster.Name, &cluster.Servers)
	if err!=nil {
		return types.Cluster{},err
	}
	return cluster,nil
}

func(m *Mysql) GetClusters() ([]types.Cluster, error) {
	slog.Info("Preparing query")
	stmt, err := m.Db.Prepare("SELECT id, name, servers FROM clusters")
	if err!=nil {
		return []types.Cluster{}, nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err!=nil {
		return []types.Cluster{}, err
	}
	
	defer rows.Close()
	var clusters []types.Cluster
	for rows.Next() {
		var cluster types.Cluster
		err = rows.Scan(&cluster.Id, &cluster.Name, &cluster.Servers)
		if err!=nil {
			return []types.Cluster{}, nil
		}
		clusters = append(clusters, cluster)
	}
	return clusters, nil
}

func(m *Mysql)  UpdateCluster(id int, servers int) (int, error){
	// Check if cluster exists
	stmt, err := m.Db.Prepare("SELECT id FROM clusters WHERE id=? LIMIT 0,1")
	if err!=nil {
		return id, err
	}
	defer stmt.Close()
	
	var clusterId int
	err = stmt.QueryRow(id).Scan(&clusterId)
	if err != nil {
		return 0, err // cluster not found
	}

	fmt.Println("Updating cluster")
	stm, err := m.Db.Prepare("UPDATE clusters SET servers = ? WHERE id = ?")
	if err!=nil {
		return 0,nil
	}
	defer stm.Close()

	_, err = stm.Exec(servers, id)
	if err != nil {
		return 0, err
	}

	return id, nil
}