# slack_default_channels

Manage default channels assigned to new team members in a Slack team.

## Example Usage

```hcl
resource "slack_default_channels" "default" {
  team_id     = "T1234567890"
  channel_ids = ["C01ABCD2EFG", "C09XYZ3HIJ"]
}
```

## Argument Reference

- `team_id` (String, Required) — Slack team ID.
- `channel_ids` (List(String), Required) — Channel IDs to set as default.

## Import

Import using the ID format `<team_id>-default-channels`:

```bash
terraform import slack_default_channels.default T1234567890-default-channels
```
