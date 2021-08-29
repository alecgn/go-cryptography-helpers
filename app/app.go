package app

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/urfave/cli"
)

func Create() *cli.App {
	app := cli.NewApp()
	app.Name = "CryptographyHelpers"
	app.Usage = "Cryptography facilities -> encryption/decryption/hash/hmac/kdf/encode/decode"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "text",
			Value: "test string",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "hex",
			Usage: "encode string to hexadecimal",
			Flags: flags,
			Action: encodeStringToHex,
		},
		{
			Name:  "base64",
			Usage: "encode string to base64",
			Flags: flags,
			Action: encodeStringToBase64,
		},
	}

	return app
}

func encodeStringToHex(c *cli.Context) {
	str := c.String("text")
	hexStr := hex.EncodeToString([]byte(str))
	fmt.Println(hexStr)
}

func encodeStringToBase64(c *cli.Context) {
	str := c.String("text")
	base64Str := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(base64Str)
}