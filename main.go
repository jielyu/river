package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jielyu/river/fs_ops"
	"github.com/jielyu/river/models"
	"github.com/urfave/cli/v2"
)

var (
	isLibProj bool = false
	isRelease bool = false
)

func main() {
	fmt.Println("********** Welcome to River C++ Package Tool **********")

	app := cli.NewApp()

	// 配置子命令
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
			Action: NewCommand,
		},
		// 编译工程
		{
			Name:  "build",
			Usage: "build target for c++ project",
			Flags: []cli.Flag{
				// -release
				&cli.BoolFlag{
					Name:        "release",
					Value:       false,
					Required:    false,
					Destination: &isRelease,
					Usage:       "生成release模式的目标"},
			},
			Action: BuildCommand,
		},
		// 运行工程
		{
			Name:  "run",
			Usage: "run executable target",
			Flags: []cli.Flag{
				// -release
				&cli.BoolFlag{
					Name:        "release",
					Value:       false,
					Required:    false,
					Destination: &isRelease,
					Usage:       "运行release模式的目标"},
			},
			Action: RunCommand,
		},
		// 测试工程
		{
			Name:   "test",
			Usage:  "run testcases",
			Flags:  []cli.Flag{},
			Action: RunTestcases,
		},
	}
	// 运行
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func NewCommand(ctx *cli.Context) error {
	// 解析工程名称
	if ctx.Args().Len() == 0 {
		return fmt.Errorf("please specify project name")
	}
	var projName = ctx.Args().Get(0)
	// 创建工程
	fmt.Printf("start to create project %v ...\r\n", projName)
	err := fs_ops.CreateProject(projName, isLibProj)
	if nil == err {
		packageType := "binary"
		if isLibProj {
			packageType = "library"
		}
		fmt.Printf("create %v project %v success\r\n", packageType, projName)
	}
	return err
}

func BuildCommand(ctx *cli.Context) error {
	cConfig := models.CommandConfig{Release: isRelease}
	modeName := "Debug"
	if isRelease {
		modeName = "Release"
	}
	// 载入toml配置

	// 检查是否需要编译

	// 进行编译
	fmt.Println("---------------------------------------------")
	fmt.Printf("start to build project with %v mode ...\r\n", modeName)
	err := fs_ops.BuildProject(cConfig)
	if nil != err {
		return fmt.Errorf("failed to build, e: %v", err)
	}
	fmt.Printf("build project success\r\n")
	fmt.Println("---------------------------------------------")
	return nil
}

func RunCommand(ctx *cli.Context) error {
	cConfig := models.CommandConfig{Release: isRelease}
	// 载入toml配置

	// 检查是否需要先编译
	err := BuildCommand(ctx)
	if nil != err {
		return err
	}

	// 运行工程
	err = fs_ops.RunProject(cConfig)
	if nil != err {
		return fmt.Errorf("failed to run, e: %v", err)
	}
	fmt.Println("run target success")
	return nil
}

func RunTestcases(ctx *cli.Context) error { return nil }
