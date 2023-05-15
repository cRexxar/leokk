package handle

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func EndWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func ByteToInterface(b []byte) (interface{}, error) {
	var tmp interface{}
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	if err := d.Decode(&tmp); err != nil {
		return nil, err
	} else {
		return tmp, nil
	}
}

func StrToFloat64(s string) float64 {
	if s == "" {
		return 0
	}
	if res, err := strconv.ParseFloat(s, 64); err != nil {
		return 0
	} else {
		return res
	}
}

func JsonToFloat64(j json.Number) float64 {
	return StrToFloat64(j.String())
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func BoolToStr(b bool) string {
	return strconv.FormatBool(b)
}
