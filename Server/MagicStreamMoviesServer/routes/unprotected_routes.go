package routes

import (
	controller "github.com/Tejas-Musale/Magic-Stream-Movies/Server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupUnProtectedRoutes(router *gin.Engine, client *mongo.Client) {

	router.POST("/register", controller.RegisterUser(client))
	router.POST("/login", controller.LoginUser(client))
	router.GET("/movies", controller.GetMovies(client))
	router.POST("/logout", controller.LogoutHandler(client))
	router.GET("/genres", controller.GetGenre(client))
	router.POST("/refresh", controller.RefreshTokenHandler(client))
}
