package ipfs

import (
	"context"
	"fmt"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"
	ma "github.com/multiformats/go-multiaddr"
)

func ConnectLocalNode(ipfsHostAddress string) (*rpc.HttpApi, error) {
	addr, err := ma.NewMultiaddr(ipfsHostAddress)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	node, err := rpc.NewApi(addr)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return node, nil
}

func PinCID(toPinCid string, node *rpc.HttpApi) error {
	ctx := context.Background()

	path, err := cidStringToPath(toPinCid)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = node.Pin().Add(ctx, path)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("CID pinned successfully")
	return nil
}

func cidStringToPath(toPinCid string) (path.ImmutablePath, error) {
	cid, err := cid.Parse(toPinCid)
	if err != nil {
		fmt.Printf(err.Error())
		return path.ImmutablePath{}, err
	}

	path := path.FromCid(cid)

	return path, nil
}
