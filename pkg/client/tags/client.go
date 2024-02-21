package tags

import (
	"context"
	"errors"
	"fmt"
	"github.com/hell-kitchen/api-gateway/internal/config"
	pb "github.com/hell-kitchen/api-gateway/pkg/api/proto/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ pb.TagsServiceClient = (*Client)(nil)

type Client struct {
	conn *grpc.ClientConn
	cfg  *config.Tags
	cli  pb.TagsServiceClient
}

func New(cfg *config.Tags) (*Client, error) {
	cli := &Client{
		cfg: cfg,
	}

	if err := cli.createConnection(); err != nil {
		return nil, err
	}
	cli.cli = pb.NewTagsServiceClient(cli.conn)
	return cli, nil
}

func (cli *Client) createConnection() (err error) {
	var opts []grpc.DialOption

	if cli.cfg.UseTLS {
		var transportCredentials credentials.TransportCredentials

		if cli.cfg.CertFile == "" || cli.cfg.KeyFile == "" {
			return errors.New("bad config")
		}

		transportCredentials, err = credentials.NewServerTLSFromFile(cli.cfg.CertFile, cli.cfg.KeyFile)
		if err != nil {
			return fmt.Errorf("failed to generate credentials: %w", err)
		}

		opts = []grpc.DialOption{
			grpc.WithTransportCredentials(transportCredentials),
		}
	} else {
		opts = []grpc.DialOption{
			grpc.WithInsecure(),
		}
	}
	cli.conn, err = grpc.Dial(cli.cfg.Addr, opts...)
	return err
}

func (cli *Client) Stop(_ context.Context) error {
	return cli.conn.Close()
}
