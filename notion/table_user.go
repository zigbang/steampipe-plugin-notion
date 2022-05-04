package notion

import (
	"context"

	"github.com/dstotijn/go-notion"
	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableNotionUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "notion_user",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING},
			{Name: "type", Type: proto.ColumnType_STRING},
			{Name: "name", Type: proto.ColumnType_STRING},
			{Name: "person_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Person.Email")},
			{Name: "avatar_url", Type: proto.ColumnType_STRING},
		},
	}
}

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// logger := plugin.Logger(ctx)

	client, _ := connect(ctx, d)
	startCursor := ""
	for {
		res, _ := client.ListUsers(ctx, &notion.PaginationQuery{
			StartCursor: startCursor,
			PageSize:    100,
		})

		for _, t := range res.Results {
			if t.Type == "person" {
				d.StreamListItem(ctx, t)
			}
		}

		if !res.HasMore {
			break
		}

		startCursor = *res.NextCursor
	}

	return nil, nil
}
