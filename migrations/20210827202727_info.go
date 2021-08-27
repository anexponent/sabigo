package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20210827202727",
		Up:      mig_20210827202727_info_up,
		Down:    mig_20210827202727_info_down,
	})
}

func mig_20210827202727_info_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE info ( name varchar(255) );")
	if err != nil {
		return err
	}
	return nil
}

func mig_20210827202727_info_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE metas")
	if err != nil {
		return err
	}
	return nil
}
