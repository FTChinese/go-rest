package go_rest

import (
	"github.com/FTChinese/go-rest/enum"
	"github.com/tomasen/realip"
	"net/http"
)

// ClientApp records the header information of a request.
type ClientApp struct {
	ClientType enum.Platform
	Version    string
	UserIP     string
	UserAgent  string
}

// NewClientApp collects information from a request.
func NewClientApp(req *http.Request) ClientApp {
	c := ClientApp{}

	c.ClientType, _ = enum.ParsePlatform(req.Header.Get("X-Client-Type"))

	c.Version = req.Header.Get("X-Client-Version")

	// Web app must forward user ip and user agent
	// For other client like Android and iOS, request comes from user's device, not our web app.
	if c.ClientType == enum.PlatformWeb {
		c.UserIP = req.Header.Get("X-User-Ip")
		c.UserAgent = req.Header.Get("X-User-Agent")
	} else {
		c.UserIP = realip.FromRequest(req)
		c.UserAgent = req.UserAgent()
	}

	return c
}
