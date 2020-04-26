package app

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhjangir/bookstore_oauth-api/src/clients/cassandra"
	"github.com/saurabhjangir/bookstore_oauth-api/src/datasource"
	"github.com/saurabhjangir/bookstore_oauth-api/src/domain/access_token"
	"github.com/saurabhjangir/bookstore_oauth-api/src/http"
	l "github.com/saurabhjangir/bookstore_oauth-api/src/utils/loggers"
)

var (
	router = gin.Default()
)
func StartApplication() {
	l.Log.Info("Initializing oauth application ...")
	session , err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	session.Close()
	atHandler := http.NewAccessTokenHandler(access_token.NewService(datasource.NewRepository(), datasource.NewUserRepository()))
	mapURL(atHandler)
	router.Run(":3301")
}

