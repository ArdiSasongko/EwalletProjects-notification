package api

import (
	"net"

	"github.com/ArdiSasongko/EwalletProjects-notification/internal/mailer"
	protohandler "github.com/ArdiSasongko/EwalletProjects-notification/internal/proto-handler"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/proto/notification"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/storage/sqlc"
	"google.golang.org/grpc"
)

func SetupGRPC() {
	app, err := SetupGRPCApplication()
	if err != nil {
		app.config.logger.Fatalf("failed to setup application (grpc): %v", err)
	}

	// Start listening on the gRPC port
	lis, err := net.Listen("tcp", app.config.addrGRPC)
	if err != nil {
		app.config.logger.Fatalf("failed to listen grpc port, err: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Connect to the database
	conn, err := ConnectDatabase(app.config.db, app.config.logger)
	if err != nil {
		app.config.logger.Fatalf("failed to start database, err: %v", err)
	}

	// Initialize dependencies
	queries := sqlc.New(conn)
	mailer := mailer.NewClientSMTP(app.config.email.fromEmail, app.config.email.apiKey)

	// Create an instance of the notifEmail service
	emailService := protohandler.NewNotifEmail(queries, mailer)

	// Register the notifEmail service with the gRPC server
	notification.RegisterNotificationServiceServer(server, emailService.(*protohandler.NotifEmail))

	// Log that the gRPC server is running
	app.config.logger.Printf("gRPC server is running on port %v", app.config.addrGRPC)

	// Start the gRPC server
	if err := server.Serve(lis); err != nil {
		app.config.logger.Fatalf("failed to start gRPC server, err: %v", err)
	}
}
