package models

type CreateQuizRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	CreatedBy   string `json:"created_by" binding:"required"`
}

type GetQuizByIdRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type UpdateQuizRequest struct {
	Id          int64  `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type DeleteQuizRequest struct {
	ID int64 `json:"id" binding:"required"`
}
