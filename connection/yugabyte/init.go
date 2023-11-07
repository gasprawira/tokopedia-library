package yugabyte

import (
	"fmt"
	"log"

	"github.com/yugabyte/gocql"
)

// CreateNewConnection for initialization yugabyte session
func CreateNewConnection(cfg YugabyteConfig) (i InterfaceYugabyte, err error) {
	if len(cfg.ClusterDSN) == 0 {
		err = fmt.Errorf("[Error][Yugabyte] ClusterDSN is nil %s", err)
		return
	}

	session, errInit := InitializeYugabyte(cfg)
	if errInit != nil {
		err = errInit
		log.Println("Error creating yugabyte session", err)
	}

	return &YugabyteInstance{
		Session:  session,
		Keyspace: cfg.Keyspace,
	}, err
}

func InitializeYugabyte(cfg YugabyteConfig) (session *gocql.Session, err error) {
	cluster := gocql.NewCluster(cfg.ClusterDSN...)
	cluster.Port = cfg.Port
	if cfg.Port <= 0 {
		cluster.Port = 9042
	}

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.YugabyteUsername,
		Password: cfg.YugabytePassword,
	}

	cluster.Consistency = ConsistencyMap[cfg.Consistency]
	cluster.Timeout = cfg.Timeout
	cluster.ConnectTimeout = cfg.ConnectTimeout
	cluster.SocketKeepalive = cfg.Keepalive
	if cfg.Consistency == "" {
		cluster.Consistency = Consistency
	}
	if cfg.Timeout == 0 {
		cluster.Timeout = Timeout
	}
	if cfg.ConnectTimeout == 0 {
		cluster.ConnectTimeout = ConnectTimeout
	}
	if cfg.DisableInitialHostLookup {
		cluster.DisableInitialHostLookup = true
	}

	if cfg.LocalDC != "" {
		cluster.PoolConfig = gocql.PoolConfig{
			HostSelectionPolicy: gocql.DCAwareRoundRobinPolicy(cfg.LocalDC),
		}
	}

	if cfg.MaxConnection > 0 {
		cluster.NumConns = cfg.MaxConnection
	}

	session, err = cluster.CreateSession()
	return
}
