package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"gitlab.com/DanielStefanK/now-playing-relay/responses"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	logger, _ := zap.NewProduction()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CORS_ORIGIN")},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/api/recent/:username", func(c *gin.Context) {
		username, _ := c.Params.Get("username")

		response, err := http.Get(fmt.Sprintf("https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&format=json&api_key=%s", username, os.Getenv("API_KEY")))

		if err != nil {
			logger.Error(err.Error())

			c.PureJSON(404, responses.ErrorResponse{
				Message: "Could not get utilization for provided studio id",
			})
		}
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		c.Data(200, "application/json", responseData)
	})

	router.Run(":8080")
}
