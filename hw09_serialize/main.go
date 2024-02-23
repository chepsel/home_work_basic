package main

import (
	"fmt"

	"github.com/chepsel/home_work_basic/hw09_serialize/books"
	"github.com/chepsel/home_work_basic/hw09_serialize/protoc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Marshaller interface {
	MarshalJSON() ([]byte, error)
}

type Unmarshaller interface {
	UnmarshalJSON([]byte) error
}

func MarshalProto(protoBook Message) ([]byte, error) {
	marshaled, err := proto.Marshal(protoBook)
	return marshaled, err
}

type Message interface {
	String() string
	ProtoReflect() protoreflect.Message
	GetID() string
	GetRate() float32
	GetYear() uint32
	GetSize() uint32
	GetTitle() string
	GetAuthor() string
}

func UnmarshalProto(marshaled []byte, unmarshaled Message) error {
	if err := proto.Unmarshal(marshaled, unmarshaled); err != nil {
		return err
	}
	return nil
}

func ToProtobufStructure(book *books.Book) Message {
	protoBook := &protoc.Book{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
		Size:   book.Size,
		Rate:   book.Rate,
	}
	return protoBook
}

func main() {
	book := books.NewBook("978-5-389-21499-6", "Элегантность ежика", "Мюриель Барбери", 2009, 400, 2.4)
	result, err := book.MarshalJSON()
	if err != nil {
		fmt.Println("error is:", err)
	}
	fmt.Println("JSON from structure:", string(result))
	bookNew := &books.Book{}
	if err = bookNew.UnmarshalJSON(result); err != nil {
		fmt.Println("error is:", err)
	}
	fmt.Println("Unmarshled JSON into bookNew:", *bookNew)
	protoBook := ToProtobufStructure(book)
	marshaled, errProto := MarshalProto(protoBook)
	if errProto != nil && protoBook.ProtoReflect().IsValid() {
		fmt.Println("error is:", errProto)
	}

	fmt.Println("Id:", protoBook.GetID())
	fmt.Println("Rate:", protoBook.GetRate())
	fmt.Println("IsValid:", protoBook.ProtoReflect().IsValid())
	fmt.Println("Marshaled protoc:", string(marshaled))

	protoBookNew := ToProtobufStructure(&books.Book{})
	UnmarshalProto(marshaled, protoBookNew)
	fmt.Printf("Unmarshled Protoc into protoBookNew: %v\n", protoBookNew)
}
