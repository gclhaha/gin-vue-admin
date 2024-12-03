package leep

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	VideosApi
	VideoHighlightApi
	VenueApi
	VenueItemApi
}

var (
	videosService         = service.ServiceGroupApp.LeepServiceGroup.VideosService
	videoHighlightService = service.ServiceGroupApp.LeepServiceGroup.VideoHighlightService
	venueService          = service.ServiceGroupApp.LeepServiceGroup.VenueService
	venueItemService      = service.ServiceGroupApp.LeepServiceGroup.VenueItemService
)
