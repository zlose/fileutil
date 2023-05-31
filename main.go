package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	pathFile, err := os.OpenFile("path.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer pathFile.Close()

	//定义一个scanner
	scanner := bufio.NewScanner(pathFile)

	//循环获取每行内容
	for scanner.Scan() {
		line := scanner.Text()
		paths := strings.Split(line, ",")
		if len(paths) < 2 {
			log.Printf("%s\t路径缺失或分割错误\n", paths)
		}

		// 使用切片定义源文件路径和目标文件路径
		sourceFilePath := paths[0]
		destFilePath := paths[1]

		// 调用重命名函数
		err := renameFileName(destFilePath)
		if err != nil {
			return
		}

		// 调用复制文件函数
		err = copyFile(sourceFilePath, destFilePath)
		if err != nil {
			return
		}
	}
}

func renameFileName(filepath string) (err error) {

	// 获取当前时间,并按照指定格式输出
	currentTime := time.Now().Format("060102")

	// 设置文件名
	fileName := fmt.Sprintf(filepath + ".bak" + currentTime)

	// 进行重命名操作
	err = os.Rename(filepath, fileName)
	if err != nil {
		log.Printf("%s\t备份失败: %v\n", filepath, err)
	} else {
		log.Printf("%s\t备份成功,备份名称为: %s\n", filepath, fileName)
	}
	return nil
}

func copyFile(sourceFilePath, destFilePath string) (err error) {

	// 打开源文件
	src, err := os.Open(sourceFilePath)
	if err != nil {
		log.Printf("%s\t打开失败: %v\n", sourceFilePath, err)
	}
	defer src.Close()

	// 创建目标文件,先进行更名操作,更名后目标文件名已经改变,这里使用create进行创建文件
	dst, err := os.Create(destFilePath)
	if err != nil {
		log.Printf("%s\t创建失败: %v\n", destFilePath, err)
	}
	defer dst.Close()

	// 进行复制操作,将源文件内容传递至目标文件
	_, err = io.Copy(dst, src)
	if err != nil {
		log.Printf("%s\t文件写入失败: %v\n", destFilePath, err)
	} else {
		log.Printf("%s\t文件写入成功", destFilePath)
	}
	return nil
}
