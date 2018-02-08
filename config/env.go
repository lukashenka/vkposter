package config

type EnvConfig struct {
	VkGroupFrom         []string `envconfig:"VK_GROUP_FROM"`
	VkGroupFromAsOwner  string   `envconfig:"VK_GROUP_FROM_AS_OWNER"`
	VkGroupTo           string   `envconfig:"VK_GROUP_TO"`
	VkGroupToToken      string   `envconfig:"VK_TO_GROUP_TOKEN_KEY"`
	VkAppId             string   `envconfig:"VK_APP_ID"`
	VkAppSecret         string   `envconfig:"VK_APP_SECRET"`
	RefreshTimePerGroup int      `envconfig:"REFRESH_TIME_PER_GROUP"`
}
