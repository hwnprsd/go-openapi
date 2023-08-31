package types

import (
	"encoding/json"
	"fmt"
	"testing"
)

type RandomReq struct {
	A     string
	Hello []string
	Lul   uint
}

type RandomRes struct {
	A        string
	ThankYou string
}

func Handler(req RandomReq) (*RandomRes, error) {
	return nil, nil
}

func Test_Spec(t *testing.T) {
	fmt.Println("Nastive OpenAPI in GO")
	openApi := NewOpenAPIData("API Docs", "v1.0")
	// Post(openApi, "/hi", Handler, WithSummary("Lulli"), WithTags([]string{"Test"}))
	openApi.XPost("/hi", Handler, WithSummary("Lulli"), WithTags([]string{"Test"}))
	jsonData, err := json.Marshal(openApi)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
