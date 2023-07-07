package xgo

import (
	"reflect"
	"testing"
)

func Test_mapValue(t *testing.T) {
	type args struct {
		v    int
		oMin int
		oMax int
		tMin int
		tMax int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "No change",
			args: args{
				v:    10,
				oMin: 0,
				oMax: 100,
				tMin: 0,
				tMax: 100,
			},
			want: 10,
		},
		{
			name: "Double",
			args: args{
				v:    10,
				oMin: 0,
				oMax: 100,
				tMin: 0,
				tMax: 200,
			},
			want: 20,
		}, {
			name: "Halve",
			args: args{
				v:    10,
				oMin: 0,
				oMax: 100,
				tMin: 0,
				tMax: 50,
			},
			want: 5,
		}, {
			name: "Shift up",
			args: args{
				v:    10,
				oMin: 0,
				oMax: 100,
				tMin: 100,
				tMax: 200,
			},
			want: 110,
		}, {
			name: "Shift down",
			args: args{
				v:    10,
				oMin: 0,
				oMax: 100,
				tMin: -100,
				tMax: 0,
			},
			want: -90,
		},
		{
			name: "Below minimum",
			args: args{
				v:    10,
				oMin: 20,
				oMax: 120,
				tMin: 20,
				tMax: 120,
			},
			want: 20,
		}, {
			name: "Above maximum",
			args: args{
				v:    1000,
				oMin: 20,
				oMax: 120,
				tMin: 20,
				tMax: 120,
			},
			want: 120,
		}, {
			name: "Reverse target",
			args: args{
				v:    50,
				oMin: 0,
				oMax: 100,
				tMin: 2000,
				tMax: 1000,
			},
			want: 1500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapValue(tt.args.v, tt.args.oMin, tt.args.oMax, tt.args.tMin, tt.args.tMax); got != tt.want {
				t.Errorf("mapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateMessage(t *testing.T) {
	type args struct {
		v []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "No data",
			args: args{
				v: []byte{},
			},
			want: []byte{0x55, 0x00, 0x06, 0xf9, 0x00, 0xaa},
		}, {
			name: "Rotate",
			args: args{
				v: []byte{0x00, 0x32, 0xa0},
			},
			want: []byte{0x55, 0x00, 0x09, 0x00, 0x32, 0xa0, 0x24, 0x00, 0xaa},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMessage(tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
