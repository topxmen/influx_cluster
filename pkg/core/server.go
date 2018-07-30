package core

import (
	"net/http"

	"../metadata"
	"../model"
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

		//cluster management
		router.GET("/cluster/nodes", ps.listNodesHandler)
		router.POST("/cluster/nodes", ps.createNodeHandler)
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

//metadata management
func (ps *ProxyServer) listNodesHandler(ctx *gin.Context) {
	var resp Response
	defer func() {
		ctx.JSON(http.StatusOK, resp)
	}()
	resp.Data = ps.metadataManager.ListNodes()
	return
}

func (ps *ProxyServer) createNodeHandler(ctx *gin.Context) {
	var resp Response
	defer func() {
		ctx.JSON(http.StatusOK, resp)
	}()

	var node model.Node
	err := ctx.BindJSON(&node)
	if err != nil {
		resp.Code = InvalidParameter
		resp.Data = err.Error()
		return
	}
}
