package notion

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type notionConfig struct {
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {Type: schema.TypeString},
}

func ConfigInstance() interface{} {
	return &notionConfig{}
}

func GetConfig(connection *plugin.Connection) notionConfig {
	if connection == nil || connection.Config == nil {
		return notionConfig{}
	}
	config, _ := connection.Config.(notionConfig)
	return config
}
