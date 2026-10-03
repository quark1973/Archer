package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"archer/internal/config"
)

func main() {
	cfg, err := config.Load("")
	if err != nil {
		fail(err)
	}
	db, err := sql.Open("postgres", cfg.PG.DSN)
	if err != nil {
		fail(err)
	}
	defer db.Close()

	for _, m := range []struct {
		provider, name, alias, kind string
	}{
		{"deepseek", "deepseek-chat", "deepseek-chat", "chat"},
		{"ollama", "bge-m3", "bge-m3", "embedding"},
	} {
		result, err := db.Exec(`
			INSERT INTO model_configs (provider_id, model_name, alias, type, is_default)
			SELECT p.id, $2::VARCHAR(64), $3::VARCHAR(64), $4::VARCHAR(16), TRUE FROM model_providers AS p
			WHERE p.code = $1 AND NOT EXISTS (
				SELECT 1 FROM model_configs WHERE alias = $3::VARCHAR(64) AND type = $4::VARCHAR(16)
			)`, m.provider, m.name, m.alias, m.kind)
		if err != nil {
			fail(err)
		}
		rows, _ := result.RowsAffected()
		fmt.Printf("%s: inserted %d default configuration(s)\n", m.alias, rows)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
