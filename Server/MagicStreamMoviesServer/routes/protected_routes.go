package routes

import(
	controller "github.com/Tejas-Musale/Magic-Stream-Movies/Server/MagicStreamMoviesServer/controllers"
	"github.com/Tejas-Musale/Magic-Stream-Movies/Server/MagicStreamMoviesServer/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"

)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleWare())

	router.POST("/addmovie", controller.AddMovie(client))
	router.GET("/movie/:imdb_id", controller.GetMovie(client))
	router.GET("/recommendedmovies", controller.GetRecommendedMovies(client))
}