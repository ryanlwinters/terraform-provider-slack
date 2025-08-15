# Terraform Provider Slack (Unofficial)

A custom Terraform provider for Slack Admin APIs. Currently supports managing default channels for new team members via `slack_default_channels`.

## Installation (local development)

1. Build the provider:

```bash
cd terraform-provider-slack
go build -o terraform-provider-slack
```

2. Create a CLI config to use a dev override:

```hcl
# ~/.terraformrc or a temp file set via TF_CLI_CONFIG_FILE
provider_installation {
  dev_overrides {
    "ryanlwinters/slack" = "/absolute/path/to/terraform-provider-slack"
  }
  direct {}
}
```

3. Use the example in `examples/default-channels` and run `terraform plan`.

## Usage

```hcl
terraform {
  required_providers {
    slack = {
      source = "ryanlwinters/slack"
    }
  }
}

provider "slack" {
  token = var.slack_token # or env SLACK_TOKEN
}

resource "slack_default_channels" "default" {
  team_id     = "T1234567890"
  channel_ids = ["C01ABCD2EFG", "C09XYZ3HIJ"]
}
```

## Authentication & Permissions

- Token: Bot or user token with appropriate Admin scopes. For Enterprise/Admin endpoints, an admin token is required.
- Example scopes: `admin.teams:write`, `channels:read` (verify in Slack docs for your workspace type).

## Development

- Go: 1.23+ (toolchain may use 1.24.x)
- SDK: `github.com/hashicorp/terraform-plugin-sdk/v2`

### Make targets

- `make build` - build provider
- `make test` - run tests
- `make docs` - generate docs (future)
- `make acc` - run acceptance tests (requires TF_ACC=1)

## License

Apache-2.0

## Publishing to Terraform Registry

1. Generate a GPG key (if you want to sign checksums yourself) and add it to repo secrets as `GPG_PRIVATE_KEY` and `GPG_PASSPHRASE`.
2. Tag a release:

```bash
git tag v0.1.0
git push origin v0.1.0
```

3. GitHub Actions will build and attach release artifacts. Once the release is live, publish the provider on the Terraform Registry under `ryanlwinters/slack`.
