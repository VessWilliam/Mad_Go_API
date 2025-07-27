package domains

type User struct {
	Id       int    `db:"id"`
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Role     []Role `db:"roles"`
}

type UserRepo interface {
	Insert(user *User) error
	GetAll() ([]*User, error)
	GetById(id string) (*User, error)
	GetRolesByUserId(userId string) ([]Role, error)
}
