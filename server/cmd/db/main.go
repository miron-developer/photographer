package main

import (
	"context"
	"net"
	"os"
	"os/signal"

	"photographer/cmd/db/internal"

	"photographer/internal/protobuf/db"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

type DBServiceGRPC struct {
	db.UnimplementedDBServer
}

func newDbServer() *DBServiceGRPC {
	return &DBServiceGRPC{}
}

func (s *DBServiceGRPC) Create(ctx context.Context, p *db.SQLInsertParams) (*db.SQLResult, error) {
	// model, ok := internal.DefineModel(p.Table)
	// if !ok {
	// 	return nil, errors.New("wrong table")
	// }

	return &db.SQLResult{Result: db.SQLResult_CREATE, Data: []*anypb.Any{}}, nil
}

func (s *DBServiceGRPC) Update(ctx context.Context, p *db.SQLUpdateParams) (*db.SQLResult, error) {
	return &db.SQLResult{}, nil
}

func (s *DBServiceGRPC) Delete(ctx context.Context, p *db.SQLDeleteParams) (*db.SQLResult, error) {
	return &db.SQLResult{}, nil
}

func (s *DBServiceGRPC) Select(ctx context.Context, p *db.SQLSelectParams) (*db.SQLResult, error) {
	return &db.SQLResult{}, nil
}

func main() {
	// initialize
	dbCtx := internal.Init()
	defer dbCtx.Log.Writer().Close()

	// listen db service place
	lis, e := net.Listen("tcp", dbCtx.Config.S_DB_ADDR)
	if e != nil {
		dbCtx.Log.Fatalf("failed to listen: %v\n", e)
	}

	// create grpc server and listen on above place
	srv := grpc.NewServer()
	db.RegisterDBServer(srv, newDbServer())
	go func() {
		if e := srv.Serve(lis); e != nil {
			dbCtx.Log.Fatal("listen grpc service error: ", e)
		}
	}()

	// Setting up signal capturing
	var stop = make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	// Waiting for SIGINT (pkill -2)
	<-stop

	// after stop signal actions here

}
