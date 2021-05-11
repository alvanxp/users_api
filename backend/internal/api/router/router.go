package router

import (
	"fmt"
	"io"
	"os"
	"users_api/internal/api/controllers"
	"users_api/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())
	// Routes
	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	secure_api := app.Group("/api")

	secure_api.Use(middlewares.AuthRequired())
	// ================== User Routes
	secure_api.GET("/users", controllers.GetUsers)
	secure_api.GET("/users/:id", controllers.GetUserById)

	public_api := app.Group("/api")

	// ================== Login Routes
	public_api.POST("/login", controllers.Login)
	public_api.POST("/login/register", controllers.RegisterUser)

	return app
}
