terraform {
  required_providers {
    slack = {
      source = "ryanlwinters/slack"
    }
  }
}

variable "slack_token" {
  type      = string
  sensitive = true
}

provider "slack" {
  token = var.slack_token
}

resource "slack_default_channels" "default" {
  team_id     = "T1234567890"
  channel_ids = ["C01ABCD2EFG", "C09XYZ3HIJ"]
}


