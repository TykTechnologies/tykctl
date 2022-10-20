package cmd

import "testing"

func Test_dashboardLogin(t *testing.T) {
	type args struct {
		loginClient   LoginClient
		url           string
		email         string
		password      string
		basicUser     string
		basicPassword string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := dashboardLogin(tt.args.loginClient, tt.args.url, tt.args.email, tt.args.password, tt.args.basicUser, tt.args.basicPassword); (err != nil) != tt.wantErr {
				t.Errorf("dashboardLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
