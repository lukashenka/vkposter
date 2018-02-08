package global

import (
	"fmt"
	"testing"
)

func TestGetPosts(t *testing.T) {

	appCl := GetVkAppClient()
	request := map[string]string{
		"owner_id":   "-",
		"from_group": "0",
		"message":    "test",
		"signed":     "1",
		"guid":       "test",
	}
	msg, err := appCl.Request("wall.post", request)
	fmt.Println(msg)
	fmt.Println(err)
	if err != nil {
		t.Fatal(err)
	}

}
