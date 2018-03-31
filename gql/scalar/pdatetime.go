package scalar

import (
	"github.com/golang/protobuf/ptypes"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

func serializeDateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case google_protobuf.Timestamp:
		t, err := ptypes.Timestamp(&value)
		if err != nil {
			return nil
		}
		buff, err := t.MarshalText()
		if err != nil {
			return nil
		}
		return string(buff)
	case *google_protobuf.Timestamp:
		return serializeDateTime(*value)
	default:
		return nil
	}

	return nil
}

func unserializeDateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case []byte:
		t := google_protobuf.Timestamp{}
		tt, err := ptypes.Timestamp(&t)
		if err != nil {
			return nil
		}
		err = tt.UnmarshalText(value)
		if err != nil {
			return nil
		}

		return t
	case string:
		return unserializeDateTime([]byte(value))
	case *string:
		return unserializeDateTime([]byte(*value))
	default:
		return nil

	}
}

var ProtoDateTime = graphql.NewScalar(graphql.ScalarConfig{
	Name: "ProtoDateTime",
	Description: "The `DateTime` scalar type represents a Protobuf DateTime." +
		" The DateTime is serialized as an RFC 3339 quoted string",
	Serialize:  serializeDateTime,
	ParseValue: unserializeDateTime,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})
