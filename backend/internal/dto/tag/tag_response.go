package tag

import "time"

type TagResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Color     *string    `json:"color"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}