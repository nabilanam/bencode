package main

// Example usage

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nabilanam/bencode/decoder"
	"github.com/nabilanam/bencode/encoder"
)

func main() {
	if len(os.Args) == 2 {
		f, err := ioutil.ReadFile(os.Args[1])
		if err == nil {
			d := decoder.New(f)
			m := d.Decode().(map[string]interface{})
			fmt.Printf("decoded:\n%+v", m)

			e := encoder.New(m)
			b := e.Encode()
			fmt.Printf("\n\nencoded:\n%+v\n", b)
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Bencoded file path is required as an argument!")
	}
}
