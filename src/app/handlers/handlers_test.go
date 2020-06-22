package Handlers

import (
	"bytes"
	"testing"
	"net/http"
    "net/http/httptest"
    st"../structs"
    "github.com/gorilla/mux"
    "github.com/tkanos/gonfig"
)

// func to check and compare expected response with given response
func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

// test encrypt handler succes
func TestEncryptHandlerSuccess(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"data":"test function"}`)
    req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", EncryptHandler).Methods("POST")
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)

}
// test encrypt handler fail because json format given.
func TestEncryptHandlerFailJson(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"something":"test product"}`)
    req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", EncryptHandler).Methods("POST")
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)

}
// test encrypt handler fail because json wrong method given
func TestEncryptHandlerFailMethod(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"something":"test product"}`)
    req, _ := http.NewRequest("GET", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", EncryptHandler).Methods("POST")
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)

}
// test decrypt handler succes
func TestDecryptHandlerSuccess(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"data":"eL8mP2if-kU-98RaW4qnLvGsN5U="}`)
    req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", DecryptHandler)
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)

}
// test decrypt handler fail when text encrypted given does not show to be correct encrypted
func TestDecryptHandlerFailText(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"data":"wdhfsdhfiwf8347"}`)
    req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", DecryptHandler)
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)
}
// test decrypt handler fail because json format given.
func TestDecryptHandlerFailJson(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"something":"eL8mP2if-kU-98RaW4qnLvGsN5U="}`)
    req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", DecryptHandler)
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)
}
// test decrypt handler fail because json wrong method given
func TestDecryptHandlerFailMethod(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"something":"wdhfsdhfiwf8347"}`)
    req, _ := http.NewRequest("GET", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", DecryptHandler)
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)
}