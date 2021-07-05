package main

import (
	"github.com/tech-botao/go-implement/app/protobufer/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)


func createPhone(n string, t pb.Person_PhoneType) *pb.Person_PhoneNumber {
	ph := &pb.Person_PhoneNumber{
		Number: n,
		Type:   t,
	}
	return ph
}


func createPerson(name string, id int64, email string) *pb.Person {
	ph := &pb.Person{
		Name: name,
		Id: id,
		Email: &email,
		Phones: []*pb.Person_PhoneNumber{},
		LastUpdated: timestamppb.Now(),
		//LastUpdate:
	}
	return ph
}

func addPhonesToPerson(person *pb.Person, phones ...*pb.Person_PhoneNumber) {
	phoneNumber :=  make([]*pb.Person_PhoneNumber, 0)
	for _, phone := range phones {
		phoneNumber = append(phoneNumber, phone)
	}
	person.Phones = phoneNumber
}

func main() {




}
