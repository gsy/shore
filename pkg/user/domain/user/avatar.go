package user

import (
	"os"
)

type Avatar struct {
	ID      AvatarID
	Path    string
	Url     string
	Content os.File
}

type AvatarID string
