package serialize

import (
	"encoding/json"
	"io"
)

// JSONSerializer writes a Point as a JSON object.
type JSONSerializer struct{}

// Serialize writes the point as JSON.
func (s *JSONSerializer) Serialize(p *Point, w io.Writer) (err error) {
	buf := scratchBufPool.Get().([]byte)

	obj := map[string]interface{}{}
	for i, tagKey := range p.tagKeys {
		obj[string(tagKey)] = string(p.tagValues[i])
	}
	for i, fieldKey := range p.fieldKeys {
		obj[string(fieldKey)] = p.fieldValues[i]
	}
	obj["_ts"] = p.timestamp.UTC().UnixNano()

	err = json.NewEncoder(w).Encode(obj)

	buf = buf[:0]
	scratchBufPool.Put(buf)

	return err
}
