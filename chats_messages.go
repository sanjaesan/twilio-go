package twilio

import (
	"context"
	"encoding/json"
	"net/url"
)

//ChatMessageService -
type ChatMessageService struct {
	client *Client
}

//ChatMessage -
type ChatMessage struct {
	Sid             string            `json:"sid"`
	AccountSid      string            `json:"account_sid"`
	ServiceSid      json.RawMessage   `json:"service_sid"`
	ChannelSid      string            `json:"channel_sid"`
	To              User              `json:"to"`
	Body            string            `json:"body"`
	From            User              `json:"from"`
	CreatedBy       string            `json:"created_by"`
	Type            string            `json:"type"`
	DateCreated     TwilioTime        `json:"date_created"`
	DateUpdated     TwilioTime        `json:"date_updated"`
	URI             string            `json:"uri"`
}

//ChatMessagePage -
type ChatMessagePage struct {
	Page
	ChatMessages []*ChatMessage `json:"chat_messages"`
}

// Create -
func (c *ChatMessageService) Create(ctx context.Context, data url.Values) (*ChatMessage, error) {
	channel := new(ChatMessage)
	err := c.client.CreateResource(ctx, channelPathPart, data, channel)
	return channel, err
}

//Update -
func (c *ChatMessageService) Update(ctx context.Context, sid string, data url.Values) (*ChatMessage, error) {
	channel := new(ChatMessage)
	err := c.client.UpdateResource(ctx, channelPathPart, sid, data, channel)
	return channel, err
}

// Get -
func (c *ChatMessageService) Get(ctx context.Context, channelServiceID string, sid string) (*ChatMessage, error) {
	channel := new(ChatMessage)
	err := c.client.GetResource(ctx, servicesPathPart+"/"+channelServiceID+"/"+channelPathPart, sid, channel)
	return channel, err
}
