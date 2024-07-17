package client

import (
	"context"
	"fmt"
	"go-axon/internal/pb"
	"google.golang.org/grpc"
)

type AxonClient interface {
	CommandGateway() CommandGateway
}

// Options for the configuration of the Axon client.
type Options struct {
	// Address is the axon server address to connect to.
	Host string

	// Port is the port of the axon server to connect to.
	Port int

	// ComponentName is the name of the component. Several instances of the same component should share this name.
	ComponentName string

	// ClientId is the unique identifier of the client. Is used to distinguish different instances of the same component.
	ClientId string

	// Tags are optional and can be used to provide hints and preferences for setting up connections.
	Tags map[string]string

	// Version is the Axon Framework version of the client.
	Version string
}

func Dial(ctx context.Context, options Options) (AxonClient, error) {
	address := fmt.Sprintf("%s:%d", options.Host, options.Port)
	grpcConnection, err := grpc.NewClient(address)
	if err != nil {
		return nil, err
	}

	platformClient := pb.NewPlatformServiceClient(grpcConnection)
	platform, err := platformClient.GetPlatformServer(ctx, &pb.ClientIdentification{
		ClientId:      options.ClientId,
		ComponentName: options.ComponentName,
		Tags:          options.Tags,
		Version:       options.Version,
	})
	if err != nil {
		return nil, err
	}

	if platform.SameConnection {
		return &axonClient{
			commandGateway: NewCommandGateway(options, grpcConnection),
		}, nil
	}

	err = grpcConnection.Close()
	if err != nil {
		return nil, err
	}

	address = fmt.Sprintf("%s:%d", platform.Primary.HostName, platform.Primary.GrpcPort)
	nodeGrpcConnection, err := grpc.NewClient(address)
	if err != nil {
		return nil, err
	}

	return &axonClient{
		commandGateway: NewCommandGateway(options, nodeGrpcConnection),
	}, nil
}

type axonClient struct {
	commandGateway CommandGateway
}

func (c *axonClient) CommandGateway() CommandGateway {
	return c.commandGateway
}
