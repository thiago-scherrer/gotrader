package convert

import (
	"reflect"
	"testing"
)

func TestBytesToString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToString(StringToBytes("42")); got != tt.want {
				t.Errorf("BytesToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToBytes(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{"Teste", StringToBytes("42")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToBytes("42"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Test", 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToInt("42"); got != tt.want {
				t.Errorf("StringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToString(42); got != tt.want {
				t.Errorf("IntToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatToString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test", "42.000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatToString(42.0); got != tt.want {
				t.Errorf("FloatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatToInt(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Test", 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatToInt(42); got != tt.want {
				t.Errorf("FloatToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
