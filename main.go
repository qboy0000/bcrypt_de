package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"strings"
)

// 加密密码
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {

	}
	return string(hash)
}

// 验证密码
func ValidatePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

func main() {
	//for i, args := range os.Args {
	//	fmt.Printf("args[%d]=%s\n", i, args)
	//}
	pwd := os.Args[1]
	fmt.Printf("decrypt password :%s\n", pwd)
	dicFile := os.Args[2]
	f, err := os.OpenFile(dicFile, os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		lineByte, err := r.ReadString('\n')
		line := strings.TrimSpace(lineByte)
		if err != nil && err != io.EOF {
			break
		}
		if err == io.EOF {
			break
		}
		result := ValidatePasswords(pwd, []byte(line))
		if result {
			fmt.Printf("%s success decrypt : %s \n", pwd, line)
		}
	}
}
