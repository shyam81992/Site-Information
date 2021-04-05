package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Site-Information/config"
	"github.com/shyam81992/Site-Information/controllers"
	"github.com/shyam81992/Site-Information/handler"
	"github.com/shyam81992/Site-Information/scrapper"
)

func main() {

	config.LoadConfig()

	//campq := camqp.CAMQP{}

	siteInfoController := controllers.NewSiteInfoController(&scrapper.Scrapper{})
	controllers.Init(&controllers.Config{
		SiteInfoController: siteInfoController,
	})

	r := gin.Default()

	r.Use(gin.Recovery())

	handler.NewHandler(&handler.Config{
		R:                  r,
		SiteInfoController: siteInfoController,
	})

	r.Run(":" + config.AppConfig["port"])
}
