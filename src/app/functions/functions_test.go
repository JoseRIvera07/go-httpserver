package functions

import (
	"testing"
	st "../structs"
	"github.com/tkanos/gonfig"
)
// test encrypt success
func TestEncryptSucess(t *testing.T){
	gonfig.GetConf("../configuration/conf.json", &st.Config)
	key := []byte(st.Config.Key)
	resultEncrypText := Encrypt(key, "TEST")
	if resultEncrypText == "-1"{
		t.Errorf("Expected response code %s. Got %s\n", "json{}", "-1")
	}
}
// test encrypt fail when key lenght is too short
func TestEncryptFail_Key(t *testing.T){
	key := []byte("")
	resultEncrypText := Encrypt(key, "TEST")
	if resultEncrypText == "-1"{
		t.Errorf("Expected response code %s. Got %s\n", "json{}", "-1")
	}
}
// test decrypt success
func TestDecryptSucess(t *testing.T){
	key := []byte(st.Config.Key)
	resultDecryptText := Decrypt(key, "eL8mP2if-kU-98RaW4qnLvGsN5U=")
	
	if resultDecryptText == "-1"{
		t.Errorf("Expected response code %s. Got %s\n", "json{}", "-1")
	}
}
//test decrypt fail when key is too short
func TestDecryptFail_Key(t *testing.T){
	key := []byte("")
	resultDecryptText := Decrypt(key, "eL8mP2if-kU-98RaW4qnLvGsN5U=")
	
	if resultDecryptText == "-1"{
		t.Errorf("Expected response code %s. Got %s\n", "json{}", "-1")
	}
}
// test decrypt fail, when text encrypted given does not show to be correct encrypted
func TestDecryptFail_Decrypt(t *testing.T){
	key := []byte(st.Config.Key)
	resultDecryptText := Decrypt(key, "dhdcqwhfhqwdfhb")
	
	if resultDecryptText == "-1"{
		t.Errorf("Expected response code %s. Got %s\n", "json{}", "-1")
	}
}