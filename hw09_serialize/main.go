package main

import (
	"fmt"

	"github.com/chepsel/home_work_basic/hw09_serialize/books"
	"github.com/chepsel/home_work_basic/hw09_serialize/protoc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type SentinelError string

func (err SentinelError) Error() string {
	return string(err)
}

const (
	EmptySlice SentinelError = "Empty slice"
)

type Marshaller interface {
	MarshalJSON() ([]byte, error)
	MarshalProto() ([]byte, error)
}

type Unmarshaller interface {
	UnmarshalJSON([]byte) error
	UnmarshalProto(marshaled []byte) error
}

type Message interface {
	String() string
	ProtoReflect() protoreflect.Message // позволяет использовать вложенные в пакет protoc методы
	GetID() string
	GetRate() float32
	GetYear() uint32
	GetSize() uint32
	GetTitle() string
	GetAuthor() string
}

func MarshalPB(protoBook *protoc.Book) ([]byte, error) {
	var marshaled []byte
	var err error
	clonedBook := proto.Clone(protoBook).(*protoc.Book)
	if protoBook.ProtoReflect().IsValid() { // комманда проверки валидности
		marshaled, err = proto.Marshal(clonedBook)
		if err != nil {
			return nil, err
		}
	}
	return marshaled, nil
}

func UnmarshalPB(marshaled []byte) (*protoc.Book, error) {
	protoBook := &protoc.Book{}
	err := proto.Unmarshal(marshaled, protoBook)
	if err != nil {
		result := proto.Clone(protoBook).(*protoc.Book) // комманда копирования которая позволяет обойти лок структуры
		return result, err
	}
	return protoBook, nil
}

func MarshalPBSlice(protoBook *protoc.BooksSlice) ([]byte, error) {
	var marshaled []byte
	var err error
	clonedBook := proto.Clone(protoBook).(*protoc.BooksSlice)
	if protoBook.ProtoReflect().IsValid() { // комманда проверки валидности
		marshaled, err = proto.Marshal(clonedBook)
		if err != nil {
			return nil, err
		}
	}
	return marshaled, nil
}

func UnmarshalPBSlice(marshaled []byte) (*protoc.BooksSlice, error) {
	protoBook := &protoc.BooksSlice{}
	err := proto.Unmarshal(marshaled, protoBook)
	if err != nil {
		result := proto.Clone(protoBook).(*protoc.BooksSlice) // комманда копирования которая позволяет обойти лок структуры
		return result, err
	}
	return protoBook, nil
}

func ToProtobufStructure(bookSlice []*books.Book) *protoc.BooksSlice {
	protoBooks := &protoc.BooksSlice{}
	protoBooks.Books = make([]*protoc.Book, len(bookSlice))
	for i, element := range bookSlice {
		protoBooks.Books[i] = &protoc.Book{
			ID:     element.ID,
			Title:  element.Title,
			Author: element.Author,
			Year:   element.Year,
			Size:   element.Size,
			Rate:   element.Rate,
		}
	}
	result := proto.Clone(protoBooks).(*protoc.BooksSlice)
	return result
}

// func MarshalJSONSlice(ctr []*books.Book) ([]byte, error) {
// 	var resultJSON []byte
// 	if ctr != nil {
// 		for i, element := range ctr {
// 			if element != nil {
// 				result, err := element.MarshalJSON()
// 				switch {
// 				case err != nil:
// 					return nil, err
// 				case i == 0:
// 					resultJSON = result

// 				default:
// 					resultJSON = append(resultJSON, []byte(`,`)...)
// 					resultJSON = append(resultJSON, result...)
// 				}
// 			}
// 		}
// 		resultJSON = append([]byte(`[`), resultJSON...)
// 		resultJSON = append(resultJSON, []byte(`]`)...)
// 		return resultJSON, nil
// 	}
// 	return nil, EmptySlice
// }

func main() {
	tmp := `{"id":"978","title":"Ку-Ка","author":"-Ре-Ку","year":2009,"size":400,"rate":2.4}`
	var book books.Book
	err := book.UnmarshalJSON([]byte(tmp))
	if err != nil {
		fmt.Println("error is:", err)
	}
	fmt.Println("Marshal one:", book)
	result, err := book.MarshalJSON()
	if err != nil {
		fmt.Println("error is:", err)
	}
	fmt.Println("Unmarshal one:", string(result))

	bookSlice := make([]*books.Book, 3)
	bookSlice[0] = &books.Book{
		ID:     "978",
		Title:  "Мюриель ежик",
		Author: "Мюриель Барбери",
		Year:   2009,
		Size:   400,
		Rate:   2.4,
	}
	bookSlice[1] = &books.Book{
		ID:     "978",
		Title:  "ежик Барбери",
		Author: "Мюриель Барбери",
		Year:   2009,
		Rate:   2.4,
	}
	bookSlice[2] = &books.Book{
		ID:     "3232-re",
		Title:  "tttt",
		Author: "aaaa",
		Rate:   3.6,
	}

	result, err = books.MarshalJSONSlice(bookSlice)
	if err != nil {
		fmt.Println("error is:", err)
	}
	fmt.Println("JSON from structure:", string(result))

	bookNew, errUnmSlice := books.UnmarshalJSONSlice(result)
	if errUnmSlice != nil {
		fmt.Println("error is:", errUnmSlice)
	}
	fmt.Println("Unmarshl JSON structure:")
	for i, element := range bookNew {
		fmt.Println("	Unmarshled JSON element of bookNew[", i, "]:", *element)
	}

	protoBook := ToProtobufStructure(bookNew)
	marshaled, errProto := MarshalPBSlice(protoBook)
	if errProto != nil {
		fmt.Println("error is:", errProto)
	}

	fmt.Printf("\n----\nMarshaled protoc message:%s\n-----\n", marshaled)

	protoBookNew, errUnmarshalProto := UnmarshalPBSlice(marshaled)
	if errUnmarshalProto != nil {
		fmt.Println("error is:", errUnmarshalProto)
	}
	fmt.Println("Unmarshled Protoc structure:")
	for i := range protoBookNew.Books {
		fmt.Println("	Unmarshled Protoc element of protoBookNew[", i, "]:", protoBookNew.Books[i])
	}
	tmp = "\n\x11978-5-389-21499-6\x12#Элегантность ежика\x1a\x1dМюриель Барбери \xd9\x0f(\x90\x035\x9a\x99\x19@"
	protoBookOne, err := UnmarshalPB([]byte(tmp))
	if err != nil {
		fmt.Println("error is:", err)
	}
	fmt.Println("Unmarshled Protobuf element of protoBookOne:", protoBookOne)

	result, err = MarshalPB(protoBookOne)
	if err != nil {
		fmt.Println("error is:", err)
	}
	fmt.Println("Marshaled protoc message:", string(result))
}
