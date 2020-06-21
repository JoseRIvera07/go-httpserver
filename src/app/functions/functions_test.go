package functions

import (
	"testing"
	st "app/structs"
	"github.com/tkanos/gonfig"
)

func TestEncryptSucess(t *testing.T){
	gonfig.GetConf("../configuration/conf.json", &st.Config)
	key := []byte(st.Config.Key)
	resultEncrypText := Encrypt(key, "TEST")
	if resultEncrypText == "-1"{
		t.Errorf("Expected response code %s. Got %s\n", "json{}", "-1")
	}
}
func TestEncryptFail_Key(t *testing.T){
	key := []byte("")
	resultEncrypText := Encrypt(key, "TEST")
	if resultEncrypText == "-1"{
		t.Errorf("Expected response code %s. Got %s\n", "json{}", "-1")
	}
}
func TestDecryptSucess(t *testing.T){
	key := []byte(st.Config.Key)
	resultDecryptText := Decrypt(key, "qAO7qQAjYwj6IJ8alB-g9GBr08MeuLg=")
	
	if resultDecryptText == "-1"{
		t.Errorf("Expected response code %s. Got %s\n", "json{}", "-1")
	}
}
func TestDecryptFail_Key(t *testing.T){
	key := []byte("")
	resultDecryptText := Decrypt(key, "qAO7qQAjYwj6IJ8alB-g9GBr08MeuLg=")
	
	if resultDecryptText == "-1"{
		t.Errorf("Expected response code %s. Got %s\n", "json{}", "-1")
	}
}

func TestDecryptFail_Decrypt(t *testing.T){
	key := []byte(st.Config.Key)
	resultDecryptText := Decrypt(key, "dhdcqwhfhqwdfhb")
	
	if resultDecryptText == "-1"{
		t.Errorf("Expected response code %s. Got %s\n", "json{}", "-1")
	}
}