package app

import (
	"github.com/gin-gonic/gin"
	"github.com/theguptaji/bookstore_oauth-api/src/repository/rest"
	"github.com/theguptaji/bookstore_oauth-api/src/services/access_token"

	"github.com/theguptaji/bookstore_oauth-api/src/http"
	"github.com/theguptaji/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(
		access_token.NewService(rest.NewRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
