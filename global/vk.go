package global

import (
	"log"

	"github.com/lukashenka/vkposter/vk"
)

var vkAppClient = new(vk.VkAppClient)

func init() {
	vkAppClient = getVkAppClient()
}

func getVkAppClient() *vk.VkAppClient {
	c := GetConfig()
	vkCl, err := vk.NewVkAppClient(c.VkAppId, c.VkAppSecret)
	if err != nil {
		log.Fatalln("Vk error:", err.Error())
		return nil
	}
	return vkCl
}

func GetVkAppClient() *vk.VkAppClient {
	return vkAppClient
}
