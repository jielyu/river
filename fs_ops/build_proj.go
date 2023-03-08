package fs_ops

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	// "path/filepath"
	"strings"

	"github.com/jielyu/river/models"
)

func BuildProject(cConfig *models.CommandConfig, tConfig *models.TOMLConfig) error {
	//执行编译
	fmt.Printf("release: %v\r\n", cConfig.Release)
	// TODO: 读取依赖关系
	// TODO: 逐个递归编译依赖

	// 读取当前项目源码文件
	files, err := findAllFiles("src/")
	if nil != err {
		return fmt.Errorf("failed to find source file from src/. error:%v", err)
	}
	fmt.Printf("files=%v\r\n", files)

	// 编译当前项目源码文件
	var sourceFiles []string
	var objFiles []string
	for _, file := range files {
		// 去掉main文件
		if file == "src/main.cpp" {
			continue
		}
		// 去掉头文件
		extName := path.Ext(file)
		if extName == ".h" || extName == ".hpp" {
			continue
		}
		// 构造obj文件路径
		objPath := strings.TrimSuffix(file, path.Ext(file))
		objPath = path.Join("build", fmt.Sprintf("%s.o", objPath))
		objDir, _ := filepath.Split(objPath)
		if !isDir(objDir) {
			os.MkdirAll(objDir, 0777)
		}
		// 记录源码文件和obj文件
		sourceFiles = append(sourceFiles, file)
		objFiles = append(objFiles, objPath)
	}
	for idx, sourceFile := range sourceFiles {
		objFile := objFiles[idx]
		c := exec.Command("g++", "-c", sourceFile, "-o", objFile, "-Iinclude")
		output, err := c.Output()
		if err != nil {
			fmt.Println(string(output))
			return fmt.Errorf("failed to compile %s. e:%v", sourceFile, err)
		}
		fmt.Printf("%v: build success\r\n", sourceFile)

	}
	// 链接当前项目的目标库
	objFileArg := strings.Join(objFiles, " ")
	libName := fmt.Sprintf("lib%s.a", tConfig.Name)
	libPath := path.Join("build", libName)
	arCmd := exec.Command("ar", "-rc", libPath, objFileArg)
	output, err := arCmd.Output()
	if err != nil {
		fmt.Printf("failed to achieve %s\r\n", libPath)
		return err
	}
	fmt.Println(string(output))
	fmt.Printf("generate %s success, save to: %s\r\n", libName, libPath)
	// 对于二进制项目编译main.cpp
	if tConfig.ProjectType != "lib" {
		if !isFile("src/main.cpp") {
			return fmt.Errorf("absense of src/main.cpp for binary project")
		}
		mainCmd := exec.Command("g++", "-c", "src/main.cpp", "-o", "build/src/main.o", "-Iinclude")
		output, err = mainCmd.Output()
		if err != nil {
			fmt.Printf("failed to compile src/main.cpp\r\n%v", string(output))
			return err
		}
		fmt.Printf("src/main.cpp: build success\r\n")
		// 链接生成可执行程序
		libArgs := fmt.Sprintf("-l%s", tConfig.Name)
		targetPath := path.Join("build", tConfig.Name)
		mainLinkCmd := exec.Command("g++", "build/src/main.o", "-o", targetPath, "-Lbuild", libArgs)
		output, err = mainLinkCmd.Output()
		if err != nil {
			fmt.Printf("failed to link src/main.cpp\r\n%v", string(output))
			return err
		}
		fmt.Printf("src/main.cpp: link success\r\n")
	}

	return nil
}
