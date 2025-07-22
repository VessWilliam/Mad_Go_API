package domains

type User struct {
	Id       int
	Email    string
	Name     string
	Password string
}

type UserRepo interface {
	Insert(user *User) error
	Get(id int) (*User, error)
}
