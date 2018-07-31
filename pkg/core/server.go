package core

import (
	"net/http"
	"time"

	"../model"
	"../util"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type ProxyServer struct {
	addr            string
	upstreamClient  *http.Client
	metadataManager *model.MetaDataManager
}

func NewProxyServer(addr string, etcdEndpoints []string, etcdDailTimeout time.Duration, serverID uint64) *ProxyServer {
	ps := &ProxyServer{
		upstreamClient:  &http.Client{},
		addr:            addr,
		metadataManager: model.NewMetaDataManager(etcdEndpoints, etcdDailTimeout, util.NewIDGenerator(serverID)),
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
		router.PUT("/cluster/nodes", ps.updateNodeHandler)
		glog.Fatal(router.Run(ps.addr))
	}()
}

func (ps *ProxyServer) pingHandler(ctx *gin.Context) {
	request := *ctx.Request
	nodes, err := ps.metadataManager.ListNodes()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
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
	nodes, err := ps.metadataManager.ListNodes()
	if err != nil {
		resp.Code = InternalServerError
		resp.Data = err.Error()
	}

	resp.Data = nodes
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

	id, err := ps.metadataManager.AddNode(node)
	if err != nil {
		resp.Code = InternalServerError
		resp.Data = err.Error()

		return
	}

	resp.Data = id
	return
}

func (ps *ProxyServer) updateNodeHandler(ctx *gin.Context) {
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

	err = ps.metadataManager.UpdateNode(node)
	if err != nil {
		resp.Code = InternalServerError
		resp.Data = err.Error()

		return
	}

	return
}
