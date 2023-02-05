package api

import (
	"database/sql"
	"net/http"

	db "github.com/achintya-7/quiz/db/sqlc"
	"github.com/achintya-7/quiz/models"
	"github.com/gin-gonic/gin"
)

func (server *Server) createQuiz(c *gin.Context) {
	var req models.CreateQuizRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateQuizParams{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
	}
	quiz, err := server.store.CreateQuiz(c, arg)
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}

	c.JSON(200, quiz)
}

func (server *Server) getQuizById(c *gin.Context) {
	var req models.GetQuizByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	quiz, err := server.store.GetQuizById(c, req.ID)
	if err != nil {

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, errorResponse(err))
			return
		}

		c.JSON(500, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, quiz)
}

func (server *Server) getAllQuiz(c *gin.Context) {
	quizzes, err := server.store.ListQuiz(c)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(200, quizzes)
}

func (server *Server) updateQuiz(c *gin.Context) {
	var req models.UpdateQuizRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.UpdateQuizParams{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}
	quiz, err := server.store.UpdateQuiz(c, arg)
	if err != nil {
		c.JSON(500, errorResponse(err))
		return
	}

	c.JSON(200, quiz)
}

func (server *Server) deleteQuiz(c *gin.Context) {
	var req models.DeleteQuizRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteQuiz(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, errorResponse(err))
			return
		}
		
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(200, gin.H{"message": "Quiz deleted successfully"})
}
