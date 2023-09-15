package todolistdto

type CreateToDoListRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
}

type UpdateToDoListRequest struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}
