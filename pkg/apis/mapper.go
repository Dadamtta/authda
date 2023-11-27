package apis

import "encoding/json"

func BodyMapper[T any](bytes []byte, t *T) error {
	return json.Unmarshal([]byte(bytes), &t)
}
