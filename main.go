package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Usage:\nsudo gmdump <pid> <offset> <kb>")
		os.Exit(0)
	}

	mempath := fmt.Sprintf("/proc/%s/mem", os.Args[1])
	f, err := os.Open(mempath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	var offset int64
	var counter int

	if len(os.Args) > 2 {
		n, err := strconv.ParseInt(os.Args[2], 16, 64)
		if err != nil {
			log.Fatal(err)
		}
		offset = int64(n)
	}

	numKb, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	f.Seek(offset, 0)
	for i := 0; i < numKb; i++ {
		counter++
		n, err := f.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Println(string(buf[:n]))
	}
	fmt.Println("Buffer counter:", counter)
}
