package main

import (
	"context"
	"fmt"
	"net"

	"photographer/internal/app"
	"photographer/internal/db"
	pdb "photographer/internal/protobuf/db"

	"google.golang.org/grpc"
)

type dbServer struct {
	pdb.UnimplementedDBServer
}

func newDbServer() *dbServer {
	return &dbServer{}
}

func (s *dbServer) Create(ctx context.Context, p *pdb.SQLInsertParams) (*pdb.SQLResult, error) {
	return &pdb.SQLResult{}, nil
}

func (s *dbServer) Update(ctx context.Context, p *pdb.SQLUpdateParams) (*pdb.SQLResult, error) {
	return &pdb.SQLResult{}, nil
}

func (s *dbServer) Delete(ctx context.Context, p *pdb.SQLDeleteParams) (*pdb.SQLResult, error) {
	return &pdb.SQLResult{}, nil
}

func (s *dbServer) Select(ctx context.Context, p *pdb.SQLSelectParams) (*pdb.SQLResult, error) {
	return &pdb.SQLResult{}, nil
}

func main() {
	// initialize
	log := app.CreateLogger("db", "", "")
	if e := db.InitDB(log); e != nil {
		log.Fatalln("init db error: ", e)
	}

	// listen db service
	lis, e := net.Listen("tcp", fmt.Sprintf("localhost:%d", app.DBPort))
	if e != nil {
		log.Fatalf("failed to listen: %v\n", e)
	}

	srv := grpc.NewServer(nil)
	pdb.RegisterDBServer(srv, newDbServer())
	srv.Serve(lis)
}
