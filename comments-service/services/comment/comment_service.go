package services

import (
	commentClient "comments/clients/comments"
	commentDto "comments/dtos/comment"
	commentModel "comments/models"
	jwtUtils "comments/utils/jwt"
	"errors"
	"strconv"
)

type commentService struct{}

type commentServiceInterface interface {
	GetComments(itemID string) (commentDto.CommentsDto, error)
	InsertComment(auth string, comment commentDto.CommentInsertDto) (commentDto.CommentDto, error)
	DeleteComment(auth string, commentId string) (commentDto.CommentDto, error)
}

var (
	CommentService commentServiceInterface
)

func init() {
	CommentService = &commentService{}
}

func (s *commentService) InsertComment(userId string, newComment commentDto.CommentInsertDto) (commentDto.CommentDto, error) {

	var newCommentModel commentModel.Comment
	newCommentModel.UserID = userId
	newCommentModel.ItemID = newComment.ItemID
	newCommentModel.Message = newComment.Message

	commentInserted, err := commentClient.InsertComment(newCommentModel)

	if err != nil {
		return commentDto.CommentDto{}, err
	}

	var commentInsertedDto commentDto.CommentDto

	commentInsertedDto.CommentID = commentInserted.CommentID
	commentInsertedDto.ItemID = commentInserted.ItemID
	commentInsertedDto.UserID = commentInserted.UserID
	commentInsertedDto.Message = commentInserted.Message
	commentInsertedDto.CreatedAt = commentInserted.CreatedAt

	return commentInsertedDto, nil
}

func (s *commentService) GetComments(itemID string) (commentDto.CommentsDto, error) {
	var commentsDto commentDto.CommentsDto

	comments, err := commentClient.GetCommentsByItemId(itemID)

	if err != nil {
		return commentDto.CommentsDto{}, err
	}

	for _, comment := range comments {
		var commentDto commentDto.CommentDto

		commentDto.CommentID = comment.CommentID
		commentDto.ItemID = comment.ItemID
		commentDto.UserID = comment.UserID
		commentDto.Message = comment.Message
		commentDto.CreatedAt = comment.CreatedAt

		commentsDto = append(commentsDto, commentDto)
	}

	return commentsDto, nil
}

func (s *commentService) DeleteComment(authToken string, commentId string) (commentDto.CommentDto, error) {
	var commentReturn commentDto.CommentDto

	// Verificar el token de autenticación
	claims, err := jwtUtils.VerifyToken(authToken)
	if err != nil {
		return commentDto.CommentDto{}, err
	}

	// Obtener el ID del usuario del token
	userId, err := strconv.Atoi(claims.Id)
	if err != nil {
		return commentDto.CommentDto{}, err
	}

	compareId := strconv.Itoa(userId)

	commentToDelete, err := commentClient.GetCommentById(commentId)

	if err != nil {
		return commentDto.CommentDto{}, err
	}

	if commentToDelete.UserID != compareId {
		err = errors.New("no tiene autorizacion para borrar el comentario")
		return commentDto.CommentDto{}, err
	}

	commentToDeleteId, err := strconv.Atoi(commentId)

	if err != nil {
		return commentDto.CommentDto{}, err
	}

	// Eliminar el comentario
	comment, err := commentClient.DeleteComment(commentToDeleteId)
	if err != nil {
		return commentDto.CommentDto{}, err
	}

	// Asignar los detalles del usuario eliminado al objeto de respuesta
	commentReturn.CommentID = comment.CommentID
	commentReturn.ItemID = comment.ItemID
	commentReturn.UserID = comment.UserID
	commentReturn.Message = comment.Message
	commentReturn.CreatedAt = comment.CreatedAt

	return commentReturn, nil
}
