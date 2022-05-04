# Notion Plugin for Steampipe

## Installing and Testing the Plugin

To install the plugin, simple run the following command.

```
% make local
go build -o  ~/.steampipe/plugins/local/notion/notion.plugin *.go
```

Check your local plugin using the following command.

```
% steampipe plugin list
+--------------------------------------------------+---------+-------------+
| Name                                             | Version | Connections |
+--------------------------------------------------+---------+-------------+
| hub.steampipe.io/plugins/turbot/aws@latest       | 0.57.0  | aws         |
| hub.steampipe.io/plugins/turbot/steampipe@latest | 0.2.0   | steampipe   |
| local/notion                                     | local   |             |
+--------------------------------------------------+---------+-------------+
```

Copy the sample `notion.spc` file to `~/.steampipe/config` folder and change the name of the `plugin` from `notion` to `local/notion`.

```
% cat ~/.steampipe/config/notion.spc
connection "notion" {
    plugin = "local/notion"

    token = "YOUR_INTEGRATION_TOKEN_HERE"
}
```

Check and see if you have a valid connection.

```
% steampipe plugin list
+--------------------------------------------------+---------+-------------+
| Name                                             | Version | Connections |
+--------------------------------------------------+---------+-------------+
| hub.steampipe.io/plugins/turbot/aws@latest       | 0.57.0  | aws         |
| hub.steampipe.io/plugins/turbot/steampipe@latest | 0.2.0   | steampipe   |
| local/notion                                     | local   | notion      |
+--------------------------------------------------+---------+-------------+
```

Let's test the plugin

```
% steampipe query "select count(*) from notion_user" --timing
+-------+
| count |
+-------+
| 406   |
+-------+

Time: 3.909126417s
```

That's it.
