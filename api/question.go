package api

import (
	"database/sql"
	"net/http"

	db "github.com/achintya-7/quiz/db/sqlc"
	"github.com/achintya-7/quiz/models"
	"github.com/gin-gonic/gin"
)

func (server *Server) createQuestion(c *gin.Context) {
	var req models.CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateQuestionParams{
		QuizID:   req.QuizID,
		Question: req.Question,
		Answer:   req.Answer,
	}
	question, err := server.store.CreateQuestion(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, question)
}

func (server *Server) getQuestionById(c *gin.Context) {
	var req models.GetQuestionByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	question, err := server.store.GetQuestionById(c, req.ID)
	if err != nil {

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, question)
}

func (server *Server) getAllQuestion(c *gin.Context) {
	questions, err := server.store.ListQuestion(c)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (server *Server) getAllQuestionsByQuizID(c *gin.Context) {
	var req models.GetAllQuestionByQuizIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	questions, err := server.store.GetAllQuestionByQuizId(c, req.QuizID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, questions)
}

func (server *Server) updateQuestion(c *gin.Context) {
	var req models.UpdateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdateQuestionParams{
		ID:       req.ID,
		Question: req.Question,
		Answer:   req.Answer,
	}
	question, err := server.store.UpdateQuestion(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, question)
}

func (server *Server) deleteQuestion(c *gin.Context) {
	var req models.DeleteQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteQuestion(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}
