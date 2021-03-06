package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const (
	a1 = 127
	a2 = 0
	a3 = 0
	a4 = 1
)

var user string

type close struct{}

func pwd() string {
	curDir, err := os.Getwd()
	//fmt.Println(curDir)
	if err != nil {
		log.Fatalln(err)
	}
	return curDir
}

func passAuth(pass string, conn net.Conn) bool {
	fmt.Println("in pathAuth")
	fmt.Println(pass)
	fmt.Fprintf(conn, "%d User %s logged in.\n", 230, user)
	return true
}

func userAuth(commandSc *bufio.Scanner, conn net.Conn) bool {
	if commandSc.Scan() {
		user = commandSc.Text()
		fmt.Println(user)
		if user != "anonymous" {
			fmt.Fprintln(conn, 530)
			return false
		}
		fmt.Fprintf(conn, "%d Please specify the password.\n", 331)
		return true
	}
	return false
}

func cd(commandSc *bufio.Scanner, conn net.Conn) {
	if commandSc.Scan() {
		path := commandSc.Text()
		fmt.Println(path)
		err := os.Chdir(path)
		fmt.Println(err)
		if err != nil {
			fmt.Fprintf(conn, "%d Failed to change directory.\n", 550)
			return
		}
		fmt.Fprintf(conn, "%d Directory successfully changed.\n", 250)
	}

}

func get(commandSc *bufio.Scanner, conn net.Conn, msg chan interface{}) {
	if commandSc.Scan() {
		path := commandSc.Text()
		fmt.Println(path)
		fp, err := os.Open(path)
		fmt.Println(err)
		if err != nil {
			fmt.Fprintf(conn, "%d Failed to open file(%s).\n", 550, err.Error())
			return
		}
		fmt.Fprintf(conn, "%d Get %s.----------->\n", 150, path)
		sc := bufio.NewScanner(fp)
		for sc.Scan() {
			msg <- fmt.Sprintf("%s\r\n", sc.Text())
		}
		msg <- close{}
		fmt.Fprintf(conn, "%d Closing data connection.\n", 226)
	}

}

func listenDataConn(port int) (net.Listener, error) {
	dataConnURL := fmt.Sprintf("localhost:%d", port)
	return net.Listen("tcp", dataConnURL)
}

func wordsScanner(s string) *bufio.Scanner {
	commandSc := bufio.NewScanner(strings.NewReader(s))
	commandSc.Split(bufio.ScanWords)
	return commandSc
}

func listName(entry os.FileInfo) interface{} {
	return fmt.Sprintf("%s\r\n", entry.Name())
}

func list(conn net.Conn, msg chan interface{}, elist func(os.FileInfo) interface{}) {
	entries, err := ioutil.ReadDir(pwd())
	if err != nil {
		fmt.Fprintln(conn, 450)
	}
	fmt.Fprintf(conn, "%d Here comes the directory listing.\n", 150)

	for _, entry := range entries {
		fmt.Println(elist(entry))
		msg <- elist(entry)
	}
	msg <- close{}
	fmt.Fprintf(conn, "%d Closing data connection.\n", 226)
}

//!+handleConn
func handleConn(conn net.Conn) {
	msg := make(chan interface{})

	//fmt.Println("start scan")
	fmt.Fprintf(conn, "%d Golang FTP server Ready.\n", 220)
	input := bufio.NewScanner(conn)

ftpCon:
	for input.Scan() {
		s := input.Text()
		fmt.Println(s)
		fmt.Println([]byte(s))
		commandSc := wordsScanner(s)
		commandSc.Scan()
		command := commandSc.Text()

		fmt.Println(command)
		switch command {
		case "USER":
			if !userAuth(commandSc, conn) {
				break ftpCon
			}
		case "PASS":
			if !passAuth(s[4:], conn) {
				break ftpCon
			}
		case "XPWD":
			//fmt.Println("PWDda")
			fmt.Fprintf(conn, "%d \"%s\" \n", 257, pwd())
			//fmt.Fprintf(conn, "%d\n", 250)
		case "CWD":
			cd(commandSc, conn)
		case "NLST":
			list(conn, msg, listName)
		case "RETR":
			get(commandSc, conn, msg)
		case "QUIT":
			fmt.Fprintf(conn, "%d See you again!\n", 221)
			break ftpCon
		default:
			fmt.Println("default")
			fmt.Fprintln(conn, 500)
		}
	}
	//io.Copy(os.Stdout, conn) // NOTE: ignoring errors
	time.Sleep(2 * time.Second)
	//done <- struct{}{} // signal the main goroutine
	log.Println("done")

	// NOTE: ignoring potential errors from input.Err()

	conn.Close()
}

func main() {
	control, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := control.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
