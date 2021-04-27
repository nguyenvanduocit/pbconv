package ptconv

import (
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/nguyenvanduocit/ptconv/testdata"
)

var fooMessage = &testdata.FooMessage{
	Message: "ahihi",
}

var fooMessageEncoded = []byte{34, 67, 103, 86, 104, 97, 71, 108, 111, 97, 81, 34}

func TestToBase64JsonString(t *testing.T) {

	type args struct {
		m proto.Message
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "foo",
			args: args{
				m: fooMessage,
			},
			want:    fooMessageEncoded,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBase64JsonString(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBase64JsonString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBase64JsonString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromBase64JsonString(t *testing.T) {
	type args struct {
		src []byte
		dst proto.Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "foo",
			args: args{
				src: fooMessageEncoded,
				dst: &testdata.FooMessage{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FromBase64JsonString(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("FromBase64JsonString() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.dst.String(), fooMessage.String()) {
				t.Errorf("ToBase64JsonString() got = %v, want %v", tt.args.dst, fooMessage)
			}
		})
	}
}
