package kit

import (
	"encoding/json"
	"regexp"
)

func GetIDFrom(url string) string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindString(url)
}

func ToBytes(i interface{}) []byte {
	b, err := json.Marshal(i)
	if err != nil {
		return nil
	}

	return b
}
