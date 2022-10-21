package util

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Uint64ToByte 将uint64转换为[]byte类型
func Uint64ToByte(data uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, data)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
