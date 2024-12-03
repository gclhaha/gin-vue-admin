package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/leep"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	leepDb := global.GetGlobalDBByDBName("leep")
	leepDb.AutoMigrate(leep.Videos{}, leep.Venue{}, leep.VenueItem{})
	return nil
}
