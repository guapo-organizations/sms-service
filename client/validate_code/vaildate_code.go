package main

import (
	"context"
	pb "github.com/guapo-organizations/sms-service/proto/sms"
	"github.com/guapo-organizations/sms-service/tls"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	// Set up a connection to the server.
	//tls配置,文件好像是通过第二个参数也就是 x.test.youtube.com生成的...fuck！！！
	creds, err := tls.GetClientTLSFromFile("ca.pem","zldz.com")
	//连接的时候添加tls配置，公钥？不懂
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSmsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//验证验证码
	r, err := c.ValidateCode(ctx, &pb.ValidateCodeRequest{
		Key:         "51785816@qq.com",
		Code:        "oh09G3",
		PublishMode: "email",
		PublishType: "code",
	})

	if err != nil {
		log.Fatalln("rpc发生错误: %v", err)
	}
	log.Printf("%v", r.Result)
}
