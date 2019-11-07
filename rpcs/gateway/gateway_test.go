package gateway

import (
	"cobra/pkg/rpc"
	"context"
	"testing"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".cobra")
	viper.AddConfigPath("../../")
	viper.ReadInConfig()
	rpc.Init()

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

	client := NewPayCenterSrvClient(rpc.GrpcConn())
	defer rpc.Close()

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
