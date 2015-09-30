package main

import "encoding/json"
import "os"

func main() {

	dec := json.NewDecoder(os.Stdin)
	var dat map[string]interface{}
	if err := dec.Decode(&dat); err != nil {
		panic(err)

	}
	b, err := json.MarshalIndent(dat, "", "  ")
	if err != nil {
		panic(err)
	}
	b2 := append(b, '\n')
	os.Stdout.Write(b2)

}
