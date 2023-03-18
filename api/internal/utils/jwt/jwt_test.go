package jwt

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestJwtGenerateToken(t *testing.T) {
	m := BaseClaims{
		UUID:     uuid.New(),
		ID:       123,
		Username: "test",

		AuthorityId: 1,
	}

	// 生成10秒过期时间的Token
	tokenString, err := GenerateToken(m, 10*time.Second)
	if err != nil {
		t.Fatalf("GenerateToken() failed, err: %v", err)
	}

	claims, err := ParseToken(tokenString)
	if err != nil {
		t.Fatalf("ParseToken() failed, err: %v", err)
	}

	if claims.Username != "test" {
		t.Errorf("Jwt token claims not match, exp: %s, got: %s", "test", claims.Username)
	}
}

//func TestJwtParseUser(t *testing.T) {
//	type args struct {
//		tokenString string
//	}
//	tests := []struct {
//		name    string
//		args
//		want    *BaseClaims
//		wantErr bool
//	}{
//		//
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := ParseToken(tt.args.tokenString)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_jwtGenerateToken(t *testing.T) {
//	type args struct {
//		m BaseClaims
//		d time.Duration
//	}
//	tests := []struct {
//		name    string
//		args
//		want    string
//		wantErr bool
//	}{
//		// .
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GenerateToken(tt.args.m, tt.args.d)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_userStdClaims_Valid(t *testing.T) {
//	type fields struct {
//		StandardClaims jwt.StandardClaims
//		BaseClaims     BaseClaims
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantErr bool
//	}{
//		//
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := userStdClaims{
//				StandardClaims: tt.fields.StandardClaims,
//				BaseClaims:     tt.fields.BaseClaims,
//			}
//			if err := c.Valid(); (err != nil) != tt.wantErr {
//				t.Errorf("Valid() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
