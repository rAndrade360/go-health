package migration

import (
	"database/sql"
	"log"
)

func CreateUserTable(db *sql.DB) {
	q := `CREATE TABLE user (
		id varchar(255) not null,
		name varchar(255),
		username varchar(255),
		email varchar(255),
		status varchar(100),
		weight DOUBLE,
		birthdate varchar(255),
		createdAt varchar(255),
		updatedAt varchar(255),
		PRIMARY KEY (id)
	);
	`

	res, err := db.Exec(q)
	if err != nil {
		log.Fatalf("Error to exec query: %s", err.Error())
	}

	log.Println(res.RowsAffected())
}
