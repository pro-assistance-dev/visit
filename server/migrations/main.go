package migrations

import "github.com/uptrace/bun/migrate"

func Init() []*migrate.Migrations {
	res := make([]*migrate.Migrations, 0)
	migrations := migrate.NewMigrations()
	if err := migrations.DiscoverCaller(); err != nil {
		panic(err)
	}
	res = append(res, migrations)
	return res
}
