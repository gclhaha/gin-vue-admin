package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type VideoHighlightSearch struct {
	request.PageInfo
	VideoId string `json:"videoId" form:"videoId"`
	UserId  string `json:"userId" form:"userId"`
}
