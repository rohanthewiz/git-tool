package util

import (
	"reflect"
	"testing"
)

func TestFloatToString(t *testing.T) {
	type args struct {
		number float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "float to string - 1",
			args: args{number: 3.1415927},
			want: "3.14",
		},
		{
			name: "float to string - 2",
			args: args{number: 7.0},
			want: "7.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatToString(tt.args.number); got != tt.want {
				t.Errorf("FloatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name:    "Join some words",
			args:    args{words: []string{"cat ", "  dog", " mice ", "men"}},
			wantOut: "cat dog mice men",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := JoinWords(tt.args.words...); gotOut != tt.wantOut {
				t.Errorf("JoinWords() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestParseStrToFloat(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Str to Float",
			args: args{str: "3.141"},
			want: 3.141,
		},
		{
			name: "Str to Float",
			args: args{str: "abc3.141"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseStrToFloat(tt.args.str); got != tt.want {
				t.Errorf("ParseStrToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlugify(t *testing.T) {
	type args struct {
		instr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Slugify - maintain upper",
			args: args{"Test->àèâ<-Test"},
			want: "test-aea-test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slugify(tt.args.instr); got != tt.want {
				t.Errorf("Slugify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrSplitAndTrim(t *testing.T) {
	type args struct {
		instr     string
		separator string
	}
	tests := []struct {
		name    string
		args    args
		wantOut []string
	}{
		{
			name:    "String Split and Trim",
			args:    args{instr: "abcd .efg . hijk", separator: "."},
			wantOut: []string{"abcd", "efg", "hijk"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := StrSplitAndTrim(tt.args.instr, tt.args.separator); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("StrSplitAndTrim() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
