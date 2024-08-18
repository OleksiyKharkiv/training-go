package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"net"
	"net/http"
)

func main() {
	var (
		hello, bye string
		age        int = 10000000
	)
	fx.New(
		fx.Provide(NewHTTPServer),
		fx.Invoke(func(server *http.Server) {}),
	).Run()
	hello = "Hallo Go!"
	bye = "Bye, poverty! "
	fmt.Println(" %s %s %s & %d $ !!!", hello, bye, hello, age)
}
func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	srv :=
		&http.Server{Addr: ":8080"}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP Server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
