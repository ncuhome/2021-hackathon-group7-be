package util

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"os"
	"strconv"
)

func WriteJSON(path string, v interface{}) error {
	outData, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, outData, 0755)
	if err != nil {
		return err
	}
	return nil
}

// v传入引用
func ReadJSON(path string, v interface{}) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, v)
	if err != nil {
		return err
	}
	return nil
}

//if exist,return true
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func RandHexStr(byteNum uint) (string, error) {
	bytes := make([]byte, byteNum)
	_, err := rand.Reader.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func RandDecStr(num uint) (string, error) {
	bytes := make([]byte, num)
	_, err := rand.Reader.Read(bytes)
	if err != nil {
		return "", err
	}
	str := ""
	for _, v := range bytes {
		str = str + string('0'+v%10)
	}
	return str, nil
}

func MD5(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}

func SHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func SHA512(data []byte) []byte {
	h := sha512.New()
	h.Write(data)
	return h.Sum(nil)
}

func Bcrypt(data []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, cost)
}

func StringAdd(data string) (string, error) {
	i, err := strconv.Atoi(data)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(i + 1), nil
}
