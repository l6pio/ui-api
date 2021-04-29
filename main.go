package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"l6p.io/ui-api/pkg/addon"
	"l6p.io/ui-api/pkg/cfg"
	"l6p.io/ui-api/pkg/middlewares"
	"l6p.io/ui-api/pkg/v1/router"
	"net/http"
)

func main() {
	var dbAddr, dbUser, dbPass string

	cmd := &cobra.Command{
		Use:   "api",
		Short: "api",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf := cfg.New(dbAddr, dbUser, dbPass)
			println("MongoDB Addr: ", conf.DbAddr)
			println("MongoDB User: ", conf.DbUser)
			return start(conf)
		},
	}

	// >>> DEBUG >>>
	//cmd.Flags().StringVar(&dbAddr, "dbAddr", "", "MongoDB address")
	//cmd.Flags().StringVar(&dbUser, "dbUser", "", "MongoDB username")
	//cmd.Flags().StringVar(&dbPass, "dbPass", "", "MongoDB password")
	// *************
	cmd.Flags().StringVar(&dbAddr, "dbAddr", "localhost:32017", "MongoDB address")
	cmd.Flags().StringVar(&dbUser, "dbUser", "root", "MongoDB username")
	cmd.Flags().StringVar(&dbPass, "dbPass", "rootpassword", "MongoDB password")
	// <<< END <<<

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func start(conf *cfg.Config) error {
	server := echo.New()
	server.HideBanner = true

	server.Use(middleware.CORS())
	server.Use(middlewares.Config(conf)...)

	addon.AddValidator(server)

	apiV1 := server.Group("/api/v1")
	router.PingRouter(apiV1.Group("/ping"))
	router.JobRouter(apiV1.Group("/job"))
	router.ReportRouter(apiV1.Group("/report"))

	server.HTTPErrorHandler = ErrorHandler

	return server.Start(":1323")
}

func ErrorHandler(err error, ctx echo.Context) {
	log.Error(err)

	code := http.StatusBadRequest
	if httpError, ok := err.(*echo.HTTPError); ok {
		code = httpError.Code
	}

	e := ctx.JSON(code, struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})
	if e != nil {
		log.Error(err)
	}
}
