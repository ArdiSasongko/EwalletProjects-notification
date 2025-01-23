package api

import (
	"net"

	"google.golang.org/grpc"
)

func SetupGRPC() {
	app, err := SetupGRPCApplication()
	if err != nil {
		app.config.logger.Fatalf("failed to setup application (grpc): %v", err)
	}

	lis, err := net.Listen("tcp", app.config.addrGRPC)
	if err != nil {
		app.config.logger.Fatalf("failed to listen grpc port, err: %v", err)
	}

	server := grpc.NewServer()

	_, err = ConnectDatabase(app.config.db, app.config.logger)
	if err != nil {
		app.config.logger.Fatalf("failed to start database, err: %v", err)
	}
	// auth := auth.NewJwt(app.config.auth.secret, app.config.auth.aud, app.config.auth.iss)

	// // register grpc
	// tokenHandler := protohandler.NewTokenService(sqlc.New(db), auth)
	// token.RegisterTokenServiceServer(server, tokenHandler)

	app.config.logger.Printf("grpc server has running, port%v", app.config.addrGRPC)

	if err := server.Serve(lis); err != nil {
		app.config.logger.Fatalf("failed to starting grpc server, err:%v", err)
	}
}
