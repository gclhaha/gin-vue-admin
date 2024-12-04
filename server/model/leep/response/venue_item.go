package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/leep"

type VenueItemDto struct {
	Id           string `json:"id"`
	VenueId      string `json:"venueId"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	Name         string `json:"name"`
	VenueName    string `json:"venueName"`
}

func ToVenueItemDto(venueItem leep.VenueItem) VenueItemDto {
	return VenueItemDto{
		Id:           venueItem.Id,
		VenueId:      venueItem.VenueId,
		ThumbnailUrl: venueItem.ThumbnailUrl,
		Name:         venueItem.Name,
		VenueName:    "", // This will be set later
	}
}
