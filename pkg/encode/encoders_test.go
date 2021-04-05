package encode

import (
	"reflect"
	"testing"
)

func TestFactory(t *testing.T) {
	tables := []struct {
		alg    Alg
		typeOf string
	}{
		{"sha1", "*encode.EncoderSha1"},
		{"md5", "*encode.EncoderMd5"},
	}

	for _, table := range tables {
		e, err := Factory(table.alg)
		if err != nil {
			t.Error("Encoders factory should never return error")
		}

		oType := reflect.TypeOf(e).String()
		if table.typeOf != oType {
			t.Errorf("Encoders factory returned incorrect object type, got: %s, want: %s.", oType, table.typeOf)
		}
	}
}
