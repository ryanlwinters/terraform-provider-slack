package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDefaultChannels() *schema.Resource {
	return &schema.Resource{
		Create: resourceDefaultChannelsCreate,
		Read:   resourceDefaultChannelsRead,
		Update: resourceDefaultChannelsUpdate,
		Delete: resourceDefaultChannelsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"channel_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceDefaultChannelsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Client)
	teamID, channelIDs := readInputs(d)
	if err := callSetDefaultChannels(client, teamID, channelIDs); err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s-default-channels", teamID))
	return resourceDefaultChannelsRead(d, meta)
}

func resourceDefaultChannelsRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDefaultChannelsUpdate(d *schema.ResourceData, meta interface{}) error {
	if d.HasChange("channel_ids") {
		client := meta.(*Client)
		teamID, channelIDs := readInputs(d)
		if err := callSetDefaultChannels(client, teamID, channelIDs); err != nil {
			return err
		}
	}
	return resourceDefaultChannelsRead(d, meta)
}

func resourceDefaultChannelsDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func readInputs(d *schema.ResourceData) (string, []string) {
	teamID := d.Get("team_id").(string)
	channels := d.Get("channel_ids").([]interface{})
	var channelIDs []string
	for _, ch := range channels {
		channelIDs = append(channelIDs, ch.(string))
	}
	return teamID, channelIDs
}

func callSetDefaultChannels(c *Client, teamID string, channelIDs []string) error {
	payload := map[string]interface{}{
		"team_id":     teamID,
		"channel_ids": channelIDs,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	url := c.BaseURL + "admin.teams.settings.setDefaultChannels"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	var lastErr error
	for i := 0; i < 3; i++ {
		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("API request failed: %w", err)
		} else {
			defer resp.Body.Close()
			if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
				if ra := resp.Header.Get("Retry-After"); ra != "" {
					if secs, convErr := strconv.Atoi(ra); convErr == nil {
						time.Sleep(time.Duration(secs) * time.Second)
					} else {
						time.Sleep(time.Duration(2*(i+1)) * time.Second)
					}
				} else {
					time.Sleep(time.Duration(2*(i+1)) * time.Second)
				}
				lastErr = fmt.Errorf("Slack API returned %d", resp.StatusCode)
				continue
			}
			data, readErr := io.ReadAll(resp.Body)
			if readErr != nil {
				return fmt.Errorf("failed to read API response: %w", readErr)
			}
			var result map[string]interface{}
			if err := json.Unmarshal(data, &result); err != nil {
				return fmt.Errorf("failed to decode API response: %w", err)
			}
			if ok, okPresent := result["ok"].(bool); !okPresent || !ok {
				return fmt.Errorf("Slack API error: %v", result["error"])
			}
			return nil
		}
	}
	return lastErr
}