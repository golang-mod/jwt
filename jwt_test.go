package jwt

import "testing"

var header = "eyJhcHAiOiJiYWNrZW5kIiwiY29kZSI6ImMxYWQ5MGYwLTg2NjUtMTRmOC0yMThjLTFmZjk3NGM5NTE3YSIsImFsZyI6ImhzMjU2IiwiZW52IjpudWxsLCJleHAiOiI3MjAwIn0"
var payload = "eyJ1aWQiOiJMaUNoYW8ifQ"
var sign = "5e61afa2616f1fc6848915acba9c85497a4a24dd"
var token = "eyJhcHAiOiJiYWNrZW5kIiwiY29kZSI6IjEwMDk5M2FiLTg0NDQtZTE1Ny1kMDAyLTU4M2VlYmVkMjc1YiIsImFsZyI6ImhzMjU2IiwiZW52IjpudWxsLCJleHAiOiI3MjAwIn0.eyJ1aWQiOiJsaWNoYW8ifQ.abe91dda193714a4eea55ba8360fd2b5779ca190d9165a33c33a722cfff4c75c"
var secretKey = "HelloWorld"

func TestSignature(t *testing.T) {
	x := Signature(header, payload, secretKey, "hs256")
	t.Log(x)
}

func TestHasSign(t *testing.T) {
	x := HasSign(token)
	t.Log(x)
	y := HasSign(header)
	t.Log(y)
}

func TestVerifySign(t *testing.T) {
	x := VerifySign(token, secretKey, "hs256")
	t.Log(x)
	y := VerifySign(header, secretKey)
	t.Log(y)
}

func TestEncode(t *testing.T) {
	data := map[string]interface{}{
		"uid": "LiChao",
	}
	str, _ := Encode(data)
	t.Log(str)
}

func TestDecode(t *testing.T) {
	data, _ := Decode(payload)
	t.Log(data)
}

func TestAssemble(t *testing.T) {
	str := Assemble(header, payload, sign)
	t.Log(str)
}

func TestDisassemble(t *testing.T) {
	h, p, s := Disassemble(token)
	t.Log(h, p, s)
}
