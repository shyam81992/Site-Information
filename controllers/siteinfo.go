package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyam81992/Site-Information/camqp"
	"github.com/shyam81992/Site-Information/models"
	"github.com/shyam81992/Site-Information/scrapper"
)

type SiteInfoCtl struct {
	Scrap scrapper.IScrapper
	camqp.ICAMQP
}

func NewSiteInfoController(scrap scrapper.IScrapper, camqp camqp.ICAMQP) models.ISiteInfo {
	return &SiteInfoCtl{Scrap: scrap, ICAMQP: camqp}
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

func (si *SiteInfoCtl) NotifySiteInfoJob(c *gin.Context) {

	var url models.URL2
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, _ := json.Marshal(url)
	si.Publishmsg(msg)

	c.JSON(200, gin.H{
		"message": "ok",
	})

}
