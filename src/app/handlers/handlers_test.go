package handlers

import (
	"bytes"
	"testing"
	"net/http"
    "net/http/httptest"
    st"app/structs"
    "github.com/gorilla/mux"
    "github.com/tkanos/gonfig"
)


func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

func TestEncryptHandlerSuccess(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"data":"test function"}`)
    req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", EncryptHandler).Methods("POST")

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)

}

func TestEncryptHandlerFailJson(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"something":"test product"}`)
    req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", EncryptHandler).Methods("POST")

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)

}

func TestEncryptHandlerFailMethod(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"something":"test product"}`)
    req, _ := http.NewRequest("GET", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", EncryptHandler).Methods("POST")

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)

}

func TestDecryptHandlerSuccess(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"data":"LiL_wkoqaHJVG8w3M6M8oZn0evtxXD0WpbpMgbQ="}`)
    req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", DecryptHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)

}

func TestDecryptHandlerFailText(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"data":"wdhfsdhfiwf8347"}`)
    req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", DecryptHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)
}

func TestDecryptHandlerFailJson(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"something":"LiL_wkoqaHJVG8w3M6M8oZn0evtxXD0WpbpMgbQ="}`)
    req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", DecryptHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)
}

func TestDecryptHandlerFailMethod(t *testing.T) {
    gonfig.GetConf("../configuration/conf.json", &st.Config)
    var jsonStr = []byte(`{"something":"wdhfsdhfiwf8347"}`)
    req, _ := http.NewRequest("GET", "/api/encrypt", bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    r := mux.NewRouter()
	r.HandleFunc("/api/encrypt", DecryptHandler)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    r.ServeHTTP(rr, req)
    checkResponseCode(t, http.StatusOK , rr.Code)
}