package util

import (
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

func MustParseUUID(s string) uuid.UUID {
	u, err := uuid.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

func MustParseUUIDs(s []string) []uuid.UUID {
	u := make([]uuid.UUID, 0, len(s))
	for _, v := range s {
		u = append(u, MustParseUUID(v))
	}
	return u
}

func ToPostgresUUIDs(s []uuid.UUID) []postgres.Expression {
	u := make([]postgres.Expression, 0, len(s))
	for _, v := range s {
		u = append(u, postgres.UUID(v))
	}
	return u
}
