package domains

type Role struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type RolesRepo interface {
	Insert(role *Role) error
	GetAll() ([]*Role, error)
	GetById(id string) (*Role, error)
	DeleteById(id string) error
}
