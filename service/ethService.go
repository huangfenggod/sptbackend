package service

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"sptbackend/config"
	"sptbackend/nosql"
	"sptbackend/token"
	"strconv"
)

var Conn  *ethclient.Client
var spt *token.Spt
var bslp *token.Bslp
func InitDial()  {
	dial, err := ethclient.Dial(config.Cfg.Ethereum.Network)
	if err != nil{
		log.Printf("connect eth fail because:%s", err)
	}
	Conn = dial
	spt1, err := token.NewSpt(common.HexToAddress(config.Cfg.Ethereum.SptAddress), dial)
	if err !=nil{
		log.Printf("spt wrong because: %s",err)
	}
	spt = spt1
	bslp1, err1 := token.NewBslp(common.HexToAddress(config.Cfg.Ethereum.Bnblp), dial)
	if err1!=nil {
		log.Printf("bslp wrong because: %s",err1)
	}
	bslp = bslp1
}

func bigIntToInt(bigint *big.Int) (int64,error){
	var in int64
	s := bigint.String()
	if len(s)<=10 {
		return int64(0),nil
	}
	s1 := s[0 : len(s)-10]
	atoi, err := strconv.ParseInt(s1,10,64)
	if err !=nil {
		return in,err
	}
	return atoi,nil
}
//这是把int64末尾添加10位0转成bigint
func intTobigInt(in int64) *big.Int {

	inString := strconv.FormatInt(in, 10)
	newString := inString+"0000000000"
	bigint, _ := new(big.Int).SetString(newString, 10)
	return bigint
}


func GetLpTotalSupply() (int64,error ){
	lpTotalSupply, err := spt.GetLpTotalSupply(&bind.CallOpts{
		From:        common.Address{},
		Pending:     false,
		BlockNumber: nil,
		Context:     nil})
	if err!=nil {
		log.Println(err)
		return 0,err
	}
	toInt, err := bigIntToInt(lpTotalSupply)
	if err !=nil {
		return 0,err
	}
	return toInt,nil
}

func GetAccountLpBalanceOf(account string) (int64,error) {
	lp, err := spt.GetAccountLp(&bind.CallOpts{
		From:        common.Address{},
		Pending:     false,
		BlockNumber: nil,
		Context:     nil}, common.HexToAddress(account))
	if err !=nil {
		return 0 ,err
	}
	toInt, err := bigIntToInt(lp)
	if err !=nil {
		return 0,err
	}
	return toInt,nil
}

func GetBnbSptLp(account string) (int64,error) {
	lp, err := bslp.BalanceOf(&bind.CallOpts{
		From:        common.Address{},
		Pending:     false,
		BlockNumber: nil,
		Context:     nil}, common.HexToAddress(account))
	if err !=nil {
		return 0 ,err
	}
	toInt, err := bigIntToInt(lp)
	if err !=nil {
		return 0,err
	}
	return toInt,nil
}


func GetLivid()(int64,error){
	lpdivid, err := spt.Lpdivid(&bind.CallOpts{
		From:        common.Address{},
		Pending:     false,
		BlockNumber: nil,
		Context:     nil})
	if err !=nil {
		return 0,nil
	}
	toInt, err := bigIntToInt(lpdivid)
	return toInt,err

}


func getTransactOpts(storePath string) (*bind.TransactOpts ,error) {
	privateKey, err := crypto.HexToECDSA(storePath)
	if err !=nil {
		log.Fatalln(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println(err)
		return nil,errors.New("has got treasure")
	}
	crypto.PubkeyToAddress(*publicKeyECDSA)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(config.Cfg.Ethereum.ChainId)))
	if err !=nil {
		log.Println(err)
		return nil,errors.New("privatekey wrong")
	}
	return auth ,nil
}


func Transfer(address string ,amount1 int64,storePath string) (bool ,error){
	tobigInt1 := intTobigInt(amount1)
	auth, err2 := getTransactOpts(storePath)
	if err2 !=nil{
		return false,err2
	}
	transaction, err := spt.Transfer(&bind.TransactOpts{
		From:      auth.From,
		//Nonce:     auth.Nonce,
		Signer:    auth.Signer,
		//Value:     big.NewInt(0),
		//GasPrice:  auth.GasPrice,
		//GasFeeCap: auth.GasFeeCap,
		//GasTipCap: auth.GasTipCap,
		//GasLimit:  auth.GasLimit,
		//Context:   auth.Context,
		//NoSend:    true,
	}, common.HexToAddress(address), tobigInt1)
	if err!=nil {
		return false ,err
	}
	mined, err2 := bind.WaitMined(context.Background(), Conn, transaction)
	if err2 !=nil {
		log.Printf("waitMined transfer fail  %s",err2)
		return false,err2
	}
	log.Printf("transfer success address: %s ,TxHash: %s",address,mined.TxHash)
	return true ,nil
}

//使用bind去查询是否存在
func CheckHash(tradHash string) (bool,error)  {
	_, _, err := Conn.TransactionByHash(context.Background(), common.HexToHash(tradHash))
	if err !=nil {
		return false,err
	}
	return true,nil
}
//验证hash
func IdentifyTransactionHash(transactionHash string)bool  {
	has := nosql.IsHasTxhash(transactionHash)
	fmt.Println(has)
	return true
	if has {
		return false
	}
	hash, _ := CheckHash(transactionHash)
	if  !hash{
		return false
	}else {
		nosql.InsertHashToRedis(transactionHash)
		return true
	}
}
