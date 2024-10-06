package ports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gsy/store/pkg/user/application"
	"github.com/gsy/store/pkg/user/application/command"
)


type HttpServer struct {
	app application.Application
}

func NewHttpServer(app application.Application) HttpServer {
	return HttpServer{app: app}
}


func (h HttpServer) RegisterHandlers(group *gin.RouterGroup) {
	group.POST("/register", h.RegisterUser)
	group.POST("/login", h.LoginUser)
}


func (h HttpServer) RegisterUser(c *gin.Context) {
	username := c.Query("username")
	_, err := h.app.Commands.RegisterUser.Handle(c, &command.RegisterUser{Username: username})
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h HttpServer) LoginUser(c *gin.Context) {
	username := c.Query("username")
	if err := h.app.Commands.LoginUser.Handle(c, &command.LoginUser{Username: username}); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
