package main

import (
	"civi-backend-challenge/handlers"
	"civi-backend-challenge/models"
	"civi-backend-challenge/routers"
	"civi-backend-challenge/services"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	port = os.Getenv("PORT")
)

func init() {
	fmt.Println(fmt.Sprintf("initializing service on port:%s", port))
}

func main() {
	endpoint := fmt.Sprintf(":%s", port)

	pointsFile, err := openFilePoints()
	if err != nil {
		log.Fatal(err)
	}

	pointsService := services.NewPointService(pointsFile)
	handlers := handlers.New(pointsService)

	initRouters := routers.InitRouter(handlers)

	server := &http.Server{
		Addr:    endpoint,
		Handler: initRouters,
	}

	_ = server.ListenAndServe()
}

func openFilePoints() ([]*models.Points, error) {
	filePath := os.Getenv("FILEPATH")
	fileName := os.Getenv("FILENAME")
	var points []*models.Points

	pointsFile, err := os.Open(filePath + fileName)
	if err != nil {
		log.Printf("Error to open file with points values. The application can't initialize")
		return nil, err
	}

	defer pointsFile.Close()

	value, _ := ioutil.ReadAll(pointsFile)
	json.Unmarshal(value, &points)

	return points, nil
}
