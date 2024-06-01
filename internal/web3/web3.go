package web3

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Web3Client struct {
	Client     *ethclient.Client
	RpcClient  *rpc.Client
	PrivateKey *ecdsa.PrivateKey
	Auth       *bind.TransactOpts
}

func NewWeb3Client(url, privateKeyHex, contractAddr string, chainId *big.Int) (*Web3Client, error) {
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	ethClient := ethclient.NewClient(rpcClient)

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	return &Web3Client{
		Client:     ethClient,
		RpcClient:  rpcClient,
		PrivateKey: privateKey,
		Auth:       auth,
	}, nil
}
