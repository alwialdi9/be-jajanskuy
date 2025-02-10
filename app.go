package bejajanskuy

import (
	"fmt"
	"os"

	"github.com/alwialdi9/be-jajanskuy/config"
	"github.com/alwialdi9/be-jajanskuy/internal/routers"
	"github.com/alwialdi9/be-jajanskuy/internal/utils"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func App() {
	os.Setenv("TZ", "Asia/Jakarta")
	godotenv.Load()

	utils.InitLogger()

	config.ConnectDatabase()

	app := routers.Route()

	log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))))
}
