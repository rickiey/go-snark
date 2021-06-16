package router

import (
	"go-snark/handlers"
	"go-snark/router/middleware"

	"github.com/gin-gonic/gin"
)

// InitRouter ..
func InitRouter() (r *gin.Engine) {
	r = gin.New()
	r.Use(middleware.LoggerHandler, middleware.RecoverHandler)

	r.GET("/ping", handlers.PingPong)
	r.POST("/seal/seal_commit_phase2", handlers.SealCommitPhase2)

	return
}
