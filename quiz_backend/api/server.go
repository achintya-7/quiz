package api

import (
	"fmt"

	db "github.com/achintya-7/quiz/db/sqlc"
	"github.com/achintya-7/quiz/util"
	"github.com/gin-gonic/gin"
	"github.com/keploy/go-sdk/integrations/kgin/v1"
	"github.com/keploy/go-sdk/keploy"
)

type Server struct {
	store  *db.Store
	config util.Config
	router *gin.Engine
}

func NewServer(store *db.Store, config util.Config, k *keploy.Keploy) (*Server, error) {
	server := &Server{store: store, config: config}

	server.setupRoutes(k)

	return server, nil
}

func (server *Server) setupRoutes(k *keploy.Keploy) {
	router := gin.Default()

	fmt.Println("Setting up keploy")
	kgin.GinV1(k, router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"pingy": "pongy"})
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
