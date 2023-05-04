package app

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/joopinho/rp-scarper/api/proto/v1/organization"
	"github.com/joopinho/rp-scarper/configs"
	"github.com/joopinho/rp-scarper/internal/profile/remote"
	"github.com/joopinho/rp-scarper/internal/services"
	gp "github.com/joopinho/rp-scarper/internal/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EnrichApplication struct {
	config    *configs.ServerConfig
	enrichers []remote.RemoteProfileEnricher
}

func (app *EnrichApplication) Init() {
	app.enrichers = append(app.enrichers,
		remote.NewApiProfileEnricher(&app.config.RemoteApiProfile),
		remote.NewPageProfileEnricher(&app.config.RemotePageProfile))
}

func NewEnrichApplication(config *configs.ServerConfig) *EnrichApplication {

	app := &EnrichApplication{config: config}
	app.Init()

	return app
}

func (app *EnrichApplication) Serve() {
	go app.runRest()
	app.runGrpc()
}

func (app *EnrichApplication) runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterOrganizationServiceHandlerFromEndpoint(ctx, mux,
		":"+app.config.GrpcServer.Port, opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at %s", app.config.RestServer.Port)
	if err := http.ListenAndServe(":"+app.config.RestServer.Port, mux); err != nil {
		panic(err)
	}
}

func (app *EnrichApplication) runGrpc() {
	lis, err := net.Listen("tcp", ":"+app.config.GrpcServer.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	organizationSvc := &services.OrganizationService{Enrichers: &app.enrichers}
	pb.RegisterOrganizationServiceServer(s, &gp.OrganizationGrpcServer{Svc: organizationSvc})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
