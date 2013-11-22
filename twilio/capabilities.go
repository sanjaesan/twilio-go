package twilio

import (
	jwt "github.com/marcuswestin/jwt-go"
	"net/url"
	"strings"
	"time"
)

type Capability struct {
	accountSid   string
	authToken    string
	capabilities []string

	incomingClientName       string
	shouldBuildIncomingScope bool

	shouldBuildOutgoingScope bool
	outgoingParams           map[string]string
	appSid                   string
}

func (c *Client) NewCapability() *Capability {
	capability := new(Capability)
	capability.accountSid = c.AccountSid
	capability.authToken = c.AuthToken
	return capability
}

// Registers this client to accept incoming calls by the given `clientName`.
// If your app TwiML <Dial>s `clientName`, this client will receive the call.
func (c *Capability) AllowClientIncoming(clientName string) {
	c.shouldBuildIncomingScope = true
	c.incomingClientName = clientName
}

// Allows this client to call your application with id `appSid` (See https://www.twilio.com/user/account/apps).
// When the call connects, Twilio will call your voiceUrl REST endpoint.
// The `appParams` argument will get passed through to your voiceUrl REST endpoint as GET or POST parameters.
func (c *Capability) AllowClientOutgoing(appSid string, appParams map[string]string) {
	c.shouldBuildOutgoingScope = true
	c.appSid = appSid
	c.outgoingParams = appParams
}

func (c *Capability) AllowEventStream(filters map[string]string) {
	params := map[string]string{
		"path": "/2010-04-01/Events",
	}
	if len(filters) > 0 {
		params["params"] = url.QueryEscape(generateParamString(filters))
	}
	c.addCapability("stream", "subscribe", params)
}

// Generate the twilio capability token.
// Deliver this token to you JS/iOS/Android Twilio client.
func (c *Capability) GenerateToken(ttl time.Duration) (string, error) {
	c.doBuildIncomingScope()
	c.doBuildOutgoingScope()
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = map[string]interface{}{
		"iss":   c.accountSid,
		"exp":   time.Duration(time.Now().Unix()) + ttl,
		"scope": strings.Join(c.capabilities, " "),
	}
	return token.SignedString([]byte(c.authToken))
}

func (c *Capability) doBuildOutgoingScope() {
	if c.shouldBuildOutgoingScope {
		values := map[string]string{}
		values["appSid"] = c.appSid
		if c.incomingClientName != "" {
			values["clientName"] = c.incomingClientName
		}

		if c.outgoingParams != nil {
			values["appParams"] = generateParamString(c.outgoingParams)
		}

		c.addCapability("client", "outgoing", values)
	}
}

func (c *Capability) doBuildIncomingScope() {
	if c.shouldBuildIncomingScope {
		values := map[string]string{}
		values["clientName"] = c.incomingClientName
		c.addCapability("client", "incoming", values)
	}
}

func (c *Capability) addCapability(service, privelege string, params map[string]string) {
	c.capabilities = append(c.capabilities, scopeUriFor(service, privelege, params))
}

func scopeUriFor(service, privelege string, params map[string]string) string {
	scopeUri := "scope:" + service + ":" + privelege
	if len(params) > 0 {
		scopeUri += "?" + generateParamString(params)
	}
	return scopeUri
}

func generateParamString(params map[string]string) string {
	values := make(url.Values)
	for key, val := range params {
		values.Add(key, val)
	}
	return values.Encode()
}
