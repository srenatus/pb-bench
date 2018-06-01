package handler

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/golang/protobuf/jsonpb"
	structpb "github.com/golang/protobuf/ptypes/struct"

	api "github.com/srenatus/pb-bench/r-3a3da3/api"
)

var (
  result api.Run
  resultBytes []byte
  resultStruct structpb.Struct
)

func BenchmarkRunMsgMarshal(b *testing.B) {
	var r []byte
	var err error
	r0 := getRunMsg(b)

	for n := 0; n < b.N; n++ {
		r, err = proto.Marshal(r0)
		if err != nil {
			b.Error(err)
		}
	}
	resultBytes = r
}

func BenchmarkRunStructMarshal(b *testing.B) {
	var r []byte
	var err error
	r0 := getRunStruct(b)

	for n := 0; n < b.N; n++ {
		r, err = proto.Marshal(&r0)
		if err != nil {
			b.Error(err)
		}
	}
	resultBytes = r
}

func BenchmarkRunMsgUnmarshal(b *testing.B) {
	r0 := getRunMsg(b)
	r := api.Run{}
	var err error
	// marshal once, unmarshal in benchmark
	bs, err := proto.Marshal(r0)
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		err = proto.Unmarshal(bs, &r)
		if err != nil {
			b.Error(err)
		}
	}
	result = r
}

func BenchmarkRunStructUnmarshal(b *testing.B) {
	var err error
  r := structpb.Struct{}
	r0 := getRunStruct(b)
	bs, err := proto.Marshal(&r0)
	if err != nil {
		b.Fatal(err)
	}

	for n := 0; n < b.N; n++ {
		err = proto.Unmarshal(bs, &r)
		if err != nil {
			b.Error(err)
		}
	}
	resultStruct = r
}

func getRunStruct(b *testing.B) structpb.Struct {
	b.Helper()
	var runStruct structpb.Struct
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}

	content, err := ioutil.ReadFile("../data/converge-success-report.json")
	if err != nil {
		b.Fatal(err)
	}
	if err := unmarshaler.Unmarshal(strings.NewReader(string(content)), &runStruct); err != nil {
		b.Fatal(err)
	}
	return runStruct
}

func getRunMsg(b *testing.B) *api.Run {
	b.Helper()
  r := api.Run{}
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}

	content, err := ioutil.ReadFile("../data/converge-success-report.json")
	if err != nil {
		b.Fatal(err)
	}
	if err := unmarshaler.Unmarshal(strings.NewReader(string(content)), &r); err != nil {
		b.Fatal(err)
	}
  return &r
}
