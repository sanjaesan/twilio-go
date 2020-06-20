package twilio

import (
	"context"
	"encoding/json"
	"net/url"
)

const channelPathPart = "Channels"

//ChannelService -
type ChannelService struct {
	client *Client
}

//Channel -
type Channel struct {
	Sid             string            `json:"sid"`
	AccountSid      string            `json:"account_sid"`
	ServiceSid      json.RawMessage   `json:"service_sid"`
	CreatedBy       string            `json:"created_by"`
	FriendlyName    string            `json:"friendly_name"`
	UniqueName      string            `json:"unique_name"`
	Type            string            `json:"type"`
	MessageCount    int64             `json:"message_count"`
	DateCreated     TwilioTime        `json:"date_created"`
	DateUpdated     TwilioTime        `json:"date_updated"`
	SubresourceURIs map[string]string `json:"subresource_uris"`
	URI             string            `json:"uri"`
}

//ChannelPage -
type ChannelPage struct {
	Page
	Channels []*Channel `json:"channels"`
}

// Create -
func (c *ChannelService) Create(ctx context.Context, data url.Values) (*Channel, error) {
	channel := new(Channel)
	err := c.client.CreateResource(ctx, channelPathPart, data, channel)
	return channel, err
}

//Update -
func (c *ChannelService) Update(ctx context.Context, sid string, data url.Values) (*Channel, error) {
	channel:= new(Channel)
	err := c.client.UpdateResource(ctx, channelPathPart, sid, data, channel)
	return channel, err
}

// Get -
func (c *ChannelService) Get(ctx context.Context, channelServiceID string, sid string) (*Channel, error) {
	channel := new(Channel)
	err := c.client.GetResource(ctx, servicesPathPart+"/"+channelServiceID+"/"+channelPathPart, sid, channel)
	return channel, err
}
