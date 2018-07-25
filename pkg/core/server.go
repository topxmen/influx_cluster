package core

import (
	"net/http"

	"../metadata"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type ProxyServer struct {
	addr            string
	upstreamClient  *http.Client
	metadataManager *metadata.MetaDataManager
}

func NewProxyServer(addr string) *ProxyServer {
	ps := &ProxyServer{
		upstreamClient: &http.Client{},
	}

	return ps
}

func (ps *ProxyServer) Start() {
	go func() {
		gin.SetMode(gin.ReleaseMode)
		router := gin.New()
		// Recovery middleware recovers from any panics and writes a 500 if there was one.
		router.Use(gin.Recovery())

		router.GET("/ping", ps.pingHandler)

		glog.Fatal(router.Run(ps.addr))
	}()
}

func (ps *ProxyServer) pingHandler(ctx *gin.Context) {
	request := *ctx.Request
	nodes := ps.metadataManager.ListNodes()
	for _, node := range nodes {
		request.URL.Host = node.Address
		resp, err := ps.upstreamClient.Do(&request)
		if err == nil {
			resp.Write(ctx.Writer)
		} else {
			ctx.Status(http.StatusInternalServerError)
		}
		break
	}
}
