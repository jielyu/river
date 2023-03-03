package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jielyu/river/fs_ops"
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
				// 解析工程名称
				if ctx.Args().Len() == 0 {
					return fmt.Errorf("please specify project name after new command")
				}
				var projName = ctx.Args().Get(0)
				// 创建工程
				err := fs_ops.CreateProject(projName, isLibProj)
				if nil == err {
					packageType := "binary"
					if isLibProj {
						packageType = "library"
					}
					fmt.Printf("create %v project %v success\r\n", packageType, projName)
				}
				return err
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
