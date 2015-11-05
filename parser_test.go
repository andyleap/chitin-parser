package main

import (
	"fmt"
	"testing"
)

func Test_Message(t *testing.T) {
	schema := `chitin v1

envelope MyProtocol {
	 messages {
	 	  1: Person v1
	 }
}

message Person v1 {
	wire format: v1

	slots {
	      age uint16
	}

	fields {
	       name string
	       phone string
	}
}

`
	parser := MakeParser()
	
	data, err := parser.ParseString(schema)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n",data)
}

func Benchmark_Message(b *testing.B) {
	schema := `chitin v1
	
message Person v1 {
	wire format: v1

	slots {
	      age uint16
	}

	fields {
	       name string
	       phone string
	}
}

message Person v2 {
	wire format: v1
	
	slots {
	      age uint16
	      siblings uint16
	}

	fields {
	       name string
	       phone string
	}
}

`
	parser := MakeParser()
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		parser.ParseString(schema)
	}
}