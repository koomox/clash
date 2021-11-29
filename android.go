package clash

import (
	"fmt"
	"strings"
)

func ForAndroid(host, port, password string) (b []byte, err error) {
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

func Parser(params string) (host, port, password string, err error) {
	errBadRequest := fmt.Errorf("bad request")
	if !strings.HasPrefix(params, "trojan://") { // trojan://password@remote_host:remote_port
		err = errBadRequest
		return
	}
	b := strings.Split(params, "trojan://")
	if len(b) < 2 {
		err = errBadRequest
		return
	}
	c := strings.Split(b[1], "@")
	if len(c) < 2 {
		err = errBadRequest
		return
	}
	password = c[0]
	d := strings.Split(c[1], ":")
	if len(d) < 2 {
		err = errBadRequest
		return
	}
	host = d[0]
	port = d[1]
	return
}
