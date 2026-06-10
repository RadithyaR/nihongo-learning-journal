package tag

type CreateTagRequest struct {
	Name  string  `json:"name" validate:"required,max=100"`
	Color *string `json:"color"`
}