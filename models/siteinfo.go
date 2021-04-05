package models

import "github.com/gin-gonic/gin"

//go:generate mockgen -destination=./mock/siteinfo.go -package=mock github.com/shyam81992/Site-Information/models ISiteInfo

type URL struct {
	URL string `form:"url" json:"url" binding:"required"`
}

type ISiteInfo interface {
	GetPageInfo(*gin.Context)
}
