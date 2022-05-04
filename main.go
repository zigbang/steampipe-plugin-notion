package main

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/ygpark80/steampipe-plugin-notion/notion"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: notion.Plugin})
}
