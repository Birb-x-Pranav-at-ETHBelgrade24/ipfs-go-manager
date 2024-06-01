package web3

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/CreedsCode/ipfs-go-manager/internal/web3/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Web3Client struct holds the Ethereum client and authentication details
type Web3Client struct {
	Client     *ethclient.Client
	RpcClient  *rpc.Client
	PrivateKey *ecdsa.PrivateKey
	Auth       *bind.TransactOpts
	Contract   *contract.Contract // Replace with your contract binding type
	Address    common.Address
}

// NewWeb3Client creates a new instance of Web3Client
func NewWeb3Client(rpcUrl, privateKeyHex, _contractAddress string, chainId *big.Int) (*Web3Client, error) {
	rpcClient, err := rpc.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC client: %w", err)
	}

	ethClient := ethclient.NewClient(rpcClient)

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to nonce. %s", err.Error())
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)

	fmt.Println(auth.GasPrice, "GAS PRICE")
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	contractAddress := common.HexToAddress(_contractAddress)
	contract, err := contract.NewContract(contractAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to bind to contract: %w", err)
	}

	return &Web3Client{
		Client:     ethClient,
		RpcClient:  rpcClient,
		PrivateKey: privateKey,
		Auth:       auth,
		Contract:   contract,
		Address:    contractAddress,
	}, nil
}

// DeductQuota deducts the quota of a user on the blockchain
func (wc *Web3Client) DeductQuota(userAddress string, amount int64) (string, error) {
	if err := wc.updateAuthOptions(); err != nil {
		return "", fmt.Errorf("failed to update auth options: %s", err.Error())
	}

	userAddr := common.HexToAddress(userAddress)
	tx, err := wc.Contract.UseQuota(wc.Auth, userAddr, big.NewInt(amount))
	if err != nil {
		return "", fmt.Errorf("failed to deduct quota: %s", err.Error())
	}
	return tx.Hash().Hex(), nil
}

// GetQuota retrieves the current quota of a user from the blockchain
func (wc *Web3Client) GetQuota(userAddress string) (*big.Int, error) {
	userAddr := common.HexToAddress(userAddress)
	quota, err := wc.Contract.GetQuota(&bind.CallOpts{}, userAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get quota: %w", err)
	}
	return quota, nil
}

func (wc *Web3Client) updateAuthOptions() error {
	fromAddress := crypto.PubkeyToAddress(*wc.PrivateKey.Public().(*ecdsa.PublicKey))

	nonce, err := wc.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %s", err.Error())
	}

	gasPrice, err := wc.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %s", err.Error())
	}

	wc.Auth.Nonce = big.NewInt(int64(nonce))
	wc.Auth.GasPrice = gasPrice

	return nil
}
