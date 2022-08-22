package controller

import (
	"errors"
	"fmt"
	"github.com/NubeIO/rubix-edge/service/system"

	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Builder(ip string, port int) (*url.URL, error) {
	return url.ParseRequestURI(CheckHTTP(fmt.Sprintf("%s:%d", ip, port)))
}

func CheckHTTP(address string) string {
	if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
		return "http://" + address
	}
	return address
}
func setInternalToken(token string) string {
	return fmt.Sprintf("Internal %s", token)
}

func (inst *Controller) FFProxy(c *gin.Context) {
	remote, err := Builder("0.0.0.0", 1660)
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	sys := system.New(&system.System{})
	data, err := sys.GetFlowToken()
	if err != nil {
		reposeHandler(nil, err, c)
		return
	}
	token := data.Token
	if token == "" {
		reposeHandler(nil, errors.New("flow-framework token is empty"), c)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
		req.Header.Set("Authorization", setInternalToken(token))

	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
