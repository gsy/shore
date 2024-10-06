package fake


type UserRepository struct {
	users []user_model.User
}

func (r *UserRepository) Create(user user_model.User) (err error) {
	r.users = append(r.users, user)
	return nil
}
