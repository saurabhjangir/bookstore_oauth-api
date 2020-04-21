package app

import "github.com/saurabhjangir/bookstore_oauth-api/src/http"

func mapURL( handler http.AccessTokenHandler){
	router.GET("/access_token/:access_token_id", handler.GetByID)
}
