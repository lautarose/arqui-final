package clients

import (
	dto "user/dtos/user"
	model "user/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) (model.User, error) {
	var user model.User

	err := Db.Where("user_id = ?", id).First(&user).Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	log.Debug("User: ", user)

	return user, nil
}

func GetUserByUsername(username string) (model.User, error) {
	var user model.User

	err := Db.Where("user_name = ?", username).First(&user).Error

	if err != nil {
		log.Println(err)
		return user, err
	}
	log.Debug("User: ", user)

	return user, nil
}

func GetUserByEmail(email string) (model.User, error) {
	var user model.User

	err := Db.Where("user_email = ?", email).First(&user).Error

	if err != nil {
		log.Println(err)
		return user, err
	}
	log.Debug("User: ", user)

	return user, nil
}

func GetUsers() (model.Users, error) {

	var users model.Users

	err := Db.Find(&users).Error

	if err != nil {
		log.Println(err)
		return users, err
	}

	log.Debug("Users: ", users)

	return users, nil
}

func InsertUser(newUser model.User) error {
	err := Db.Create(&newUser).Error
	if err != nil {
		log.Println(err)
		return err
	}

	log.Debug("User inserted: ", newUser)

	return nil
}

func UpdateUser(id int, updatedUser dto.UserUpdateDto) (model.User, error) {
	var user model.User

	// Buscar el usuario por su ID
	err := Db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		log.Println(err)
		return user, err
	}

	// Actualizar los campos del usuario

	user.Name = updatedUser.Name
	user.LastName = updatedUser.LastName
	user.Pwd = updatedUser.Password

	// Guardar los cambios en la base de datos
	err = Db.Save(&user).Error
	if err != nil {
		log.Println(err)
		return user, err
	}

	log.Debug("User updated: ", user)

	return user, nil
}

func DeleteUser(id int) (model.User, error) {
	var user model.User

	// Buscar el usuario por su ID
	err := Db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		log.Println(err)
		return user, err
	}

	// Eliminar el usuario de la base de datos
	err = Db.Delete(&user).Error
	if err != nil {
		log.Println(err)
		return user, err
	}

	log.Debug("User deleted: ", user)

	return user, nil
}
