package gateway

import (
	"context"
	fmt "fmt"
	"log"
	"testing"

	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
)

func init() {
	viper.SetConfigName(".cobra")
	viper.AddConfigPath("../../")
	viper.ReadInConfig()

}

func NewService() *Service {
	return &Service{}
}

func TestPayList(t *testing.T) {

	s := NewService()
	tests := []struct {
		AppID        int64
		ShouldBE     string
		ShouldBECode string
	}{
		{
			AppID:        1,
			ShouldBE:     "OK",
			ShouldBECode: "200",
		},
	}

	for _, tt := range tests {
		req := &CallPayListReq{
			AppID: tt.AppID,
		}

		resp, _ := s.CallPayList(context.Background(), req)
		if resp.Message != tt.ShouldBE && resp.Code != tt.ShouldBECode {
			t.Errorf("CallPayListRep(%v)=%v, wanted message: %v or code: %v", req, resp.Message, tt.ShouldBE, tt.ShouldBECode)
		}
	}
}

func TestPayListDial(t *testing.T) {
	grpcConn, err := grpc.Dial(fmt.Sprintf("%s:%d", viper.GetString("rpc.host"), viper.GetInt("rpc.port")), grpc.WithInsecure())
	if err != nil {
		log.Printf("rpc dail err:%v", err)
	}
	defer grpcConn.Close()

	client := NewPayCenterSrvClient(grpcConn)
	tests := []struct {
		AppID        int64
		ShouldBE     string
		ShouldBECode string
	}{
		{
			AppID:        1,
			ShouldBE:     "OK",
			ShouldBECode: "200",
		},
		{
			AppID:        2,
			ShouldBE:     "OK",
			ShouldBECode: "200",
		},
	}

	for _, tt := range tests {
		req := &CallPayListReq{
			AppID: tt.AppID,
		}

		resp, _ := client.CallPayList(context.Background(), req)
		if resp.Message != tt.ShouldBE && resp.Code != tt.ShouldBECode {
			t.Errorf("CallPayListRep(%v)=%v, wanted message: %v or code: %v", req, resp.Message, tt.ShouldBE, tt.ShouldBECode)
		}
	}
}
