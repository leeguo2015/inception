package captcha

import (
	"inception/api/internal/global"
	"testing"
)

func TestBaseApi_Captcha(t *testing.T) {

	global.InitConfig()
	tests := []struct {
		name     string
		wantId   string
		wantB64s string
		wantErr  bool
	}{
		{"a", "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BaseApi{}
			gotId, gotB64s, err := b.Captcha()
			if (err != nil) != tt.wantErr {
				t.Errorf("Captcha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log("gotId: ", gotId)
			t.Log("gotB64s: ", gotB64s)
			//if gotId != tt.wantId {
			//	t.Errorf("Captcha() gotId = %v, want %v", gotId, tt.wantId)
			//}
			//if gotB64s != tt.wantB64s {
			//	t.Errorf("Captcha() gotB64s = %v, want %v", gotB64s, tt.wantB64s)
			//}
		})
	}
}
