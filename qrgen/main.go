package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mdp/qrterminal/v3"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	// 改行を削除
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	// 改行を削除
	password = strings.TrimSpace(password)

	// URLの生成
	url := fmt.Sprintf("https://%s:%s@example.com", username, password)

	// QRコードをターミナルに表示
	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig(url, config)

	fmt.Println("QR code displayed in terminal")
}
