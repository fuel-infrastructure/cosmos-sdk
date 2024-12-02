package codec

import (
	"bytes"

	"github.com/cosmos/gogoproto/jsonpb"
	"github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/codec/types"
)

var defaultJM = &jsonpb.Marshaler{OrigName: true, EmitDefaults: true, AnyResolver: nil}

// ProtoMarshalJSON provides an auxiliary function to return Proto3 JSON encoded
// bytes of a message.
func ProtoMarshalJSON(msg proto.Message, resolver jsonpb.AnyResolver) ([]byte, error) {
	// We use the OrigName because camel casing fields just doesn't make sense.
	// EmitDefaults is also often the more expected behavior for CLI users
	jm := defaultJM
	if resolver != nil {
		jm = &jsonpb.Marshaler{OrigName: true, EmitDefaults: true, AnyResolver: resolver}
	}
	err := types.UnpackInterfaces(msg, types.ProtoJSONPacker{JSONPBMarshaler: jm})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := jm.Marshal(buf, msg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ProtoMarshalJSONSkipEmpty provides an auxiliary function to return Proto3 JSON
// encoded bytes of a message with empty fields omitted.
// It is especially useful for encoding messages to be sent over the wire.
func ProtoMarshalJSONSkipEmpty(msg proto.Message, resolver jsonpb.AnyResolver) ([]byte, error) {
	jm := defaultJM
	jm.EmitDefaults = false
	if resolver != nil {
		jm = &jsonpb.Marshaler{OrigName: true, EmitDefaults: false, AnyResolver: resolver}
	}
	err := types.UnpackInterfaces(msg, types.ProtoJSONPacker{JSONPBMarshaler: jm})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := jm.Marshal(buf, msg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
