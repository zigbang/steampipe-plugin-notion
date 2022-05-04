# Table: notion_user

notion_user includes users in a Notion workspace. Users include full workspace members and bots but guests are not included. Refer to [User object](https://developers.notion.com/reference/user) for more information.

## Examples

### List all Notion members

```sql
select
  *
from notion_user
where
  type = 'person';
```

### List all Notion bots

```sql
select
  *
from notion_user
where
  type = 'bot';
```
