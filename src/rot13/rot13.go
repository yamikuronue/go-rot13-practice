package rot13

import (
	"github.com/micro/go-micro"
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

func (g *Rot13) Encode(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.output = EncodeRot13(req.input)
	return nil
}