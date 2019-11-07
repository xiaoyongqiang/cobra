package rpc

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var grpcConn *grpc.ClientConn
var mtx sync.RWMutex

//Init 初始化redis对象
func Init() error {
	mtx.Lock()
	defer mtx.Unlock()

	log.Printf("Initializing Grpc server ... %v \n", viper.Get("rpc"))
	client, err := grpc.Dial(fmt.Sprintf("%s:%d", viper.GetString("rpc.host"), viper.GetInt("rpc.port")), grpc.WithInsecure())
	if err != nil {
		log.Printf("rpc dail err:%v", err)
		return err
	}

	grpcConn = client
	return nil
}

// Close 关闭redis连接
func Close() {
	if grpcConn != nil {
		grpcConn.Close()
	}
}

//GrpcConn 服务 Dail
func GrpcConn() *grpc.ClientConn {
	mtx.Lock()
	defer mtx.Unlock()

	return grpcConn
}
