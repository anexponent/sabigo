package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20210827201457",
		Up:      mig_20210827201457_metas_up,
		Down:    mig_20210827201457_metas_down,
	})
}

func mig_20210827201457_metas_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE metas ( name varchar(255) );")
	if err != nil {
		return err
	}
	return nil
}

func mig_20210827201457_metas_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE metas")
	if err != nil {
		return err
	}
	return nil
}
