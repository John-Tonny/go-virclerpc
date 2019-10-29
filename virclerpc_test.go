package virclerpc

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	err := godotenv.Load()
	if err != nil && os.Getenv("GO_ENV") != "test" {
		log.Fatal("Error loading .env file")
	}
}

func TestRPCClient(t *testing.T) {
	Convey("Success", t, func() {
		basicAuth := &BasicAuth{
			Username: os.Getenv("USERNAME"),
			Password: os.Getenv("PASSWORD"),
		}
		c := NewRPCClient(os.Getenv("VIRCLED_ENDPOINT"), basicAuth)
		_, err := c.GetBlockCount()
		So(err, ShouldBeNil)
	})
	Convey("InvalidBasicAuth", t, func() {
		basicAuth := new(BasicAuth)
		c := NewRPCClient(os.Getenv("VIRCLED_ENDPOINT"), basicAuth)
		_, err := c.GetBlockCount()
		So(err, ShouldNotBeNil)
	})
}

func TestNewAddress(t *testing.T) {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("VIRCLED_ENDPOINT"), basicAuth)
	Convey("Success", t, func() {
		_, err := c.GetNewAddress("KeisukeYamashita")
		So(err, ShouldBeNil)
	})
}

func TestGetBalance(t *testing.T) {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("VIRCLED_ENDPOINT"), basicAuth)
	Convey("Success", t, func() {
		result, err := c.GetBalance("*")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, 0.00000000) // Block Height
	})

	Convey("No addresss", t, func() {
		result, err := c.GetBalance("I do not exists")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, 0.00000000) // 0.00000000
	})
}

func TestGetBlockHash(t *testing.T) {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("VIRCLED_ENDPOINT"), basicAuth)
	Convey("Success", t, func() {
		result, err := c.GetBlockHash(54000)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "dbad92afd4e5e185efcad9b7dd30e6ec3b865b194f03c520f9d71f3da6a4b90c ") // Block Hash
	})

	Convey("Invalid block index", t, func() {
		_, err := c.GetBlockHash(-1)
		So(err.Error(), ShouldEqual, "Block height out of range")
	})
}

// TODO: change for testnet*testing.T) {
// 	basicAuth := &BasicAuth{
// 		Username: os.Getenv("USERNAME"),
// 		Password: os.Getenv("PASSWORD"),
// 	}
// 	c := NewRPCClient(os.Getenv("VIRCLED_ENDPOINT"), basicAuth)

// 	Convey("Invalid block hash", t, func() {
// 		_, err := c.GetBlock("Invalid transactionId")
// 		So(err.Error(), ShouldEqual, "Block not found")
// 	})
// }

func TestGetBlockCount(t *testing.T) {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("VIRCLED_ENDPOINT"), basicAuth)
	Convey("Success", t, func() {
		_, err := c.GetBlockCount()
		So(err, ShouldBeNil)
	})
}

// TODO: change for testnet
// func TestGetRawTransactions(t *testing.T) {
// 	basicAuth := &BasicAuth{
// 		Username: os.Getenv("USERNAME"),
// 		Password: os.Getenv("PASSWORD"),
// 	}
// 	c := NewRPCClient(os.Getenv("VIRCLED_ENDPOINT"), basicAuth)
// 	Convey("Success", t, func() {
// 		result, err := c.GetRawTransactions([]string{"e78be2880e53a278442d736e44c9bdc7822255a3d072639858af8a2d7412d124","9ac157891582c9f5267c2f1d6becedabdfcce1651d42fa032bc907fa36fbfb79"}) // Genesis Block
// 		So(err, ShouldBeNil)
// 		So(result[0], ShouldEqual, "a0000000000001683e7f80aaea19b6ee966e7d59521b4ff95961e456cfe7f4d3c96a74e26f24410000000000ffffffff0201188f1939000000001976a9147d2b20f6a8c25cd7bad11ce0d0c7104b108810c088ac042e516d54636d4c345547625a4e32474a4d4c476f67676856784b706b554e6574336e787939626431756637666b353702473044022044ce6f677452afbc93d0611fdb257cf732520ee4d280d24d6d40f0b7f21a059902202bbf0598f14012648c37ea62f6dc2090c3b9a35ec52432735831510ae410b44201210342551aabf5e08d33f3fe1741714bd826820b50ef63a960c01f2ad97c3848fb68")
// 		So(result[1], ShouldEqual, "a00200000000018d7d484f4779cb392f1a4bf7f7fe3b5442ae234998552770d93567940f3729070100000000ffffffff030404f0d200000141c40b480000000017a914ff8624af288f7a7452c4a35f898f62069038a4cb87018fa14572280000001976a9144c5d2193c3542234c03668a4683ac8227a75379588ac0247304402206f014048e7a2f3348a1d25b1e5e6a72273d21e9c57404371c11a9df64699c96502201a36995e168f0b6547bc606f8ec2a1abafd2d528320b95894799bd0f85741d4f012103e5e55d666ea118d40e419ee14309b09c7a7681d65f3174df79ba393f4a33b6d4")
// 	})

// 	Convey("Invalid transactionId", t, func() {
// 		_, err := c.GetRawTransactions([]string{"Invalid transactionId"})
// 		So(err.Error(), ShouldContainSubstring, "parameter 1 must be hexadecimal string")
// 	})
// }

func TestDecodeRawTransaction(t *testing.T) {
	basicAuth := &BasicAuth{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}
	c := NewRPCClient(os.Getenv("VIRCLED_ENDPOINT"), basicAuth)
	Convey("Success", t, func() {
		result, err := c.DecodeRawTransactions([]string{"a0000000000001683e7f80aaea19b6ee966e7d59521b4ff95961e456cfe7f4d3c96a74e26f24410000000000ffffffff0201188f1939000000001976a9147d2b20f6a8c25cd7bad11ce0d0c7104b108810c088ac042e516d54636d4c345547625a4e32474a4d4c476f67676856784b706b554e6574336e787939626431756637666b353702473044022044ce6f677452afbc93d0611fdb257cf732520ee4d280d24d6d40f0b7f21a059902202bbf0598f14012648c37ea62f6dc2090c3b9a35ec52432735831510ae410b44201210342551aabf5e08d33f3fe1741714bd826820b50ef63a960c01f2ad97c3848fb68","a00200000000018d7d484f4779cb392f1a4bf7f7fe3b5442ae234998552770d93567940f3729070100000000ffffffff030404f0d200000141c40b480000000017a914ff8624af288f7a7452c4a35f898f62069038a4cb87018fa14572280000001976a9144c5d2193c3542234c03668a4683ac8227a75379588ac0247304402206f014048e7a2f3348a1d25b1e5e6a72273d21e9c57404371c11a9df64699c96502201a36995e168f0b6547bc606f8ec2a1abafd2d528320b95894799bd0f85741d4f012103e5e55d666ea118d40e419ee14309b09c7a7681d65f3174df79ba393f4a33b6d4"})
		So(err, ShouldBeNil)
		So(result[0].Txid, ShouldEqual, "e78be2880e53a278442d736e44c9bdc7822255a3d072639858af8a2d7412d124")
		So(result[1].Txid, ShouldEqual, "9ac157891582c9f5267c2f1d6becedabdfcce1651d42fa032bc907fa36fbfb79")
	})

	Convey("Invalid rawtx", t, func() {
		_, err := c.DecodeRawTransactions([]string{"Invalid rawTransaction"})
		So(err.Error(), ShouldEqual, "TX decode failed")
	})

}
