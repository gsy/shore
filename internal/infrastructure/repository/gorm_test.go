package repository

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewTestDatabase(t *testing.T) *gorm.DB {
	t.Helper()
	dsn := "store:store.123@tcp(127.0.0.1:3306)/store?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("failed to create database")
	}

	return db
}

type userModel struct {
	gorm.Model
	UserID   string `gorm:"index"`
	Username string
}

func (m userModel) TableName() string {
	return "user"
}

type userEntity struct {
	ID       string
	Username string
}

func (e userEntity) Entity() {}

type conv struct{}

func (c conv) FromEntity(entity *userEntity) (model *userModel) {
	return &userModel{
		UserID:   entity.ID,
		Username: entity.Username,
	}
}

func (c conv) ToEntity(model *userModel) (entity *userEntity) {
	return &userEntity{
		ID:       model.UserID,
		Username: model.Username,
	}
}

func TestCreate(t *testing.T) {
	db := NewTestDatabase(t)

	err := db.AutoMigrate(&userModel{})
	require.NoError(t, err)

	repo := NewRepository[userModel, userEntity](db, conv{})

	newUser := &userEntity{
		ID:       gofakeit.LetterN(5),
		Username: gofakeit.LetterN(5),
	}
	ctx := context.Background()
	_, err = repo.Create(ctx, newUser)
	assert.NoError(t, err)
}
