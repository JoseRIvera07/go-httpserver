package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"crypto/x509"
	st"./structs"
	rt"./handlers"
	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
)

func main() {
	//Load struct configuration
	gonfig.GetConf("./configuration/conf.json", &st.Config)
	//Use gorilla mux to route handlers
	r := mux.NewRouter()
	r.HandleFunc(st.Config.EncryptRoute, rt.EncryptHandler)
	r.HandleFunc(st.Config.DecryptRoute, rt.DecryptHandler)
	http.Handle("/", r)
	// Create CA certificate
	caCert, err := ioutil.ReadFile("cert.pem") //add cert.pem to it
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Enable Client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs: caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	// Server instance 
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}
	
	// Set HTTPS connections with the server instance  with the certificate 
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}