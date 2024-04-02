package main

import "testing"

func Test_getUrlString(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Handles single input with dashes",
			args: args{
				args: []string{"canalave-city-area"},
			},
			want:    "https://pokeapi.co/api/v2/location-area/canalave-city-area",
			wantErr: false,
		},
		{
			name: "Handles multiple words with spaces",
			args: args{
				args: []string{"canalave", "city", "area"},
			},
			want:    "https://pokeapi.co/api/v2/location-area/canalave-city-area",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUrlString(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUrlString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getUrlString() = %v, want %v", got, tt.want)
			}
		})
	}
}
