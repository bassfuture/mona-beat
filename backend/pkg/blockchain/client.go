package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BlockchainClient 区块链客户端
type BlockchainClient struct {
	client      *ethclient.Client
	privateKey  *ecdsa.PrivateKey
	auth        *bind.TransactOpts
	contract    *CaptureNFT
	contractAddr common.Address
}

// Config 区块链配置
type Config struct {
	RPCEndpoint     string
	PrivateKey      string
	ContractAddress string
	ChainID         int64
}

// NewBlockchainClient 创建新的区块链客户端
func NewBlockchainClient(config Config) (*BlockchainClient, error) {
	// 连接到区块链网络
	client, err := ethclient.Dial(config.RPCEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to blockchain: %v", err)
	}

	// 解析私钥
	privateKeyHex := config.PrivateKey
	if strings.HasPrefix(privateKeyHex, "0x") {
		privateKeyHex = privateKeyHex[2:]
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// 设置链ID
	chainID := big.NewInt(config.ChainID)

	// 创建交易授权
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	// 设置Gas限制
	auth.GasLimit = uint64(3000000) // 3M gas limit
	auth.GasPrice = big.NewInt(20000000000) // 20 gwei

	// 连接到合约
	contractAddress := common.HexToAddress(config.ContractAddress)
	contract, err := NewCaptureNFT(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to contract: %v", err)
	}

	return &BlockchainClient{
		client:       client,
		privateKey:   privateKey,
		auth:         auth,
		contract:     contract,
		contractAddr: contractAddress,
	}, nil
}

// GetDefaultConfig 获取默认配置
func GetDefaultConfig() Config {
	return Config{
		RPCEndpoint:     getEnv("RPC_ENDPOINT", "https://testnet-rpc.monad.xyz"),
		PrivateKey:      getEnv("PRIVATE_KEY", ""),
		ContractAddress: getEnv("CONTRACT_ADDRESS", "0x572d38104998De89fA161D8002A727d6a1FFd9c2"),
		ChainID:         getEnvInt64("CHAIN_ID", 10143), // 正确的链ID
	}
}



// getEnv 获取环境变量
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt64 获取int64类型的环境变量
func getEnvInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := fmt.Sscanf(value, "%d", &defaultValue); err == nil && intValue == 1 {
			return defaultValue
		}
	}
	return defaultValue
}

// Close 关闭客户端连接
func (bc *BlockchainClient) Close() {
	if bc.client != nil {
		bc.client.Close()
	}
}

// GetClient 获取以太坊客户端
func (bc *BlockchainClient) GetClient() *ethclient.Client {
	return bc.client
}

// GetContract 获取合约实例
func (bc *BlockchainClient) GetContract() *CaptureNFT {
	return bc.contract
}

// GetAuth 获取交易授权
func (bc *BlockchainClient) GetAuth() *bind.TransactOpts {
	return bc.auth
}

// UpdateGasPrice 更新Gas价格
func (bc *BlockchainClient) UpdateGasPrice(ctx context.Context) error {
	gasPrice, err := bc.client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("failed to get gas price: %v", err)
	}
	
	// 增加20%的Gas价格以确保交易被快速确认
	gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(120))
	gasPrice = new(big.Int).Div(gasPrice, big.NewInt(100))
	
	bc.auth.GasPrice = gasPrice
	log.Printf("Updated gas price to: %s wei", gasPrice.String())
	
	return nil
}

// GetBalance 获取账户余额
func (bc *BlockchainClient) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	return bc.client.BalanceAt(ctx, address, nil)
}

// GetNonce 获取账户nonce
func (bc *BlockchainClient) GetNonce(ctx context.Context) (uint64, error) {
	publicKey := bc.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return 0, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return bc.client.PendingNonceAt(ctx, fromAddress)
}