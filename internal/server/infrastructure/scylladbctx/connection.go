package scylladbctx

import (
	"pokemon-lab-api/internal/server/infrastructure/config"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
	"go.uber.org/zap"
)

func NewSession(cfg *config.Config, logger *zap.Logger) (*gocqlx.Session, error) {
	// Create gocql cluster.
	cluster := gocql.NewCluster(cfg.ScyllaHost...)
	// Wrap session on creation, gocqlx session embeds gocql.Session pointer.
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		logger.Fatal("unable to connect to scylla", zap.Error(err))
	}

	return &session, nil
}
