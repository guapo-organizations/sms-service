package test

import (
	"context"
	"github.com/guapo-organizations/go-micro-secret/tls"
	pb "github.com/guapo-organizations/sms-service/proto/sms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"testing"
	"time"
)

//email 发送验证测试
//func TestSendEmailCode(t *testing.T) {
//	// Set up a connection to the server.
//	//tls配置,文件好像是通过第二个参数也就是 x.test.youtube.com生成的...fuck！！！
//	creds, err := credentials.NewClientTLSFromFile(tls.Path("ca.pem"), "zldz.com")
//	//连接的时候添加tls配置，公钥？不懂
//	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(creds))
//
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//
//	c := pb.NewSmsClient(conn)
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	//发送验证码
//	r, err := c.SendEmailCode(ctx, &pb.SendEmailCodeRequest{
//		Email: "51785816@qq.com",
//	})
//
//	if err != nil {
//		t.Fatalf("rpc发生错误: %v", err)
//	}
//
//	t.Logf("结果:%v", r)
//}

//
////验证码校验测试
func TestValidateEmailCode(t *testing.T) {
	t.Log("什么鬼？？")
	// Set up a connection to the server.
	//tls配置,文件好像是通过第二个参数也就是 x.test.youtube.com生成的...fuck！！！
	creds, err := credentials.NewClientTLSFromFile(tls.Path("ca.pem"), "zldz.com")
	//连接的时候添加tls配置，公钥？不懂
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	t.Log("什么鬼？？")
	c := pb.NewSmsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	t.Log("什么鬼？？")
	//验证验证码
	r, err := c.ValidateCode(ctx, &pb.ValidateCodeRequest{
		Key:         "51785816@qq.com",
		Code:        "123456",
		PublishMode: "email",
		PublishType: "code",
	})

	if err != nil {
		t.Log("什么鬼？？")
		t.Fatalf("rpc发生错误: %v", err)
	}
	t.Log("什么鬼？？")
	t.Log(r)
	t.Logf("%v",r.Result)
}
