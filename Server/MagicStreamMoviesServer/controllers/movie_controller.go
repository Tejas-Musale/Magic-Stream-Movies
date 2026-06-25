package controllers

import (
	"context"
	"log"
	"os"
	"strconv"

	// "strings"
	"errors"
	"fmt"
	"net/http"

	// "strings"
	"time"

	"github.com/Tejas-Musale/Magic-Stream-Movies/Server/MagicStreamMoviesServer/database"
	"github.com/Tejas-Musale/Magic-Stream-Movies/Server/MagicStreamMoviesServer/models"
	"github.com/Tejas-Musale/Magic-Stream-Movies/Server/MagicStreamMoviesServer/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// var rankingCollection *mongo.Collection = database.OpenCollection("rankings")
var validate = validator.New()

func GetMovies(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var movieCollection *mongo.Collection = database.OpenCollection("movies", client)

		var movies []models.Movie

		cursor, err := movieCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
			return
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies"})
			return
		}

		c.JSON(http.StatusOK, movies)
	}
}

func GetMovie(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		movieID := c.Param("imdb_id")
		// fmt.Println("Received request for movie with imdb_id:", movieID)

		if movieID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "imdb_id is required"})
			return
		}

		var movieCollection *mongo.Collection = database.OpenCollection("movies", client)

		var movie models.Movie

		err := movieCollection.FindOne(ctx, bson.M{"imdb_id": movieID}).Decode(&movie)

		// fmt.Println("Error:", err)
		// fmt.Printf("Movie found: %+v\n", movie)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movie"})
			return
		}

		c.JSON(http.StatusOK, movie)

	}
}

func AddMovie(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var movie models.Movie

		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if err := validate.Struct(movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error, validation failed": err.Error()})
			return
		}

		var movieCollection *mongo.Collection = database.OpenCollection("movies", client)

		result, err := movieCollection.InsertOne(ctx, movie)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add movie"})
			return
		}

		c.JSON(http.StatusCreated, result)

	}
}

// func GetRecommededMovies(client *mongo.Client) gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		userId, err := utils.GetUserIdFromContext(c)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user ID"})
// 			return
// 		}

// 		ctx, cancel := context.WithTimeout(c, 100*time.Second)
// 		defer cancel()

// 		favourite_genres, err := GetUsersFavouriteGenres(userId, client, c)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user's favourite genres"})
// 			return
// 		}

// 		err = godotenv.Load(".env")
// 		if err != nil {
// 			log.Println("Warning: Could not load .env file, proceeding with environment variables")
// 		}
// 		var recommendedMoviesLimitVal int64 = 5

// 		recommendedMoviesLimitStr := os.Getenv("RECOMMENDED_MOVIES_LIMIT")

// 		if recommendedMoviesLimitStr != "" {
// 			recommendedMoviesLimitVal, _ = strconv.ParseInt(recommendedMoviesLimitStr, 10, 64)
// 		}

// 		findOptions := options.Find()

// 		findOptions.SetSort(bson.D{{Key: "ranking.ranking_value", Value: 1}})

// 		findOptions.SetLimit(recommendedMoviesLimitVal)

// 		filter := bson.M{"genre.genre_name": bson.M{"$in": favourite_genres}}

// 		var movieCollection *mongo.Collection = database.OpenCollection("movies", client)

// 		cursor, err := movieCollection.Find(ctx, filter, findOptions)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recommended movies"})
// 			return
// 		}
// 		defer cursor.Close(ctx)

// 		var recommendedMovies []models.Movie
// 		if err := cursor.All(ctx, &recommendedMovies); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode recommended movies"})
// 			return
// 		}
//         c.JSON(http.StatusOK, recommendedMovies)

// 	}
// }

