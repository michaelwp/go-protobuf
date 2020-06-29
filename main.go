package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/michaelwp/go-grpc/protofiles"
)

func main(){
	p := &pb.Person{
		Id:1234,
		Name: "Roger F",
		Email: "rf@example.com",
		Phones: []*pb.Person_PhoneNumber {
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	printRes(p)
}

func printRes(p *pb.Person) {
	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	jsonRes, _ := json.Marshal(p)

	fmt.Println("Original struct loaded from proto file: ", p, "\n")
	fmt.Println("Marshaled proto data: ", body, "\n")
	fmt.Println("Unmarshaled struct: ", p1)
	fmt.Println("Json format: ", string(jsonRes))
}


