package guid

import (
	"errors"
	"fmt"

	"github.com/twcclan/go-guid/pb-md5"
)

const (
	Seed1 uint32 = 0x00b684a3
	Seed2 uint32 = 0x00051a56
)

var (
	ErrInvalidEtKey = errors.New("Key needs to be 18 bytes long")
)

func Calculate(etkey string) (guid string, err error) {
	if len(etkey) != 18 {
		return "", ErrInvalidEtKey
	}

	guid = fmt.Sprintf("%x", md5.Sum([]byte(etkey), Seed1))
	guid = fmt.Sprintf("%X", md5.Sum([]byte(guid), Seed2))

	return
}
