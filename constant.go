package clash

import (
	"encoding/base64"
	"github.com/koomox/ext"
)

type Config struct {
	rules   []byte
	prefix  []byte
	isValid bool
}

var (
	current   = &Config{isValid: false}
	rawRules  = "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",
	rawPrefix = "eJxUkMFugzAMhu+R8g5W1Wvo6Napykts2mk7pmBaVjfJnEDH208hoWKKEObzL/LZwTXXoLzjqKF+2h+OUhgid1dkrIbOUEApbq5FDTwQSkHurAhHJA297ZwUvR9fH8nWBi0FAFpzInxgAGtuGJBH5LkPoKCuX6rVs+BjNZ/lM1LQu11rQ8XDaer6cKkaq4+H5yVwidGnSF3NJ0XVz4A8SSGFZ/fbY3ZSs4SGTWT3bewm/yBOHjVklEnRhG0uMswb2qZXASaEu+M2wVLmhiFvy4yz3n5Vx+iTZ7nn2nvVIEc1IvfdlCwGXKwndWY3+P/q7x9vn19r74CEzWK0mrXsrkz1FwAA//+AgYVJ"
)

func deCompress(raw string) (b []byte, err error) {
	if b, err = base64.RawStdEncoding.DecodeString(raw); err != nil {
		return
	}
	return ext.NewEncoding().DeCompress(b)
}

func Initial() {
	if current.isValid {
		return
	}
	b, err := deCompress(rawRules)
	if err != nil {
		return
	}
	current.rules = b
	if b, err = deCompress(rawPrefix); err != nil {
		return
	}
	current.prefix = b
	current.isValid = true
}