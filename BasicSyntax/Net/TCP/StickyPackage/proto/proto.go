package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// 为了防止粘包，需要自定义协议

// 编码
func Encode(message string) ([]byte, error) {
	length := int32(len(message))
	pkg := new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length) 
	if err != nil {
		return nil, err
	}
	// 写入消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// 解码
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息头
	lengthByte, _ := reader.Peek(4)
	lengthBuffer := bytes.NewBuffer(lengthByte)
	var length int32
	// 需要把读取到的字节数组转化成int32
	err := binary.Read(lengthBuffer, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	// 校验信息时是否完整
	if int32(reader.Buffered()) < length + 4 {
		return "", err
	}

	// 读取真实数据
	pack := make([]byte, int(length+4))
	_, err = reader.Read(pack)
	if err != nil {
		return "", nil
	}
	return string(pack[4:]), nil
}