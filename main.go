package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	isLibProj bool = false
)

func main() {
	fmt.Println("hello river")

	app := cli.NewApp()

	app.Commands = []*cli.Command{
		// 创建工程的命令
		{
			Name:  "new",
			Usage: "create a c++ project",
			Flags: []cli.Flag{
				// -lib
				&cli.BoolFlag{
					Name:        "lib",
					Value:       false,
					Required:    false,
					Destination: &isLibProj,
					Usage:       "new创建工程时指定library类型的工程"},
			},
			Action: func(ctx *cli.Context) error {
				fmt.Printf("create c++ project, isLibProj=%v\r\n", isLibProj)
				return nil
			},
		},
		// 编译工程
		// 运行工程
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
