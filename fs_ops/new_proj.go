package fs_ops

import (
	"fmt"
	"log"
	"os"
	"path"
)

/*创建工程，包括目录结构和关键的示例文件

Args:
    projName string, 工程名称
	isLib bool,      是否为库工程
*/
func CreateProject(projName string, isLib bool) error {
	log.Printf("create project %v, isLib=%v\r\n", projName, isLib)
	// 创建工程root目录
	projRoot := fmt.Sprintf("./%v", projName)
	err := os.Mkdir(projRoot, 0777)
	if nil != err {
		return err
	}
	// 创建 src 目录
	srcDir := path.Join(projRoot, "src")
	err = os.Mkdir(srcDir, 0777)
	if nil != err {
		return err
	}
	// 创建 include 目录
	incDir := path.Join(projRoot, "include")
	err = os.Mkdir(incDir, 0777)
	if nil != err {
		return err
	}
	// 创建 test 目录
	testDir := path.Join(projRoot, "test")
	err = os.Mkdir(testDir, 0777)
	if nil != err {
		return err
	}
	// 创建 .gitignore 文件
	gitignorePath := path.Join(projRoot, ".gitignore")
	err = createGitignore(gitignorePath)
	if nil != err {
		return err
	}
	// 创建库对应的头文件和源码文件
	err = createLibSource(incDir, srcDir)
	if nil != err {
		return err
	}
	if !isLib {
		// 创建二进制工程的入口main文件
		err = createMainFile(srcDir)
		if nil != err {
			return err
		}
	}
	// 创建包管理配置文件 River.toml
	err = createTOML(projRoot, projName)
	if nil != err {
		return err
	}
	return nil
}

/* 创建 .gitignore 文件

Args:
    p string, gitignore文件的路径

Returns:
    error, 错误信息

*/
func createGitignore(p string) error {
	gitCont := `
output/
build/
*.exe
	`
	err := createAndFillFile(p, gitCont)
	if nil != err {
		return fmt.Errorf("failed to create .gitignore file, e:%v", err)
	}
	return nil
}

/* 创建 include/libs.h 和 src/libs.cpp 文件

Args:
    incDir string, include 目录，用于存放对外暴露接口的头文件
	srcSir string, src 目录，用于存放cpp源码文件和不对外暴露接口的头文件

Returns:
    error

*/
func createLibSource(incDir, srcDir string) error {
	// 创建 include/libs.h 文件
	incLibPath := path.Join(incDir, "libs.h")
	incCont := `#ifndef __LIBS_H__
#define __LIBS_H__	

int add(int a, int b);

#endif`
	err := createAndFillFile(incLibPath, incCont)
	if nil != err {
		return fmt.Errorf("failed to create include/libs.h, e:%v", err)
	}
	// 创建 src/libs.cpp 文件
	srcLibPath := path.Join(srcDir, "libs.cpp")
	srcCont := `#include "libs.h"

int add(int a, int b) {
	return a + b;
}
`
	err = createAndFillFile(srcLibPath, srcCont)
	if nil != err {
		return fmt.Errorf("failed to create src/libs.cpp, e:%v", err)
	}
	return err
}

/* 创建 main.cpp 文件

Args:
    srcDir string, 指定src目录

Returns:
    error

*/
func createMainFile(srcDir string) error {
	mainCont := `
#include <iostream>
#include "libs.h"

int main(int argc, char ** argv) {
	int a = 5, b = 7;
	int c = add(a, b);
	std::cout << "a=" << a ", b=" << b << std::endl;
	std::cout << "add(a, b)=" << c << std::endl;
	return 0;
}
`
	mainPath := path.Join(srcDir, "main.cpp")
	err := createAndFillFile(mainPath, mainCont)
	if nil != err {
		return fmt.Errorf("failed to create src/main.cpp, e:%v", err)
	}
	return nil
}

/* 创建工包管理配置文件 River.toml

Args:
    projRoot string, 工程根目录
    projName string, 工程名，也作为包名

Returns:
    error

*/
func createTOML(projRoot, projName string) error {
	tomlCont := fmt.Sprintf(`
name="%v"	
`, projName)
	tomlPath := path.Join(projRoot, "River.toml")
	err := createAndFillFile(tomlPath, tomlCont)
	if nil != err {
		return fmt.Errorf("failed to create %v, e:%v", tomlPath, err)
	}
	return nil
}
