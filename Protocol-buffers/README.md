# Protocol buffer
~~~protoc
syntax = "proto3";
package tutorial;

message Person {
	string name = 1;
	int32 id = 2;
	string email = 3;

	enum PhoneType {
		MOBILE = 0;
		HOME = 1;
		WORK = 2;
	}

	message PhoneNumber {
		string number = 1;
		PhoneType type = 2;
	}

	repeated PhoneNumber phones = 4;
}

message AddressBook {
	repeated Person people = 1;
}
~~~

# Melakukan generate file .proto ke file pb.go
~~~bash
protoc -I=go/src/github.com/jiharAkakom/Protocol-buffers --go_out=$go/src/github.com/jiharAkakom/Protocol-buffers/server go/src/github.com/jiharAkakom/Protocol-buffers/addressbook.proto
~~~
