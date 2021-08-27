package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20210827200048",
		Up:      mig_20210827200048_welcome_up,
		Down:    mig_20210827200048_welcome_down,
	})
}

func mig_20210827200048_welcome_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE welcome ( name varchar(255) );")
	if err != nil {
		return err
	}
	return nil
}

func mig_20210827200048_welcome_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE welcome")
	if err != nil {
		return err
	}
	return nil
}
