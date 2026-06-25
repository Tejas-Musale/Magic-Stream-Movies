package routes

import(
	controller "github.com/Tejas-Musale/Magic-Stream-Movies/Server/MagicStreamMoviesServer/controllers"
	"github.com/Tejas-Musale/Magic-Stream-Movies/Server/MagicStreamMoviesServer/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"

)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {

    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleWare())

    protected.POST("/addmovie", controller.AddMovie(client))
    protected.GET("/movie/:imdb_id", controller.GetMovie(client))
    protected.GET("/recommendedmovies", controller.GetRecommendedMovies(client))
}