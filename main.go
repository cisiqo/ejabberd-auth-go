package main

import (
	"bufio"
	"encoding/binary"
	"os"
	"strings"
)

func main() {
	bioIn := bufio.NewReader(os.Stdin)
	bioOut := bufio.NewWriter(os.Stdout)

	var length uint16
	var result uint16
	var ret bool = true

	for {
		_ = binary.Read(bioIn, binary.BigEndian, &length)
		buf := make([]byte, length)
		r, err := bioIn.Read(buf)
		if r == 0 || err != nil {
			break
		}

		std := string(buf[:])
		ret = excute(std)

		length = 2
		binary.Write(bioOut, binary.BigEndian, &length)

		if ret {
			result = 1
		} else {
			result = 0
		}
		binary.Write(bioOut, binary.BigEndian, &result)
		bioOut.Flush()
	}
}

func excute(std string) bool {
	args := strings.Split(std, ":")
	switch args[0] {
	case "isuser":
		return isuser(args[1], args[2])
	case "auth":
		return auth(args[1], args[2], args[3])
	case "setpass", "tryregister", "removeuser", "removeuser3":
		return false
	}
	return false
}

func isuser(username, server string) bool {
	// TODO: Check user ejabberd user
	return false
}

func auth(username, server, password string) bool {
	// TODO: Auth your ejabberd user
	return false
}
