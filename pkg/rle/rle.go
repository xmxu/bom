package rle

import (
	"bytes"
	"io"
)

func Decode(r io.Reader) ([]byte, error) {
	var decoded bytes.Buffer

	for {
		// 读取一个字符
		ch := make([]byte, 1)
		_, err := r.Read(ch)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		// 解析连续字符的重复次数
		count := 0
		for '0' <= ch[0] && ch[0] <= '9' {
			count = count*10 + int(ch[0]-'0')
			_, err := r.Read(ch)

			if err == io.EOF {
				break
			}

			if err != nil {
				return nil, err
			}
		}

		// 将字符重复count次写入结果
		for i := 0; i < count; i++ {
			decoded.WriteByte(ch[0])
		}
	}

	return decoded.Bytes(), nil

}
