package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	config "github.com/Questee29/taxi-app_driverService/configs"
	"github.com/Questee29/taxi-app_driverService/database"
	"github.com/Questee29/taxi-app_driverService/middleware"
	_ "github.com/Questee29/taxi-app_driverService/migrations"
	server "github.com/Questee29/taxi-app_driverService/pkg/grpc"
	grpcHandlers "github.com/Questee29/taxi-app_driverService/pkg/grpc/handler"
	handlers "github.com/Questee29/taxi-app_driverService/pkg/handlers"
	Repository "github.com/Questee29/taxi-app_driverService/pkg/repository/authorization"
	OrderRepository "github.com/Questee29/taxi-app_driverService/pkg/repository/order"
	Service "github.com/Questee29/taxi-app_driverService/pkg/service/authorization"
)

func main() {
	config, err := config.LoadConfig("app", ".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	db, err := database.New()
	if err != nil {
		log.Fatalln(errors.New(`failed to load database`))
	}

	usersRepository := Repository.New(db)
	ordersRepository := OrderRepository.NewTaxiRepository(db)
	authService := Service.New(usersRepository)
	orderService := Service.NewOrderService(ordersRepository)

	handlerSignUp := handlers.NewSignup(authService)
	handlerSignIn := handlers.NewSignIn(authService)
	handlerWelcome := handlers.NewWelcome(authService)
	handlerLogout := handlers.NewLogOut(authService)
	grpcOrderHandler := grpcHandlers.NewOrderHandler(orderService)

	grpcServ := server.NewServer(server.Deps{
		OrderHandler: grpcOrderHandler,
	})

	http.Handle("/sign-up", middleware.SetContentTypeJSON(handlerSignUp))
	http.Handle("/sign-in", middleware.SetContentTypeJSON(handlerSignIn))
	http.Handle("/welcome", middleware.SetContentTypeJSON(middleware.CheckAuthorizedBearer(handlerWelcome, authService)))
	http.Handle("/logout", middleware.CheckAuthorizedBearer(handlerLogout, authService))

	go func() {
		log.Println("Starting listening REST server")
		http.ListenAndServe(config.Server.Port, nil)
	}()

	go func() {
		log.Println("Starting listening GRPC server")
		if err := grpcServ.ListenAndServe(config.Server.GrpcPort); err != nil {
			log.Printf("grpc ListenAndServe err %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Println("shutting down server")
	db.Close()

}
