package main

import (
	"github.com/bowillkin/ipip-grpc/handler"
	pb "github.com/bowillkin/proto"
	"github.com/bowillkin/proto/ipip"
	"github.com/ipipdotnet/ipdb-go"
	"github.com/koding/multiconfig"
)

type DefaultConfig struct {
	Name    string `default:"ipip-grpc"`
	Address string `default:"0.0.0.0:5000"`
}

func main() {
	var config *DefaultConfig
	config = new(DefaultConfig)
	m := multiconfig.New()
	m.MustLoad(config)
	server := pb.DefaultGrpcServer(config.Name, nil, nil)
	db, err := ipdb.NewCity("ipipfree.ipdb")
	if err != nil {
		pb.GetLogger().Fatal(err)
	}
	ipip.RegisterIpipServer(server, &handler.IpIp{
		DB: db,
	})

	if err := pb.Run(server, config.Address); err != nil {
		pb.GetLogger().Fatalf("run: %s", err)
	}
}
