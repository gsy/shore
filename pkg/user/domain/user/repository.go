package user

import "context"

type UserRepository interface {
	// Save save user
	Save(ctx context.Context, user *User) (result *User, err error)
}

// AvatarRepository avatar repository
type AvatarRepository interface {
	// Save save avatar
	Save(ctx context.Context, avatar *Avatar) (result *Avatar, err error)
	// Delete delete avatar
	Delete(ctx context.Context, avatar *Avatar) (err error)
}
