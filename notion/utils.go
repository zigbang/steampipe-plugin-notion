package notion

import (
	"context"
	"errors"
	"os"

	"github.com/dstotijn/go-notion"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*notion.Client, error) {
	token := os.Getenv("NOTION_TOKEN")

	notionConfig := GetConfig(d.Connection)
	if notionConfig.Token != nil {
		token = *notionConfig.Token
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	api := notion.NewClient(token)
	return api, nil
}
