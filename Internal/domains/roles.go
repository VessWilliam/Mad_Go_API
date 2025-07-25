package domains

type Roles struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type RolesRepo interface {
}
