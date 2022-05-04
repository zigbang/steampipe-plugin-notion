install:
	go build -o  ~/.steampipe/plugins/hub.steampipe.io/plugins/ygpark80/notion@latest/steampipe-plugin-notion.plugin *.go

local:
	go build -o  ~/.steampipe/plugins/local/notion/notion.plugin *.go
