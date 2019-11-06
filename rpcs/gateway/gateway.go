package gateway

import (
	fmt "fmt"

	context "golang.org/x/net/context"
)

// Service provides main service
type Service struct {
}

// CallPayList returns a list of payment types that are enabled.
func (s *Service) CallPayList(ctx context.Context, req *CallPayListReq) (*CallPayListRep, error) {
	resp := &CallPayListRep{
		Code:    "200",
		Message: "OK",
	}

	fmt.Printf("query data: %v ...", req)
	return resp, nil
}
