package app

import "github.com/saurabhjangir/bookstore_oauth-api/src/http"

func mapURL( handler http.IaccessTokenHandler){
	router.GET("/oauth/access_token/:access_token", handler.GetByID)
	router.POST("/oauth/access_token", handler.Create)
	router.POST("/oauth/access_token/", handler.Create)
}
