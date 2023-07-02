package dto

type CommentInsertDto struct {
	ItemID  string `json:"item_id"`
	Message string `json:"message"`
}
