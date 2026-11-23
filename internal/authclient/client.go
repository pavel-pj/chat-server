package authclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	accessv1 "github.com/pavel-pj/auth/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

type Config struct {
	Address    string
	ServerName string
	RootCAs    *x509.CertPool
}

type Client struct {
	conn   *grpc.ClientConn
	access accessv1.AccessV1Client
}

func New(cfg Config) (*Client, error) {
	tlsConfig := &tls.Config{
		RootCAs:    cfg.RootCAs,
		ServerName: cfg.ServerName,
		MinVersion: tls.VersionTLS12,
	}

	conn, err := grpc.NewClient(
		cfg.Address,
		grpc.WithTransportCredentials(
			credentials.NewTLS(tlsConfig),
		),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		access: accessv1.NewAccessV1Client(conn),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Check(ctx context.Context, endpoint string) error {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md.Copy())
	}

	_, err := c.access.Check(ctx, &accessv1.CheckRequest{
		EndpointAddress: endpoint,
	})

	return err
}
