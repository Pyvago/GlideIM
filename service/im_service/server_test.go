package im_service

import (
	"go_im/pkg/rpc"
	"testing"
)

func TestRunServer(t *testing.T) {
	options := rpc.ServerOptions{
		Addr: "0.0.0.0",
		Port: 9081,
		Name: "im_service",
	}
	err := RunServer(&options)
	if err != nil {
		t.Error(err)
	}
}
