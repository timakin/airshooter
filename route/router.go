package route

import (
    "github.com/timakin/airshooter/controller"

    "github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api")
    api = AddNotificationAPI(api)
    api = AddMessageAPI(api)
    return api
}

func AddNotificationAPI(e *echo.Echo) (combined *echo.Echo) {
	combined := e.Group("/notifications")
    combined.POST("/enqueue", func(c echo.Context) error {
        err := controller.ValidationHandler("")
        if err != nil {
            return c.JSON(http.StatusBadRequest, "{}")
        }
        return nil
    }, controller.EnqueueNotification)
	combined.PUT("/publish", controller.PublishNotification)
	combined.GET("/subscriptions", controller.GetNotifications)
	combined.PUT("/subscriptions", controller.MarkReadNotifications)
}

func AddMessageAPI(e *echo.Echo) (combined *echo.Echo) {
	combined := e.Group("/messages")
    combined.POST("", controller.PostMessage)
	combined.GET("", controller.GetMessages)
}
