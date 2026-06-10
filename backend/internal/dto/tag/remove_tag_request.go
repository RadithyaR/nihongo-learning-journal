package tag

type RemoveTagRequest struct {
	TagID    string `json:"tagId" validate:"required,uuid"`
	ItemType string `json:"itemType" validate:"required"`
	ItemID   string `json:"itemId" validate:"required,uuid"`
}