func GetRecommendedMovies(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("=== RECOMMENDED MOVIES ===")

		value, exists := c.Get("userId")
		fmt.Println("userId exists:", exists)
		fmt.Println("userId value:", value)

		userId, err := utils.GetUserIdFromContext(c)

		if err != nil {
			fmt.Println("GetUserIdFromContext error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "User Id not found in context"})
		// 	return
		// }

		favourite_genres, err := GetUsersFavouriteGenres(userId, client, c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = godotenv.Load(".env")
		if err != nil {
			log.Println("Warning: .env file not found")
		}
		var recommendedMovieLimitVal int64 = 5

		recommendedMovieLimitStr := os.Getenv("RECOMMENDED_MOVIE_LIMIT")

		if recommendedMovieLimitStr != "" {
			recommendedMovieLimitVal, _ = strconv.ParseInt(recommendedMovieLimitStr, 10, 64)
		}

		findOptions := options.Find()

		findOptions.SetSort(bson.D{{Key: "ranking.ranking_value", Value: 1}})

		findOptions.SetLimit(recommendedMovieLimitVal)

		filter := bson.D{
			{Key: "genre.genre_name", Value: bson.D{
				{Key: "$in", Value: favourite_genres},
			}},
		}

		var ctx, cancel = context.WithTimeout(c, 100*time.Second)
		defer cancel()

		var movieCollection *mongo.Collection = database.OpenCollection("movies", client)

		cursor, err := movieCollection.Find(ctx, filter, findOptions)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching recommended movies"})
			return
		}
		defer cursor.Close(ctx)

		var recommendedMovies []models.Movie

		if err := cursor.All(ctx, &recommendedMovies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, recommendedMovies)
	}
}

func GetUsersFavouriteGenres(userId string, client *mongo.Client, c *gin.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(c, 100*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userId}

	projection := bson.M{"favourite_genres.genre_name": 1, "_id": 0}

	opts := options.FindOne().SetProjection(projection)
	var result bson.M

	var usersCollection *mongo.Collection = database.OpenCollection("users", client)

	err := usersCollection.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []string{}, nil
		}
		return nil, err
	}

	favGenresArray, ok := result["favourite_genres"].(bson.A)
	if !ok {
		return []string{}, errors.New("Failed to parse favourite genres")
	}

	var genreNames []string
	for _, item := range favGenresArray {
		if genreMap, ok := item.(bson.D); ok {
			for _, elem := range genreMap {
				if elem.Key == "genre_name" {
					if name, ok := elem.Value.(string); ok {
						genreNames = append(genreNames, name)
					}
				}
			}
		}
	}

	return genreNames, nil

}

func GetGenre(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var genres []models.Genre

		var genreCollection *mongo.Collection = database.OpenCollection("genres", client)

		cursor, err := genreCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching movie genres"})
			return
		}
		defer cursor.Close(ctx)

		if err := cursor.All(ctx, &genres); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, genres)

	}
}

// func AdminReviewUpdate() gin.HandlerFunc {
// 	return func(c *gin.Context){
// 		movieID := c.Param("imdb_id")
// 		if movieID == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "imdb_id is required"})
// 			return
// 		}

// 		var req struct{
// 			AdminReview string `json:"admin_review" validate:"required"`
// 		}
// 		var resp struct {
// 			RankingName string `json:"ranking_name"`
// 			AdminReview string `json:"admin_review"`
// 		}

// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 			return
// 		}

// 	}
// }

// func GetReviewRanking(admin_review string) (string, int, error) {
// 	rankings, err := GetRankings()
// 	if err != nil {
// 		return "", 0, err
// 	}

// 	sentimentDelimited := ""

// 	for _, ranking := range rankings {
// 		if ranking.RankingValue != 999{
// 			sentimentDelimited += ranking.RankingName + ","
// 		}
// 	}

// 	sentimentDelimited = strings.Trim(sentimentDelimited, ",")

// 	err = godotenv.Load(".env")

// 	if err != nil {
// 		log.Println("Warning: Could not load .env file, proceeding with environment variables")
// 	}

// }

// func GetRankings() ([]models.Ranking, error) {
// 	var rankings []models.Ranking

// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	cursor, err := rankingCollection.Find(ctx, bson.M{})

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	if err = cursor.All(ctx, &rankings); err != nil {
// 		return nil, err
// 	}

// 	return rankings, nil
// }
