package v1

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/IDarar/notifications-service/pb"

	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

//Each service will be available
type NotificationClient struct {
	Client pb.NotificationsServiceClient
}

func InitNotificationServiceClient(cfg *config.Config) *NotificationClient {
	tlsCredentials, err := loadTLSCredentialsClient(cfg)
	if err != nil {
		logger.Error(err)
		return nil
	}
	transportOption := grpc.WithTransportCredentials(tlsCredentials)
	conn, err := grpc.Dial("0.0.0.0:"+cfg.GRPC.Port, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	c := pb.NewNotificationsServiceClient(conn)
	return &NotificationClient{
		Client: c,
	}
}

func loadTLSCredentialsClient(cfg *config.Config) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	logger.Info(cfg.GRPC)
	pemServerCA, err := ioutil.ReadFile(cfg.GRPC.ClientCACertFile)
	if err != nil {
		logger.Error("err opening cert ", err)
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(cfg.GRPC.ClientCertFile, cfg.GRPC.ClientKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
