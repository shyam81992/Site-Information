package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Site-Information/models"
)

type Handler struct {
	siteInfoController models.ISiteInfo
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R                  *gin.Engine
	SiteInfoController models.ISiteInfo
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		siteInfoController: c.SiteInfoController,
	} // currently has no properties
	r := c.R.Group("/")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/static/siteinfo", h.siteInfoController.GetPageInfo)
	r.POST("/dynamic/siteinfo", h.siteInfoController.GetPageInfo)

}
