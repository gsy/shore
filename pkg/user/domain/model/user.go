package user_model

type User struct {
	ID   string
	Name string
}

func (u User) Entity() {}
