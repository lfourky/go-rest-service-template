package integration

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/lfourky/go-rest-service-template/pkg/repository/postgres"
	"github.com/stretchr/testify/require"
)

func SeedDatabase(config postgres.Config, t *testing.T, sqlSeed ...string) {
	require := require.New(t)
	db, teardown := OpenDB(config, t)
	defer teardown()

	for _, seed := range sqlSeed {
		queries := strings.Split(seed, ";")
		// Remove last element, which is just an empty string.
		queries = queries[:len(queries)-1]
		for _, query := range queries {
			_, err := db.Exec(query)
			require.NoError(err, "unable to execute query: %s, got error: %s", query, err)
		}
	}
}

func OpenDB(config postgres.Config, t *testing.T) (*sql.DB, func()) {
	require := require.New(t)

	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DatabaseName,
	))
	require.NoError(err, "unable to connect to database: %s", err)

	err = db.Ping()
	require.NoError(err, "unable to ping database: %s", err)

	return db, func() {
		// Close the database connection, since the application uses a separate connection, anyway.
		err = db.Close()
		require.NoError(err, "unable to close database: %s", err)
	}
}

// This function should be called at the beginning of every test,
// in order to truncate tables and enable the test to start from an empty database.
func ClearDatabase(config postgres.Config, t *testing.T) {
	require := require.New(t)
	// Important: maintain this list with existing tables to truncate.
	tables := []string{
		"item", "user",
	}

	statements := make([]string, len(tables))

	db, teardown := OpenDB(config, t)
	defer teardown()

	// We need to disable foreign key checks, in order to truncate tables in any order.
	for i := range tables {
		statements[i] = `TRUNCATE TABLE "` + tables[i] + `" CASCADE`
	}

	for _, statement := range statements {
		_, err := db.Exec(statement)
		require.NoError(err, "unable to execute statement: %s, got error: %s", statement, err)
	}
}
