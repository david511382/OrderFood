package swag

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

func Documents(c *gin.Context) {
	if c.Request.RequestURI == "/docs" || c.Request.RequestURI == "/docs/" {
		c.Redirect(http.StatusTemporaryRedirect, "/docs/index.html")
	} else {
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	}
}
