package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// open our file for reading
	file, _ := os.Open("ports.conf")

	// create a scanner object that reads the file we opened before
	// by default, NewScanner splits on newlines
	scanner := bufio.NewScanner(file)

	// we need to initialize an empty slice here for text in order
	// for the append function on line 25 to work
	var text []string

	// read in the file line by line, save it to text and close the file
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	// os.Args[1] parses argv
	TARGET := os.Args[1]
	fmt.Printf("The target is %s", TARGET)
	var openPorts []string
	for _, port := range text {
		fmt.Printf("\r\nCurrently scanning: %s:%s", TARGET, port)
		address := TARGET + ":" + port
		conn, err := net.Dial("tcp", address)
		if err != nil {
			//port is either filtered or closed
			continue
		}
		conn.Close()
		openPorts = append(openPorts, port)
		fmt.Printf("\r\n[*]Port %s is open!\n", port)
	}
	fmt.Printf("\n\nThe ports open for %s are %s", TARGET, openPorts)
}
