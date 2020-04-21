package app

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhjangir/bookstore_oauth-api/src/datasource"
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/access_token"
	"github.com/saurabhjangir/bookstore_oauth-api/src/http"
)

var (
	router = gin.Default()
)
func StartApplication() {
	atHandler := http.NewAccessTokenHandler(access_token.NewAccessTokenService(datasource.NewRepository()))
	mapURL(atHandler)
	router.Run(":3301")
}

