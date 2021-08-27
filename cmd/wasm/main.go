package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"img_to_css/cmd/process"
	"log"
	"syscall/js"
)

func convertWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		imageArr := args[0]
		colorDiff := args[1]
		minLineLen := args[2]
		callback := args[3]
		log.Println(args)
		inBuf := make([]uint8, imageArr.Get("byteLength").Int())
		js.CopyBytesToGo(inBuf, imageArr)

		reader := bytes.NewReader(inBuf)
		img, imgType, err := image.Decode(reader)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(imgType)

		html := process.Image(img, colorDiff.Float(), minLineLen.Int())

		callback.Invoke(html)

		return true
	})
	return jsonFunc
}

func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("convertToCSS", convertWrapper())
	<-make(chan bool)
}
