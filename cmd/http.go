package cmd

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mauwahid/kafman/internal/infra/api"
	"github.com/mauwahid/kafman/internal/infra/config"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunHttp() {

	fmt.Println("==== Start Run HTPP Server ====")

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))
	e.HideBanner = true

	api.SetupApiRoute(e)

	sPort := config.Get().GetString("app.port")
	port := fmt.Sprintf(":%s", sPort)

	// start server
	go func() {
		if err := e.Start(port); err != nil {
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
