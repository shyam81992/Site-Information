package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Site-Information/models"
	"github.com/shyam81992/Site-Information/scrapper"
)

type SiteInfoCtl struct {
	Scrap scrapper.IScrapper
}

func NewSiteInfoController(scrap scrapper.IScrapper) models.ISiteInfo {
	return &SiteInfoCtl{Scrap: scrap}
}

//CreateCity function
func (si *SiteInfoCtl) GetPageInfo(c *gin.Context) {

	var url models.URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := si.Scrap.GetPageInfo(url.URL)

	c.JSON(200, result)

}
