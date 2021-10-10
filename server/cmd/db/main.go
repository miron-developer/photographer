package main

import (
	"context"
	"fmt"
	"net"

	"photographer/internal/app"
	"photographer/internal/orm"
	"photographer/internal/protobuf/db"

	"google.golang.org/grpc"
)

type dbServer struct {
	db.UnimplementedDBServer
}

func newDbServer() *dbServer {
	return &dbServer{}
}

func (s *dbServer) Create(ctx context.Context, p *db.SQLInsertParams) (*db.SQLResult, error) {
	return &db.SQLResult{}, nil
}

func (s *dbServer) Update(ctx context.Context, p *db.SQLUpdateParams) (*db.SQLResult, error) {
	return &db.SQLResult{}, nil
}

func (s *dbServer) Delete(ctx context.Context, p *db.SQLDeleteParams) (*db.SQLResult, error) {
	return &db.SQLResult{}, nil
}

func (s *dbServer) Select(ctx context.Context, p *db.SQLSelectParams) (*db.SQLResult, error) {
	return &db.SQLResult{}, nil
}

func main() {
	// initialize
	log := app.CreateLogged("db")
	if e := orm.InitDB(log); e != nil {
		log.Fatalln("init db error: ", e)
	}

	// listen db service
	lis, e := net.Listen("tcp", fmt.Sprintf("localhost:%d", app.DBPort))
	if e != nil {
		log.Fatalf("failed to listen: %v\n", e)
	}

	srv := grpc.NewServer(nil)
	db.RegisterDBServer(srv, newDbServer())
	srv.Serve(lis)
}
