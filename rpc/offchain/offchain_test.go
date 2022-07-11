package offchain

import (
	"os"
	"testing"

	"github.com/gaoqinying/substrate-rpc/client"
	"github.com/gaoqinying/substrate-rpc/config"
)

var offchain *Offchain

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	offchain = NewOffchain(cl)
	os.Exit(m.Run())
}
