package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

// Signature 签发签名
func Signature(header string, payload string, secretKey string, algorithm ...string) string {
	alg := ""
	plain := header + "." + payload
	if len(algorithm) > 0 {
		alg = algorithm[0]
	}
	switch strings.ToLower(alg) {
	case "hs256":
		return hs256(plain, secretKey)
	default:
		return simple(plain)
	}
}

// HasSign 判断签名
func HasSign(token string) bool {
	arr := strings.Split(token, ".")
	return len(arr) == 3
}

// VerifySign 验证签名
func VerifySign(token string, secretKey string, algorithm ...string) bool {
	alg := ""
	if len(algorithm) > 0 {
		alg = algorithm[0]
	}
	header, payload, sign := Disassemble(token)
	return sign != "" && Signature(header, payload, secretKey, alg) == sign
}

// Encode 编码
func Encode(data map[string]interface{}) (jsonStr string, err error) {
	jsonData, err := json.Marshal(data)
	jsonStr = base64.RawURLEncoding.EncodeToString(jsonData)
	return jsonStr, err
}

// Decode 解码
func Decode(str string) (jsonData map[string]interface{}, err error) {
	jsonStr, err := base64.RawURLEncoding.DecodeString(str)
	err = json.Unmarshal([]byte(jsonStr), &jsonData)
	return jsonData, err
}

// Assemble 组装
func Assemble(header string, payload string, sign string) string {
	return header + "." + payload + "." + sign
}

// Disassemble 拆解
func Disassemble(token string) (string, string, string) {
	arr := strings.Split(token, ".")
	if len(arr) > 2 {
		return arr[0], arr[1], arr[2]
	} else if len(arr) == 1 {
		return arr[0], "", ""
	} else {
		return "", "", ""
	}
}
