package database_test

import (
	"mangosteen/internal/database"
	"testing"
)

func BenchmarkCrud(b *testing.B) {
	database.PgConnect()
	database.PgCreateTables()
	database.Migrate()
	defer database.PgClose()
	for i := 0; i < b.N; i++ {
		database.Crud()
	}
}
