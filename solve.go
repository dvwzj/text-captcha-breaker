package tcb

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
)

var (
	ErrInvalidImage = errors.New("invalid image")
)

// Solve returns the text from the image
// input can be a file path or a base64 string
func Solve(input string) (string, error) {
	// check input is starts with "data:" and contains "base64,"
	// if not, assume it's a file path and read the file
	// if it is, decode the base64
	// send the base64 image to the server
	// receive the text from the server
	// return the text

	// if it not a base64 string, assume it's a file path
	if !strings.HasPrefix(input, "data:") && !strings.Contains(input, "base64,") {
		// read the file
		// encode the file to base64
		// set image to the base64 string
		file, err := os.Open(input)
		if err != nil {
			return "", err
		}
		defer file.Close()
		fi, err := file.Stat()
		if err != nil {
			return "", err
		}
		size := fi.Size()
		data := make([]byte, size)
		_, err = file.Read(data)
		if err != nil {
			return "", err
		}
		// get image format
		_, format, err := image.DecodeConfig(file)
		if err != nil {
			return "", err
		}
		// encode to base64
		input = "data:image/" + format + ";base64," + base64.StdEncoding.EncodeToString(data)
	}
	if !strings.HasPrefix(input, "data:") || !strings.Contains(input, "base64,") {
		return "", ErrInvalidImage
	}
	// connect to the server
	client := NewClient()
	err := client.Connect()
	if err != nil {
		return "", err
	}
	defer client.Close()
	generatedHash := GenerateHash()
	for {
		msg, err := client.Receive()
		if err != nil {
			return "", err
		}
		if len(msg) == 0 {
			return "", errors.New("empty message")
		}
		jsonMsg := make(map[string]interface{})
		err = json.Unmarshal(msg, &jsonMsg)
		if err != nil {
			return "", err
		}
		if _, ok := jsonMsg["msg"]; !ok {
			return "", errors.New("invalid message")
		}
		v := jsonMsg["msg"]
		switch v {
		case "send_hash":
			sendHash := []byte(`{"fn_index":0,"session_hash":"` + generatedHash + `"}`)
			err = client.Send(sendHash)
			if err != nil {
				return "", err
			}
		case "estimation":
			// do nothing
		case "send_data":
			sendData := []byte(`{"data": ["` + input + `"],"event_data": null,"fn_index": 0,"session_hash": "` + generatedHash + `"}`)
			err = client.Send(sendData)
			if err != nil {
				return "", err
			}
		case "process_starts":
			// do nothing
		case "process_completed":
			if v, ok := jsonMsg["output"]; !ok {
				return "", errors.New("invalid message")
			} else {
				output := v.(map[string]interface{})
				if v, ok := output["data"]; !ok {
					return "", errors.New("invalid message")
				} else {
					data := v.([]interface{})
					if len(data) == 0 {
						return "", errors.New("invalid message")
					}
					return data[0].(string), nil
				}
			}
		default:
			fmt.Println("default", v)
		}
	}
}
