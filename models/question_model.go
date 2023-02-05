package models

type CreateQuestionRequest struct {	
	QuizID int64 `json:"quiz_id" binding:"required"`
	Question string `json:"question" binding:"required"`
	Answer string `json:"answer" binding:"required"`
}

type GetQuestionByIdRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type UpdateQuestionRequest struct {
	ID int64 `json:"id" binding:"required"`
	Question string `json:"question" binding:"required"`
	Answer string `json:"answer" binding:"required"`
}

type DeleteQuestionRequest struct {
	ID int64 `json:"id" binding:"required"`
}

type GetAllQuestionByQuizIDRequest struct {
	QuizID int64 `uri:"quiz_id" binding:"required"`
}