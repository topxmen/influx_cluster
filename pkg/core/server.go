package core

import (
	"net/http"

	"../model"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type ProxyServer struct {
	addr            string
	upstreamClient  *http.Client
	metadataManager *model.MetaDataManager
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

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

//metadata management
func (ps *ProxyServer) listNodesHandler(ctx *gin.Context) {
	var resp Response
	defer func() {
		ctx.JSON(http.StatusOK, resp)
	}()
	nodes := ps.metadataManager.ListNodes()
	resp.Data = nodes

	return
}

func (ps *ProxyServer) createNodeHandler(ctx *gin.Context) {

}
