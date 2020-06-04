package server

import (
	"crypto/tls"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/lumeng689/go-micro/pkg/ui/data/swagger"
	"github.com/lumeng689/go-micro/pkg/util"
	pb "github.com/lumeng689/go-micro/proto/demo"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	ServerPort     string
	SecureServer   bool
	CertServerName string
	CertPemPath    string
	CertKeyPath    string
	SwaggerDir     string
	EndPoint       string

	tlsConfig *tls.Config
)

func Run() (err error) {
	initLog()
	EndPoint = ":" + ServerPort
	conn, err := net.Listen("tcp", EndPoint)
	if err != nil {
		log.Printf("TCP Listen err:%v\n", err)
	}

	// 判断是否启用证书保护系统
	if SecureServer {
		srv := newSecureServer(conn)
		log.Printf("gRPC and https listen on: %s\n", ServerPort)
		tlsConfig = util.GetTLSConfig(CertPemPath, CertKeyPath)
		if err = srv.Serve(util.NewTLSListener(conn, tlsConfig)); err != nil {
			log.Printf("Use SSL, Listen And Serve: %v\n", err)
		}
	} else {
		// 不使用证书
		srv := newUnSecureServer(conn)
		if err = srv.Serve(conn); err != nil {
			log.Printf("Listen And Serve: %v\n", err)
		}
	}
	return err
}

// 初始化日志
func initLog() {
	fmt.Println("init logrus log")
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func newSecureServer(conn net.Listener) (*http.Server) {
	grpcServer := newGrpc()
	gwmux, err := newGateway()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	//serveSwaggerUI(mux)

	return &http.Server{
		Addr:      EndPoint,
		Handler:   util.GrpcSecureHandlerFunc(grpcServer, mux),
		TLSConfig: tlsConfig,
	}
}

func newGrpc() *grpc.Server {
	creds, err := credentials.NewServerTLSFromFile(CertPemPath, CertKeyPath)
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(creds),
	}

	server := grpc.NewServer(opts...)

	pb.RegisterGreeterServer(server, &HelloworldServer{})

	return server
}

func newGateway() (http.Handler, error) {
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(CertPemPath, CertServerName)
	if err != nil {
		return nil, err
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterGreeterHandlerFromEndpoint(ctx, gwmux, EndPoint, dopts); err != nil {
		return nil, err
	}

	return gwmux, nil
}

func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if ! strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(SwaggerDir, p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

// 非安全方式启动
func newUnSecureServer(conn net.Listener) (*http.Server) {
	grpcServer := newGrpc()
	gwmux, err := newGateway()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger/", serveSwaggerFile)
	serveSwaggerUI(mux)

	server := &http.Server{
		Addr:    EndPoint,
		Handler: util.GrpcUnSecureHandlerFunc(grpcServer, mux),
	}

	return server
}

func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
