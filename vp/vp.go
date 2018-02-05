package vp

import (
	"log"

	"github.com/lukashenka/vkposter/config"
)

type VkPoster struct {
	FromID      string
	ToId        []string
	RefreshTime int
}

func InitProcessing() *VkPoster {
	c := config.GetConfig()
	vp := &VkPoster{
		FromID: c.VkGroupFrom,
		ToId:   c.VkGroupTo,
	}
	return vp

}

func (vp *VkPoster) Start() {
	log.Printf("Start processing")

}
