package wal

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func MarshalLogEntry(entry *LogEntry) []byte {
	marshaledEntry, err := proto.Marshal(entry)
	if err != nil {
		panic(fmt.Sprintf("marshalling should never fail (%v)", err))
	}
	return marshaledEntry
}

func UnMarshalLogEntry(data []byte, entry *LogEntry) {
	err := proto.Unmarshal(data, entry)
	if err != nil {
		panic(fmt.Sprintf("unmarshalling should never fail (%v)", err))
	}
}
