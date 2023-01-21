package impl

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gin-clean-architecture/entity"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/RizkiMufrizal/gin-clean-architecture/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewUserRepositoryImpl(DB *sqlx.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*sqlx.DB
}

func (userRepository *userRepositoryImpl) Create(username string, password string, roles []string) {
	var userRoles []entity.UserRole
	for _, role := range roles {
		userRoles = append(userRoles, entity.UserRole{
			Id:       uuid.New(),
			Username: username,
			Role:     role,
		})
	}
	user := entity.User{
		Username:  username,
		Password:  password,
		IsActive:  true,
		UserRoles: userRoles,
	}

	tx, err := userRepository.DB.Beginx()
	exception.PanicLogging(err)

	_, err = tx.NamedExec("INSERT INTO tb_user (username, password, is_active) VALUES(:username, :password, :is_active)", user)
	_, err = tx.NamedExec("INSERT INTO tb_user_role (user_role_id, role, username) VALUES(:user_role_id, :role, :username)", userRoles)
	if err != nil {
		err := tx.Rollback()
		exception.PanicLogging(err)
	}
	//commit
	err = tx.Commit()
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) DeleteAll() {
	//begin transaction
	tx, err := userRepository.DB.Beginx()
	exception.PanicLogging(err)

	_, err = tx.Exec("DELETE FROM tb_user")
	_, err = tx.Exec("DELETE FROM tb_user_role")
	if err != nil {
		err := tx.Rollback()
		exception.PanicLogging(err)
	}
	//commit
	err = tx.Commit()
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) Authentication(ctx context.Context, username string) (entity.User, error) {
	var userResult entity.User
	err := userRepository.DB.GetContext(ctx, &userResult, ""+
		"SELECT username, password, is_active "+
		"FROM tb_user user "+
		"INNER JOIN tb_user_role user_role ON user_role.username = user.username "+
		"WHERE user.username = $1 and user.is_active = $2", username, true)
	if err != nil {
		return entity.User{}, errors.New("user not found")
	}
	return userResult, nil
}
