package scylladbctx

import (
	"fmt"
	"pokemon-lab-api/internal/server/infrastructure/config"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
	"go.uber.org/zap"
)

func NewSession(cfg *config.Config, logger *zap.Logger) (*gocqlx.Session, error) {
	// Create gocql cluster.
	cluster := gocql.NewCluster(cfg.ScyllaHosts...)
	cluster.ProtoVersion = 4
	cluster.Consistency = gocql.Quorum
	// Wrap session on creation, gocqlx session embeds gocql.Session pointer.
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		logger.Error("unable to connect to scylla", zap.Error(err))
		return nil, err
	}

	query := fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};`, cfg.ScyllaKeySpaces)

	errKey := session.ExecStmt(query)

	if errKey != nil {
		logger.Error("unable to create key to scylla", zap.Error(errKey))
		return nil, errKey
	}

	session.Close()

	// TODO: fine a way to replace this reconection patter to set the keyspace.
	cluster.Keyspace = cfg.ScyllaKeySpaces
	session, err = gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		logger.Error("unable to connect to scylla", zap.Error(err))
		return nil, err
	}

	return &session, nil
}
