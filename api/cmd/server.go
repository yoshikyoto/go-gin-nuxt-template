package cmd

import (
	"api-server/config"
	"api-server/route"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Server は gin のサーバーを起動する
func Server() {
	engine := gin.Default()
	env := config.NewEnv()

	// cookie の設定
	// ルーティングの設定より前に処理する必要がある
	secret := env.GetAppSecret()
	store := cookie.NewStore([]byte(secret))
	engine.Use(sessions.Sessions("apply_app_session", store))

	route.SetRoutes(engine)

	port := env.GetServerPort()
	if port == "" {
		port = "3000"
	}
	address := fmt.Sprintf(":%s", port)
	engine.Run(address)
}
