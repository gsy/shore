package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	user_model "github.com/gsy/store/pkg/user/domain/model"
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

func TestNewRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "create repository",
			args: args{
				db: NewTestDatabase(t),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRepository(tt.args.db)
			assert.NotNil(t, got)

		})
	}
}

func Test_repoImpl_FindByID(t *testing.T) {
	db := NewTestDatabase(t)
	repo := NewRepository(db)
	type args struct {
		ctx context.Context
		id  uint64
	}
	tests := []struct {
		name       string
		args       args
		fixture    func(args args) *user_model.User
		wantEntity *user_model.User
		wantErr    bool
	}{
		{
			name: "found",
			args: args{
				ctx: context.Background(),
				id:  gofakeit.Uint64(),
			},
			fixture: func(args args) *user_model.User {
				expected := &user_model.User{
					ID:   gofakeit.LetterN(5),
					Name: gofakeit.Username(),
				}
				expected, err := repo.general.Create(args.ctx, expected)
				require.NoError(t, err)
				return expected
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantEntity = tt.fixture(tt.args)
			gotEntity, err := repo.FindByID(tt.args.ctx, tt.wantEntity.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repoImpl.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEntity, tt.wantEntity) {
				t.Errorf("repoImpl.FindByID() = %v, want %v", gotEntity, tt.wantEntity)
			}
		})
	}
}
