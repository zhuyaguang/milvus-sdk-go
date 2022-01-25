package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func main()  {
	err :=ReadFile("claims_embedding")
	if err !=nil{
		fmt.Println("err",err)
	}
	return
}

func ReadFile(filePath string) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	for {
		line, _,err := buf.ReadLine()
		fmt.Println(string(line))
		if err != nil {
			if err == io.EOF{
				return nil
			}
			return err
		}
		return nil
	}
}

func loadBinary()  {
	f, err := os.Open("myfile")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func ExampleRead() {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Print(pi)
	// Output: 3.141592653589793
}

func ExampleRead_multi() {
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40, 0xff, 0x01, 0x02, 0x03, 0xbe, 0xef}
	r := bytes.NewReader(b)

	var data struct {
		PI   float64
		Uate uint8
		Mine [3]byte
		Too  uint16
	}

	if err := binary.Read(r, binary.LittleEndian, &data); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(data.PI)
	fmt.Println(data.Uate)
	fmt.Printf("% x\n", data.Mine)
	fmt.Println(data.Too)
	// Output:
	// 3.141592653589793
	// 255
	// 01 02 03
	// 61374
}
