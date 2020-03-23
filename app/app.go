package app

import (
	"context"
	"github.com/govindamurali/COVID-selfreport-go-api/app/handler"
	"github.com/govindamurali/COVID-selfreport-go-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// App has router and db instances
type App struct {
	Router      *mux.Router
	mongoClient *mongo.Client
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(config.MongoConfig.ConnectionUrl)
	a.mongoClient, _ = mongo.Connect(ctx, clientOptions)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the reports
	a.Get("/reports", a.handleRequest(handler.GetAllReports))
	a.Post("/reports", a.handleRequest(handler.CreateReport))

	// Routing for handling the persons
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(mongoClient *mongo.Client, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.mongoClient, w, r)
	}
}
