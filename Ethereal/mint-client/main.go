package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"zk/mint/mint"
	"zk/mint/utils"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	SRS kzg.SRS
)

const (
	SRS_SIZE = 17
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func loadSRSFromContract(mintContract *mint.MintCaller) *kzg.SRS {
	SRS := kzg.SRS{
		Pk: kzg.ProvingKey{
			G1: make([]bn254.G1Affine, SRS_SIZE),
		},
		Vk: kzg.VerifyingKey{
			G2: [2]bn254.G2Affine{},
			G1: bn254.G1Affine{},
		},
	}
	for i := 0; i < SRS_SIZE; i++ {
		g1_x, err := mintContract.SRSG1X(nil, big.NewInt(int64(i)))
		panicErr(err)
		SRS.Pk.G1[i].X.SetBigInt(g1_x)
		g1_y, err := mintContract.SRSG1Y(nil, big.NewInt(int64(i)))
		panicErr(err)
		SRS.Pk.G1[i].Y.SetBigInt(g1_y)
	}
	SRS.Vk.G1.X.Set(&SRS.Pk.G1[0].X)
	SRS.Vk.G1.Y.Set(&SRS.Pk.G1[0].Y)

	zero := big.NewInt(0)
	one := big.NewInt(1)
	g2_x0_0, err := mintContract.SRSG2X0(nil, zero)
	panicErr(err)
	SRS.Vk.G2[0].X.A0.SetBigInt(g2_x0_0)
	g2_x0_1, err := mintContract.SRSG2X0(nil, one)
	panicErr(err)
	SRS.Vk.G2[0].X.A1.SetBigInt(g2_x0_1)

	g2_x1_0, err := mintContract.SRSG2X1(nil, zero)
	panicErr(err)
	SRS.Vk.G2[1].X.A0.SetBigInt(g2_x1_0)
	g2_x1_1, err := mintContract.SRSG2X1(nil, one)
	panicErr(err)
	SRS.Vk.G2[1].X.A1.SetBigInt(g2_x1_1)

	g2_y0_0, err := mintContract.SRSG2Y0(nil, zero)
	panicErr(err)
	SRS.Vk.G2[0].Y.A0.SetBigInt(g2_y0_0)
	g2_y0_1, err := mintContract.SRSG2Y0(nil, one)
	panicErr(err)
	SRS.Vk.G2[0].Y.A1.SetBigInt(g2_y0_1)

	g2_y1_0, err := mintContract.SRSG2Y1(nil, zero)
	panicErr(err)
	SRS.Vk.G2[1].Y.A0.SetBigInt(g2_y1_0)
	g2_y1_1, err := mintContract.SRSG2Y1(nil, one)
	panicErr(err)
	SRS.Vk.G2[1].Y.A1.SetBigInt(g2_y1_1)
	return &SRS
}

func G1AffineToG1Point(p *bn254.G1Affine) *mint.PairingG1Point {
	var X, Y big.Int
	p.X.BigInt(&X)
	p.Y.BigInt(&Y)
	return &mint.PairingG1Point{
		X: &X,
		Y: &Y,
	}
}

func ordinal(n int) string {
	suffix := []string{"th", "st", "nd", "rd"}
	if n%100 >= 11 && n%100 <= 13 {
		return strconv.Itoa(n) + "th"
	}
	return strconv.Itoa(n) + suffix[n%10%4]
}

func main() {
	var action = flag.String("action", "register", "action to perform")
	var contractAddr = flag.String("contract", "", "contract address")
	var rpcAddr = flag.String("rpc", "http://47.76.89.7:8545", "rpc address")
	// var rpcAddr = flag.String("rpc", "http://127.0.0.1:8545", "rpc address")
	// --------------------------------------------------
	// Note: Please change the account and pk to your own
	var accountAddr = flag.String("account", "0x78B8Af83b0DEF7e7f89Cd964f9E3921F9685844b", "account address")
	var privateKey = flag.String("pk", "e5e8bb4e61f5afe8bf82afec48adf1a5ddd799f042b1c71d3fe5cf26f7f00c31", "private key")

	// var accountAddr = flag.String("account", "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", "account address")
	// var privateKey = flag.String("pk", "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "private key")

	// --------------------------------------------------
	flag.Parse()
	contractAddress := common.HexToAddress(*contractAddr)
	accountAddress := common.HexToAddress(*accountAddr)
	ethclient, err := ethclient.Dial(*rpcAddr)
	transactor := utils.NewTransactor(ethclient, *privateKey)

	mintContract, err := mint.NewMint(contractAddress, ethclient)
	panicErr(err)

	SRS := loadSRSFromContract(&mintContract.MintCaller)

	switch *action {
	case "register":
		var sword []fr.Element
		f, err := os.Open("sword.json")
		panicErr(err)
		defer f.Close()
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&sword)
		panicErr(err)

		bladeCommitment, bladeProof := CraftBladeSignature(sword, SRS)
		commitmentPoint := G1AffineToG1Point(&bladeCommitment)
		bladeProofPoint := G1AffineToG1Point(&bladeProof.H)
		// soulBox := G1AffineToG1Point(&bladeProof.PublicKeyG1Aff)
		publicKeyG2 := new(bn254.G1Affine)
		soulBox := G1AffineToG1Point(publicKeyG2)
		fmt.Println("PrivateKey", bladeProof.PrivateKey.String())
		mintContract.Register(transactor, *commitmentPoint, *bladeProofPoint, *soulBox)

		transactor.Value = big.NewInt(0)
	case "replay":
		mintContract.Replay(transactor)
	case "mint":
		var sword []fr.Element
		f, err := os.Open("sword.json")
		panicErr(err)
		defer f.Close()
		decoder := json.NewDecoder(f)
		err = decoder.Decode(&sword)
		panicErr(err)

		nonce, err := mintContract.GetNonce(nil, transactor.From)
		panicErr(err)
		nonceFr := new(fr.Element)
		nonceFr.SetBigInt(nonce)
		fmt.Printf("Minting the %s gem\n", ordinal(int(nonce.Int64())))
		proof, err := kzg.Open(sword, *nonceFr, SRS.Pk)
		panicErr(err)
		proofPoint := G1AffineToG1Point(&proof.H)

		var value big.Int
		proof.ClaimedValue.BigInt(&value)
		// fmt.Println("transactor", transactor)
		// fmt.Println("proofPoint", proofPoint)
		// fmt.Println("value", value.String())

		mintContract.Mint(transactor, *proofPoint, &value)
	case "query":
		accountDeposit, err := mintContract.Deposits(nil, accountAddress)
		panicErr(err)
		collectedY, err := mintContract.GetCollectedY(nil, accountAddress)
		panicErr(err)
		spew.Dump(collectedY, accountDeposit)
	}
}
