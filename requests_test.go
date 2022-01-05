package erh

import (
	"reflect"
	"testing"
)

func TestArgAmount(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name string
		args args
		want Arg
	}{
		{"Usage", args{amount: 10}, Arg{data: "amount=10.000000"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArgAmount(tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgSymbols(t *testing.T) {
	type args struct {
		symbols []string
	}
	tests := []struct {
		name string
		args args
		want Arg
	}{
		{"Usage", args{symbols: []string{"USD", "EUR"}}, Arg{data: "symbols=USD,EUR"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArgSymbols(tt.args.symbols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgPlaces(t *testing.T) {
	type args struct {
		places int
	}
	tests := []struct {
		name string
		args args
		want Arg
	}{
		{"Usage", args{places: 2}, Arg{data: "places=2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArgPlaces(tt.args.places); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgPlaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgBase(t *testing.T) {
	type args struct {
		base string
	}
	tests := []struct {
		name string
		args args
		want Arg
	}{
		{"Usage", args{base: "USD"}, Arg{data: "base=USD"}},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArgBase(tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArgBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_argsToString(t *testing.T) {
	type args struct {
		args []Arg
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Usage", args{[]Arg{ArgAmount(10), ArgSymbols([]string{"USD", "EUR"}), ArgPlaces(2), ArgBase("JPY")}},"amount=10.000000&symbols=USD,EUR&places=2&base=JPY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := argsToString(tt.args.args); got != tt.want {
				t.Errorf("argsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}