package main

import "testing"

func Test_mapForward(t *testing.T) {
	type args struct {
		c *config
	}

	nonUrl := "notaUrl.really"
	myConfig := config{
		prev: nil,
		next: &nonUrl,
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Broken config",
			args: args{
				&myConfig,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mapForward(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("mapForward() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
