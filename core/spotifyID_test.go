package core

import (
	"reflect"
	"testing"
)

func TestSpotifyID_ToBase62(t *testing.T) {
	tests := []struct {
		name string
		s    SpotifyID
		want string
	}{
		{
			s:    SpotifyID([]byte{0x00, 0x0d, 0x53, 0x65, 0x35, 0x86, 0x4e, 0x0f, 0x99, 0x76, 0x1f, 0x9d, 0xa9, 0x00, 0xb1, 0xc1}),
			want: "0065zxtT6XKaQww7cLne0h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToBase62(); got != tt.want {
				t.Errorf("SpotifyID.ToBase62() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSpotifyIDFromBase62(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    SpotifyID
		wantErr bool
	}{
		{
			args:    args{id: "0065zxtT6XKaQww7cLne0h"},
			want:    SpotifyID([]byte{0x00, 0x0d, 0x53, 0x65, 0x35, 0x86, 0x4e, 0x0f, 0x99, 0x76, 0x1f, 0x9d, 0xa9, 0x00, 0xb1, 0xc1}),
			wantErr: false,
		},
		{
			args:    args{id: "0065z*(&}{}}||^^**&***&*&*&*tT6XKaQww7cLne0h"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSpotifyIDFromBase62(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSpotifyIDFromBase62() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpotifyIDFromBase62() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSpotifyIDFromRaw(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    SpotifyID
		wantErr bool
	}{
		{
			args:    args{data: []byte{0x00, 0x0d, 0x53, 0x65, 0x35, 0x86, 0x4e, 0x0f, 0x99, 0x76, 0x1f, 0x9d, 0xa9, 0x00, 0xb1, 0xc1}},
			want:    SpotifyID([]byte{0x00, 0x0d, 0x53, 0x65, 0x35, 0x86, 0x4e, 0x0f, 0x99, 0x76, 0x1f, 0x9d, 0xa9, 0x00, 0xb1, 0xc1}),
			wantErr: false,
		},
		{
			args:    args{data: []byte{0x00, 0x00, 0x00, 0x0d, 0x53, 0x65, 0x35, 0x86, 0x4e, 0x0f, 0x99, 0x76, 0x1f, 0x9d, 0xa9, 0x00, 0xb1, 0xc1}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSpotifyIDFromRaw(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSpotifyIDFromRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpotifyIDFromRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}
