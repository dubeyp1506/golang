package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// 1. MARSHAL: Go Struct -> JSON
	user := User{
		Name:  "Priyanshu Dubey",
		Email: "Priyanshu.dubey@gmail.com",
	}

	fmt.Println("+++++++++Doing Marshal+++++++++")

	JsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(JsonData))

	// 2. UNMARSHAL: JSON -> Go Struct
	fmt.Println("++++++++Doing Unmarshal++++++++")
	var newUser User
	unmarshalErr := json.Unmarshal([]byte(JsonData), &newUser)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	fmt.Println(newUser.Name)
	fmt.Println(newUser.Email)

	// 3. CHECKING TYPES
	fmt.Println("++++++++Checking Types++++++++")
	fmt.Printf("Type of 'user': %T\n", user)         // Should be main.User
	fmt.Printf("Type of 'JsonData': %T\n", JsonData) // Should be []uint8 (which is a byte slice)
	fmt.Printf("Type of 'newUser': %T\n", newUser)   // Should be main.User

	// 4. DECODING: JSON -> Go Data (using Decoder)
	fmt.Println("++++++++Decoding with Decoder++++++++")
	userData := `{"Name": "Priyanshu", "Email": "Priyanshu@gmail.com"}`
	fmt.Printf("Type of 'userData': %T\n", userData) // Should be string
	reader := strings.NewReader(userData)
	fmt.Printf("Type of 'reader': %T value is %v\n", reader, reader) // Should be *strings.Reader
	decoder := json.NewDecoder(reader)

	var decodedUser User
	decodeErr := decoder.Decode(&decodedUser)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}
	fmt.Println(decodedUser)

	// 5. ENCODING: Go Data -> JSON
	fmt.Println("++++++++Encoding with Encoder++++++++")
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	encoderErr := encoder.Encode(&decodedUser)
	if encoderErr != nil {
		log.Fatal(encoderErr)
	}
	fmt.Println(buf.String())
	fmt.Printf("Type of 'buf': %T\n", buf) // Should be bytes.Buffer

}
