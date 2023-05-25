package domain

type PagingQuery struct {
	After  *string
	Before *string
	Limit  *int
	Order  *string
}

type CursorTransform struct {
	After  *string `json:"after,omitempty" query:"after"`
	Before *string `json:"before,omitempty" query:"before"`
}
