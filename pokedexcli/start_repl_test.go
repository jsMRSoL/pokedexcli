package main

import (
	"reflect"
	"testing"
)

func Test_normalizeInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		wantCmd  string
		wantArgs []string
		wantErr  bool
	}{
		// TODO: Add test cases.
    {
      name: "Parse command with no args",
      args: args{
        input: "map",
      },
      wantCmd: "map",
      wantArgs: nil,
      wantErr: false,
    },
    {
      name: "Parse command preceded by space, with no args",
      args: args{
        input: " map",
      },
      wantCmd: "map",
      wantArgs: nil,
      wantErr: false,
    },
    {
      name: "Parse command with multiple args",
      args: args{
        input: "map lots of args",
      },
      wantCmd: "map",
      wantArgs: []string{ "lots", "of", "args" },
      wantErr: false,
    },
    {
      name: "Parse command preceded by space, with multiple args",
      args: args{
        input: " map lots of args",
      },
      wantCmd: "map",
      wantArgs: []string{ "lots", "of", "args" },
      wantErr: false,
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmd, gotArgs, err := normalizeInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("normalizeInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCmd != tt.wantCmd {
				t.Errorf("normalizeInput() gotCmd = %v, want %v", gotCmd, tt.wantCmd)
			}
			if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("normalizeInput() gotArgs = %v, want %v", gotArgs, tt.wantArgs)
			}
		})
	}
}
