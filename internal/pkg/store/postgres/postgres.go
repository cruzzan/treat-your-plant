package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type store struct {
	log *logrus.Logger
	db *sqlx.DB
}

// New creates a postgres backed storage implementing the store.Store interface
func New(pgUrl string, maxConn int, l *logrus.Logger) (*store, error) {
	db, err := sqlx.Open("postgres", pgUrl)

	if err != nil {
		return nil, errors.Wrap(err, "Could not open database connection")
	}

	db.SetMaxOpenConns(maxConn)

	return &store{
		db: db,
		log: l,
	}, nil
}

func (s *store) Ping() error {
	return errors.Wrap(s.db.Ping(), "ping request failed")
}

func (s *store) Grace() {
	err := s.db.Close()
	if err != nil {
		s.log.WithError(err).Error("Could not close the database connection")
	}
}
