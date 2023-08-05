package postgres

import "APG6/pkg/db/postgres_db"

type Repository struct {
	admin   *postgres_db.PostgresDb
	visitor *postgres_db.PostgresDb
}

func New(admin *postgres_db.PostgresDb, visitor *postgres_db.PostgresDb) *Repository {
	return &Repository{
		admin:   admin,
		visitor: visitor,
	}
}
