package main

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/guapo-organizations/go-micro-secret/frame_tool"
	grpc_gateway_service_info "github.com/guapo-organizations/go-micro-secret/frame_tool/service"
	"github.com/guapo-organizations/sms-service/tls"
	gw "github.com/guapo-organizations/sms-service/proto/sms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net/http"
)

//grpc网关

func run() error {
	my_frame_tool := frame_tool.LyFrameTool{
		ConfigPath: "./config",
	}
	my_frame_tool.Run()
	gate_way_service_info := grpc_gateway_service_info.GetGrpcGateWayServiceInfo()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//grpc网关服务
	mux := runtime.NewServeMux()

	//ssl/tls 初始化
	creds, err := credentials.NewClientTLSFromFile(tls.Path("ca.pem"), "zldz.com")
	//带ssl/tls的拨号
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	err = gw.RegisterSmsHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%s", gate_way_service_info.Ip, gate_way_service_info.Port), opts)
	if err != nil {
		return err
	}

	//grpc网关的服务端监听,也就是开启Web服务
	return http.ListenAndServe(fmt.Sprintf(":%s", gate_way_service_info.GatewayPort), mux)
}

func main() {
	//不知道这个是什么，github上是这样写的，照样抄就对了,应该是个日志组件之类的
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
