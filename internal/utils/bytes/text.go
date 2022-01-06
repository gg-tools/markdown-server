package bytes

import "bytes"

func ToUnix(data []byte) []byte {
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))       // 删除Windows文本文件头BOM
	data = bytes.ReplaceAll(data, []byte("\r\n"), []byte("\n")) // 替换Windows换行符CRLF为UNIX换行符
	return data
}
