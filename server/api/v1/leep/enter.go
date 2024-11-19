package leep

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ VideosApi }

var videosService = service.ServiceGroupApp.LeepServiceGroup.VideosService
