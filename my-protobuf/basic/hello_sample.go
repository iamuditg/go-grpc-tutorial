package basic

import (
	"github.com/iamudit/go-grpc-tutorial/my-protobuf/protogen/basic"
	"log"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Clark Kent",
	}
	log.Print(&h)
}
