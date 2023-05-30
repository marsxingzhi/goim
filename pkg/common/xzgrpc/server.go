package xzgrpc

import (
	"fmt"
	"github.com/marsxingzhi/goim/pkg/common/xzetcd"
	"github.com/marsxingzhi/goim/pkg/config"
	"github.com/marsxingzhi/goim/pkg/utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const ttl = 10

type GrpcServer struct {
	grpcConf *config.Grpc
	etcdConf *config.Etcd
}

func NewGrpcServer(grpcConf *config.Grpc, etcdConf *config.Etcd) *GrpcServer {
	return &GrpcServer{grpcConf: grpcConf, etcdConf: etcdConf}
}

func (gs *GrpcServer) Run(s *grpc.Server) {
	addr := "0.0.0.0:" + strconv.Itoa(gs.grpcConf.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("[server] failed to Listen: ", err)
		return
	}

	// TODO 先固定
	host := utils.GetServerIP()
	fmt.Println("host: ", host)

	err = xzetcd.Register(gs.etcdConf.Schema, gs.etcdConf.Endpoints, host, gs.grpcConf.Port, gs.grpcConf.Name, ttl)

	//err = xzetcd.RegisterV2(gs.etcdConf.Schema, host, gs.grpcConf.Port, gs.grpcConf.Name, ttl)
	if err != nil {
		fmt.Println("[xzgrpc] faield to register grpc to etcd: ", err)
		return
	}

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
		return
	}
}
