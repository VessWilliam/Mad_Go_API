package seed

import (
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func Seeder(dbx *sqlx.DB) error {

	name := "ADMIN"
	email := "Admin@gmail.com"
	password := "root123"

	_, err := dbx.Exec(`insert into roles (name) values ($1) ON CONFLICT DO NOTHING`, name)
	if err != nil {
		log.Fatal("insert role error", err)
		return err
	}

	var count int
	err = dbx.Get(&count, `Select Count(*) from users`)
	if err != nil {
		log.Fatal("count user error", err)
		return err
	}

	if count > 0 {
		log.Println("user already exist. Skipping seedings")
		return err
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	var userId int
	err = dbx.QueryRowx(
		`insert into users (email, name, password) values ($1, $2, $3) returning id`,
		email,
		strings.ToUpper(name),
		string(hash),
	).Scan(&userId)

	if err != nil {
		log.Fatal("insert admin user error", err)
		return err
	}

	var roleId int
	err = dbx.Get(&roleId, `select id from roles where name = $1`, name)
	if err != nil {
		log.Fatal("get role id have error", err)
		return err
	}

	_, err = dbx.Exec(
		`insert into user_roles (user_id, role_id) values ($1, $2)`,
		userId,
		roleId)

	if err != nil {
		log.Fatal("Assign role error", err)
		return err
	}

	log.Println(`User have seeded:`, name, "/", password)
	return nil
}
