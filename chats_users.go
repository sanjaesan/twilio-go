package twilio

import (
	"context"
	"encoding/json"
	"net/url"
)

const userPathPart = "Users"

//UserService -
type UserService struct {
	client *Client
}

//User -
type User struct {
	Sid        string          `json:"sid"`
	AccountSid string          `json:"account_sid"`
	ServiceSid json.RawMessage `json:"service_sid"`
	RoleSid    string          `json:"role_sid"`
	Identity   string          `json:"identity"`
	// Attributes       Attributes        `json:"attributes"`
	JoinedUsersCount int64             `json:"joined_users_count"`
	IsOnline         bool              `json:"is_online"`
	IsNotifiable     bool              `json:"is_notifiable"`
	FriendlyName     string            `json:"friendly_name"`
	DateCreated      TwilioTime        `json:"date_created"`
	DateUpdated      TwilioTime        `json:"date_updated"`
	Links            map[string]string `json:"subresource_uris"`
	URL              string            `json:"uri"`
}

//UserPage -
type UserPage struct {
	Page
	Users []*User `json:"users"`
}

// Create -
func (u *UserService) Create(ctx context.Context, data url.Values) (*User, error) {
	user := new(User)
	err := u.client.CreateResource(ctx, userPathPart, data, user)
	return user, err
}

//Update -
func (u *UserService) Update(ctx context.Context, sid string, data url.Values) (*User, error) {
	user := new(User)
	err := u.client.UpdateResource(ctx, userPathPart, sid, data, user)
	return user, err
}

// Get -
func (u *UserService) Get(ctx context.Context, userServiceID string, sid string) (*User, error) {
	user := new(User)
	err := u.client.GetResource(ctx, servicesPathPart+"/"+userServiceID+"/"+userPathPart, sid, user)
	return user, err
}
