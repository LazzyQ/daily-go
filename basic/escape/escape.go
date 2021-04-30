package main
import (
    "crypto/sha1"
    "crypto/sha256"
    "fmt"
    "golang.org/x/crypto/pbkdf2"
    "hash"
    "math/rand"
)

// harbor admin 密码修改方法：
// update harbor_user set salt='', password='', password_version ='' where username='admin';

const (
    // EncryptHeaderV1 ...
    EncryptHeaderV1 = "<enc-v1>"
    // SHA1 is the name of sha1 hash alg
    SHA1 = "sha1"
    // SHA256 is the name of sha256 hash alg
    SHA256 = "sha256"
)

// HashAlg used to get correct alg for hash
var HashAlg = map[string]func() hash.Hash{
    SHA1:   sha1.New,
    SHA256: sha256.New,
}

func GenerateRandomStringWithLen(length int) string {
    const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    l := len(chars)
    result := make([]byte, length)
    _, err := rand.Read(result)
    if err != nil {
        fmt.Printf("error reading random bytes: %v", err)
    }
    for i := 0; i < length; i++ {
        result[i] = chars[int(result[i])%l]
    }
    return string(result)
}

func Encrypt(content string, salt string, encrptAlg string) string {
    return fmt.Sprintf("%x", pbkdf2.Key([]byte(content), []byte(salt), 4096, 16, HashAlg[encrptAlg]))
}

func main() {
    // 这要设置这里的明文密码变量，就可以生成对应 salt、password 信息
    password := "123456"
    salt := GenerateRandomStringWithLen(32)
    passwordEncry := Encrypt(password, salt, SHA256)
    fmt.Printf("明文密码: %s\nsalt: %s\npassword: %s\npassword_version: sha256\n", password, salt, passwordEncry)
}
