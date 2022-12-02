package vk_service

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"github.com/woodhds/vk.service/api/handlers"
	"github.com/woodhds/vk.service/database"
	vkMessages "github.com/woodhds/vk.service/gen/messages"
	vkUsers "github.com/woodhds/vk.service/gen/users"
	"github.com/woodhds/vk.service/internal/messages"
	"github.com/woodhds/vk.service/internal/users"
	"log"
	"net/http"
	"time"
)

type App struct {
	router              *mux.Router
	messageQueryService database.MessagesQueryService
	usersQueryService   database.UsersQueryService
	token               string
	version             string
	count               int
	factory             database.ConnectionFactory
	messagesService     handlers.VkMessagesService
	grpcMux             *runtime.ServeMux
	srv                 *http.Server
}

func (app *App) Initialize() {
	router := mux.NewRouter()
	app.router = router.PathPrefix("/api").Subrouter()
	app.grpcMux = runtime.NewServeMux()

	app.initializeRoutes()

}

func (app *App) Run(port int) {

	app.srv = &http.Server{
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      cors.Default().Handler(app.router),
		Addr:         fmt.Sprintf(":%d", port),
	}

	if err := app.srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (app *App) Stop(ctx context.Context) {
	app.srv.Shutdown(ctx)
	log.Println("shutting down")
}

func NewApp(
	messageQueryService database.MessagesQueryService,
	usersQueryService database.UsersQueryService,
	factory database.ConnectionFactory,
	messagesService handlers.VkMessagesService,
	token string,
	version string,
	count int) *App {
	return &App{
		router:              nil,
		messageQueryService: messageQueryService,
		usersQueryService:   usersQueryService,
		token:               token,
		version:             version,
		count:               count,
		factory:             factory,
		messagesService:     messagesService,
	}
}

func (app *App) initializeRoutes() {
	app.router.Path("/like").Handler(handlers.LikeHandler()).Methods(http.MethodPost)
	app.router.Path("/grab").Handler(handlers.ParserHandler(app.factory, app.messagesService, app.count, app.usersQueryService)).Methods(http.MethodGet)

	vkMessages.RegisterMessagesServiceHandlerServer(context.Background(), app.grpcMux, messages.NewMessageHandler(app.messageQueryService, app.token, app.version, app.factory))
	vkUsers.RegisterUsersServiceHandlerServer(context.Background(), app.grpcMux, users.NewUsersHandler(app.usersQueryService, app.token, app.version))
	app.router.PathPrefix("").Handler(app.grpcMux)
}
