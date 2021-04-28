package pbconv

import (
	"encoding/base64"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

const wrapper = byte(34)

// ToBase64JsonString convert proto message to base 64 then wrap with `"`
func ToBase64JsonString(m proto.Message) ([]byte, error) {
	body, err := proto.Marshal(m)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to marshal the message")
	}

	buf := make([]byte, base64.RawStdEncoding.EncodedLen(len(body)))

	base64.RawStdEncoding.Encode(buf, body)

	buf = append([]byte{wrapper}, buf...)
	buf = append(buf, wrapper)

	return buf, nil
}

// ToBase64 convert proto message to base64
func ToBase64(m proto.Message) ([]byte, error) {
	body, err := proto.Marshal(m)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to marshal the message")
	}

	buf := make([]byte, base64.RawStdEncoding.EncodedLen(len(body)))

	base64.RawStdEncoding.Encode(buf, body)

	return buf, nil
}

// FromBase64JsonString unwrap `"` then unmarshal to base64 then unmarshal to proto message
func FromBase64JsonString(src []byte, dst proto.Message) error {
	srcLen := len(src)
	unWrappedSrc := src[0:]
	if src[0] == wrapper && src[srcLen-1] == wrapper {
		unWrappedSrc = src[1 : srcLen-1]
	}

	buf := make([]byte, base64.RawStdEncoding.DecodedLen(len(unWrappedSrc)))

	if _, err := base64.RawStdEncoding.Decode(buf, unWrappedSrc); err != nil {
		return err
	}

	if err := proto.Unmarshal(buf, dst); err != nil {
		return errors.WithMessage(err, "failed to marshal the message")
	}

	return nil
}

// FromBase64  base64 then unmarshal to proto message
func FromBase64(src []byte, dst proto.Message) error {
	buf := make([]byte, base64.RawStdEncoding.DecodedLen(len(src)))

	if _, err := base64.RawStdEncoding.Decode(buf, src); err != nil {
		return err
	}

	if err := proto.Unmarshal(buf, dst); err != nil {
		return errors.WithMessage(err, "failed to marshal the message")
	}

	return nil
}
