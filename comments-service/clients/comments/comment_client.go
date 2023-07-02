package clients

import (
	model "comments/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCommentById(commentId string) (model.Comment, error) {
	var comment model.Comment

	err := Db.Where("comment_id = ?", commentId).First(&comment).Error

	if err != nil {
		log.Println(err)
		return model.Comment{}, err
	}

	log.Debug("Comment: ", comment)

	return comment, nil
}

func GetCommentsByUserId(userId string) (model.Comments, error) {
	var comments model.Comments

	err := Db.Where("user_id = ?", userId).Find(&comments).Error

	if err != nil {
		log.Println(err)
		return model.Comments{}, err
	}

	log.Debug("Comments: ", comments)

	return comments, nil
}

func GetCommentsByItemId(itemId string) (model.Comments, error) {
	var comments model.Comments

	err := Db.Where("item_id = ?", itemId).Find(&comments).Error

	if err != nil {
		log.Println(err)
		return model.Comments{}, err
	}
	log.Debug("Comments: ", comments)

	return comments, nil
}

func InsertComment(newComment model.Comment) (model.Comment, error) {
	err := Db.Create(&newComment).Error
	if err != nil {
		log.Println(err)
		return model.Comment{}, err
	}

	log.Debug("Comment inserted: ", newComment)

	return newComment, nil
}

func DeleteComment(commentId int) (model.Comment, error) {
	var comment model.Comment

	// Buscar el usuario por su ID
	err := Db.Where("comment_id = ?", commentId).First(&comment).Error
	if err != nil {
		log.Println(err)
		return model.Comment{}, err
	}

	// Eliminar el usuario de la base de datos
	err = Db.Delete(&comment).Error
	if err != nil {
		log.Println(err)
		return model.Comment{}, err
	}

	log.Debug("Comment deleted: ", comment)

	return comment, nil
}
