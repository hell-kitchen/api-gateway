package tags

import (
	"context"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/tags"
	"google.golang.org/grpc"
)

func (cli *Client) Create(ctx context.Context, in *pb.CreateRequest, opts ...grpc.CallOption) (*pb.CreateResponse, error) {
	return cli.cli.Create(ctx, in, opts...)
}

func (cli *Client) CreateMany(ctx context.Context, in *pb.CreateManyRequest, opts ...grpc.CallOption) (*pb.CreateManyResponse, error) {
	return cli.cli.CreateMany(ctx, in, opts...)
}

func (cli *Client) Get(ctx context.Context, in *pb.GetRequest, opts ...grpc.CallOption) (*pb.GetResponse, error) {
	return cli.cli.Get(ctx, in, opts...)
}

func (cli *Client) GetAll(ctx context.Context, in *pb.GetAllRequest, opts ...grpc.CallOption) (*pb.GetAllResponse, error) {
	return cli.cli.GetAll(ctx, in, opts...)
}

func (cli *Client) Delete(ctx context.Context, in *pb.DeleteRequest, opts ...grpc.CallOption) (*pb.DeleteResponse, error) {
	return cli.cli.Delete(ctx, in, opts...)
}

func (cli *Client) Update(ctx context.Context, in *pb.UpdateRequest, opts ...grpc.CallOption) (*pb.UpdateResponse, error) {
	return cli.cli.Update(ctx, in, opts...)
}
