package tls

import (
	"google.golang.org/grpc/credentials"
	"path/filepath"
	"runtime"
)

var basepath string

func init() {
	//获取当前文件，也就是tls文件夹的绝对路径
	_, currentFile, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(currentFile)
}

//文件的绝对路径
func path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(basepath, rel)
}

//获取tls客户端加密的组件
//certFile 为证书
//serverNameOverride 是生成证书的那个字符串，例如 zldz.com
func GetClientTLSFromFile(certFile string, serverNameOverride string) (credentials.TransportCredentials, error) {
	certFile = path(certFile)
	creds, err := credentials.NewClientTLSFromFile(certFile, serverNameOverride)
	if err != nil {
		return nil, err
	}
	return creds, nil
}

//获取服务端tls加密的组件
func GetServiceTLSFromFile(certFile, keyFile string) (credentials.TransportCredentials, error) {
	certFile = path(certFile)
	keyFile = path(keyFile)
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

	if err != nil {
		return nil, err
	}
	return creds, nil
}
