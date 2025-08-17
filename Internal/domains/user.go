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
	Update(user *User) error
	GetAll() ([]*User, error)
	GetById(id int) (*User, error)
	GetRolesByUserId(userId int) ([]Role, error)
	AssignRolesToRoles(userId int, roleIds []int) error
	GetByEmail(email string) (*User, error)
}
