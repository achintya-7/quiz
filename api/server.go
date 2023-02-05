package api

import (
	db "github.com/achintya-7/quiz/db/sqlc"
	"github.com/achintya-7/quiz/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	config util.Config
	router *gin.Engine
}

func NewServer(store *db.Store, config util.Config) (*Server, error) {
	server := &Server{store: store, config: config}

	server.setupRoutes()

	return server, nil
}

func (server *Server) setupRoutes() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})

	// * Quiz Routes
	quiz := router.Group("/quiz")
	quiz.GET("/", server.getAllQuiz)
	quiz.GET("/:id", server.getQuizById)
	quiz.POST("/", server.createQuiz)
	quiz.PUT("/", server.updateQuiz)
	quiz.DELETE("/", server.deleteQuiz)

	// * Question Routes
	question := router.Group("/question")
	question.GET("/", server.getAllQuestion)
	question.GET("/:id", server.getQuestionById)
	question.POST("/", server.createQuestion)
	question.PUT("/", server.updateQuestion)
	question.DELETE("/", server.deleteQuestion)
	question.GET("/quiz/:quiz_id", server.getAllQuestionsByQuizID)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
