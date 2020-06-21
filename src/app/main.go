package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"crypto/x509"
	st"app/structs"
	rt "app/handlers"
	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
)

func main() {
	gonfig.GetConf("configuration/conf.json", &st.Config)

	r := mux.NewRouter()
	r.HandleFunc(st.Config.EncryptRoute, rt.EncryptHandler)
	r.HandleFunc(st.Config.DecryptRoute, rt.DecryptHandler)
	http.Handle("/", r)
	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create the TLS Config with the CA pool and enable Client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs: caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	// Create a Server instance to listen on port 8443 with the TLS config
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}
	
	// Listen to HTTPS connections with the server certificate and wait
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}