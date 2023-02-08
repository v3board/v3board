package password

import (
	"encoding/base64"
	"v3board/global"

	"golang.org/x/crypto/argon2"
)

// Hash生成密文
func Hash(password string) string {
	bytePasword := argon2.IDKey([]byte(password), []byte(global.Sault), 1, 64*1024, 4, 64)

	return base64.StdEncoding.EncodeToString(bytePasword)
}

// Verify 验证密码
func Verify(password, hashPassword string) bool {
	return Hash(password) == hashPassword
}
