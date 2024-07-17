package client

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"go-axon/internal/pb"
	"google.golang.org/grpc"
	"time"
)

type CommandGateway interface {
	Send(ctx context.Context, name string, command interface{}) error
}

type commandGateway struct {
	clientId      string
	componentName string
	client        pb.CommandServiceClient
}

func NewCommandGateway(options Options, grpcConnection *grpc.ClientConn) CommandGateway {
	return &commandGateway{
		clientId:      options.ClientId,
		componentName: options.ComponentName,
		client:        pb.NewCommandServiceClient(grpcConnection),
	}
}

func (cg *commandGateway) Send(ctx context.Context, name string, command interface{}) error {
	payload, err := marshalPayload(command)
	if err != nil {
		return err
	}

	pbCommand := &pb.Command{
		MessageIdentifier:      uuid.New().String(),
		Name:                   name,
		Timestamp:              time.Now().UnixNano() / int64(time.Millisecond),
		Payload:                payload,
		MetaData:               nil, //TODO: populate metadata from context
		ProcessingInstructions: nil,
		ClientId:               cg.clientId,
		ComponentName:          cg.componentName,
	}

	response, err := cg.client.Dispatch(ctx, pbCommand)
	if err != nil {
		return err
	}

	if response.ErrorMessage != nil {
		return CommandDispatchingError{
			Message:   response.ErrorMessage.Message,
			Location:  response.ErrorMessage.Location,
			Details:   response.ErrorMessage.Details,
			ErrorCode: response.ErrorMessage.ErrorCode,
		}
	}

	return nil
}

func marshalPayload(payload interface{}) (*pb.SerializedObject, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return &pb.SerializedObject{
		Type:     "application/json",
		Revision: "",
		Data:     data,
	}, nil
}
