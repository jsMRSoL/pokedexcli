package main

import (
	"testing"
	"time"

	"github.com/jsMRSoL/pokedexcli/internal/pokecache"
)

func Test_mapForward(t *testing.T) {
	type args struct {
		c    *config
		args []string
	}

	nonUrl := "notaUrl.really"
	pc := pokecache.NewCache(time.Minute * 5)
  
	myConfig := config{
		prev: nil,
		next: &nonUrl,
    cache: pc,
	}

	emptyArgs := []string{}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Broken config",
			args: args{
				c:    &myConfig,
				args: emptyArgs,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mapForward(tt.args.c, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("mapForward() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
