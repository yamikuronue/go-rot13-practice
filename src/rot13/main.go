package main

import (
	"github.com/micro/go-micro";
	"context";
	rot13 "rot13/proto"
	)

func EncodeRot13(s string) string {
	var output []byte

	for i := 0; i < len(s); i++ {
		if s[i] >= 97 && s[i] <= 109 || s[i] >= 65 && s[i] <= 77 {
			output = append(output, s[i]+13)
		} else if s[i] >= 110 && s[i] <= 122 || s[i] >= 78 && s[i] <= 90 {
			output = append(output, s[i]-13)
		} else {
			output = append(output, s[i])
		}
	}
	return string(output[:])
}

func main() {
	service := micro.NewService(
		micro.Name("rot13"),
	)

	service.Init()
	service.Run()
}

type Rot13 struct{}

func (g *Rot13) Encode(ctx context.Context, req *rot13.EncodeRequest, rsp *rot13.EncodeResponse) error {
	rsp.Output = EncodeRot13(req.Input)
	return nil
}
