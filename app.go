package bejajanskuy

import (
	"fmt"
	"os"

	"github.com/alwialdi9/be-jajanskuy/config"
	"github.com/alwialdi9/be-jajanskuy/internal/middlewares"
	"github.com/alwialdi9/be-jajanskuy/internal/routers"
	"github.com/alwialdi9/be-jajanskuy/internal/utils"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func App() {
	os.Setenv("TZ", "Asia/Jakarta")
	_ = godotenv.Load()

	utils.InitLogger()

	config.ConnectDatabase()

	app := routers.Route()
	app.Use(cors.New())
	app.Use(middlewares.LoggerMiddleware)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
