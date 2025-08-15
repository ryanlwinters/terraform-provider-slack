package provider

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Client struct {
	HTTPClient	*http.Client
	Token		string
	BaseURL		string
	UserAgent	string
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SLACK_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"slack_default_channels": resourceDefaultChannels(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	token := d.Get("token").(string)

	client := &Client{
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
		Token:      token,
		BaseURL:    "https://slack.com/api/",
		UserAgent:  "terraform-provider-slack/" + Version,
	}

	return client, diags
}