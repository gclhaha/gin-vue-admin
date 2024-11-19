package leep

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ VideosRouter }

var videosApi = api.ApiGroupApp.LeepApiGroup.VideosApi
