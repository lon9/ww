package main

import (
	"context"
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/google/subcommands"
	"github.com/lon9/ww/client"
	"github.com/lon9/ww/server"
)

type serverCmd struct {
	port string
}

func (sc *serverCmd) Name() string     { return "server" }
func (sc *serverCmd) Synopsis() string { return "Run server." }
func (sc *serverCmd) Usage() string {
	return `server -p <Port number>:
	Run server.
`
}
func (sc *serverCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&sc.port, "p", "9999", "Port number")
}
func (sc *serverCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	server := server.NewTestServer()
	server.Run(sc.port)
	return subcommands.ExitSuccess
}

type clientCmd struct {
	addr string
	port string
}

func (cc *clientCmd) Name() string     { return "client" }
func (cc *clientCmd) Synopsis() string { return "Run client." }
func (cc *clientCmd) Usage() string {
	return `client -h <Host name or ip>:
	Run client
`
}
func (cc *clientCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cc.addr, "h", "localhost", "host name or ip")
	f.StringVar(&cc.port, "p", "9999", "port number")
}
func (cc *clientCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	client := client.NewClient()
	client.Run(cc.addr, cc.port)
	return subcommands.ExitSuccess
}

func main() {
	rand.Seed(time.Now().UnixNano())
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(new(serverCmd), "")
	subcommands.Register(new(clientCmd), "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))

}
