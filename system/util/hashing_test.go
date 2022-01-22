package util

import "testing"

func TestBase64RandomToken(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Base64 Random Token Test 1", args: args{length: 30}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := Base64RandomToken(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64RandomToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("B64Token->", gotToken)
			// if len(gotToken) != tt.args.length {
			// 	t.Errorf("Base64RandomToken() gotToken length = %d, want %v", len(gotToken), tt.args.length)
			// }
		})
	}
}

func TestRandomToken(t *testing.T) {
	type args struct {
		ln int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test Random token - 1", args: args{ln: 30}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := RandomToken(tt.args.ln)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("RandomToken->", gotToken)
		})
	}
}
