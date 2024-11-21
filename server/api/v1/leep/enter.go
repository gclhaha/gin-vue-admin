package leep

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	VideosApi
	VideoHighlightApi
}

var (
	videosService         = service.ServiceGroupApp.LeepServiceGroup.VideosService
	videoHighlightService = service.ServiceGroupApp.LeepServiceGroup.VideoHighlightService
)
