package clash

import (
	"fmt"
	"strings"
)

func ForAndroid(password, host, port string) (b []byte, err error) {
	if !current.isValid {
		err = fmt.Errorf("initial failed")
		return
	}

	prefix := string(current.prefix)
	prefix = strings.Replace(prefix, "$server", host, -1)
	prefix = strings.Replace(prefix, "$port", port, -1)
	prefix = strings.Replace(prefix, "$password", password, -1)

	b = []byte(prefix)
	b = append(b, current.rules...)

	return
}

func ParseTrojanLink(link string) (password, host, port, tag string, err error) {
	if strings.ToLower(link[:9]) == "trojan://" {
		link = link[9:]
	} else {
		err = fmt.Errorf("bad request")
		return
	}
	offset := 0
	for i := 0; i < len(link); i++ {
		switch link[i] {
		case '@':
			password = link[:i]
			offset = i + 1
		case ':':
			host = link[offset:i]
			offset = i + 1
			port = link[offset:]
		case '#':
			port = link[offset:i]
			offset = i + 1
			tag = link[offset:]
		}
	}

	return
}