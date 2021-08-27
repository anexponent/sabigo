package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20210827200539",
		Up:      mig_20210827200539_users_up,
		Down:    mig_20210827200539_users_down,
	})
}

func mig_20210827200539_users_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE users ( name varchar(255) );")
	if err != nil {
		return err
	}
	return nil
}

func mig_20210827200539_users_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users")
	if err != nil {
		return err
	}
	return nil
}