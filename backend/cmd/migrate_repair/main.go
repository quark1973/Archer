package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"archer/internal/config"
)

func main() {
	repair := flag.Bool("repair", false, "clear only a failed first migration with no application tables")
	flag.Parse()

	cfg, err := config.Load("")
	if err != nil {
		fail(err)
	}
	db, err := sql.Open("postgres", cfg.PG.DSN)
	if err != nil {
		fail(err)
	}
	defer db.Close()

	var version int
	var dirty bool
	if err := db.QueryRow(`SELECT version, dirty FROM schema_migrations`).Scan(&version, &dirty); err != nil {
		fail(err)
	}
	var tables int
	if err := db.QueryRow(`SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name <> 'schema_migrations'`).Scan(&tables); err != nil {
		fail(err)
	}
	fmt.Printf("migration version=%d dirty=%t application_tables=%d\n", version, dirty, tables)
	if !*repair {
		return
	}
	if tables != 0 || !((version == 1 && dirty) || (version == 0 && !dirty)) {
		fail(fmt.Errorf("refusing repair: expected failed first migration with no application tables"))
	}
	m, err := migrate.New("file://migrations", cfg.PG.DSN)
	if err != nil {
		fail(err)
	}
	defer m.Close()
	if err := m.Force(database.NilVersion); err != nil {
		fail(err)
	}
	fmt.Println("first migration marked ready to retry")
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
