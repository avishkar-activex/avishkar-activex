package user

type User struct {
	Id        int64
	Name      string
	Password  string
	Email     string
	AccountId int64
}

func FindByName(name string) (User, error) {
	usr := User{}

	return usr, nil
}
