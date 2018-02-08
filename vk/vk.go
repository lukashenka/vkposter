package vk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const apiMethodUrl = "https://api.vk.com/method/"
const oauthUrl = "https://oauth.vk.com/access_token"

type VkAppClient struct {
	AppId       string
	AppSecret   string
	AccessToken string
}
type VkTokenResponse struct {
	AccessToken string `json:"access_token,required"`
}

type VkErrorResponse struct {
	Error struct {
		ErrorCode     int    `json:"error_code"`
		ErrorMsg      string `json:"error_msg"`
		RequestParams []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"request_params"`
	} `json:"error"`
}

func NewVkAppClient(appId string, appSecret string) (*VkAppClient, error) {

	vk := &VkAppClient{
		AppId:     appId,
		AppSecret: appSecret,
	}
	err := vk.InitAccessToken()
	if err != nil {
		return nil, err
	}
	return vk, nil
}

func (vac *VkAppClient) InitAccessToken() (err error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	urlR, err := url.Parse(oauthUrl)
	if err != nil {
		return
	}
	q := urlR.Query()
	q.Add("client_id", vac.AppId)
	q.Add("client_secret", vac.AppSecret)
	q.Add("grant_type", "client_credentials")
	urlR.RawQuery = q.Encode()
	response, err := netClient.Get(urlR.String())
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	vkResp := &VkTokenResponse{}
	err = json.Unmarshal(body, vkResp)
	if err != nil {
		return
	}
	if vkResp.AccessToken == "" {
		err = errors.New("get_access_token_fail")
		return
	}
	vac.AccessToken = vkResp.AccessToken
	return
}

func (vac *VkAppClient) Request(methodName string, params map[string]string) (string, error) {
	u, err := url.Parse(apiMethodUrl + methodName)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	q.Set("access_token", vac.AccessToken)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return string(content), errors.New("bad_request")
	}
	vkErrResp := &VkErrorResponse{}
	json.Unmarshal(content, &vkErrResp)
	if vkErrResp.Error.ErrorCode != 0 {
		return string(content), errors.New("bad_request")
	}

	return string(content), nil
}
