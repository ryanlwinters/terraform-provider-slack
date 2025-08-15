# Slack Provider

The Slack provider allows managing certain Slack Admin settings.

## Authentication

- Configure with a token via provider `token` or environment variable `SLACK_TOKEN`.
- Admin endpoints require appropriate admin scopes.

## Example Usage

```hcl
provider "slack" {
  token = var.slack_token
}
```

## Resources

- `slack_default_channels` — Manage default channels for new team members.
