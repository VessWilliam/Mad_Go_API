package domains

type Role struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type RolesRepo interface {
	Insert(role *Role) error
	Update(role *Role) error
	GetAll() ([]*Role, error)
	GetById(id int) (*Role, error)
	DeleteById(id int) error
	GetRoleByEmail(email string) ([]string, error)
}
