# **Dependencies**
Run:

$go get github.com/gorilla/mux \ && go get github.com/tkanos/gonfig 
    
# **Build**
Inside $go-httpServer/src/app: **go build -o main .**

**I'm using relative paths, so take care of configure your GOPATH first.**

My GOPATH, looks as follows:

**export GOPATH=$HOME/go-httpserver/src/app**

# **Run**
## **IP & PORT**
**IP:** localhost
**PORT:** 8443
**URL:** https://localhost:8443/
## **Run app**
Inside $go-httpServer/src/app: **./main**

Take care of configurate your client with the certificate and key, because the server was configure to use https method. 
Inside $go-httpServer/src/app you're going to find cer.pem and key.pem, certificate and key correspondingly.
To test this app, I used postman, here is a screenshot of my certificate configuration.
![](./img/postman_ca.png)
![](./img/postman_test_encrypt.png)
![](./img/postman_test_decrypt.png)

# **JSON**

This is how the json must be to send de request to the server:
**{"data": "here goes your string"}**

# **Response**
The go-httpserver is going to response with a json like this:

**{
    "code": 200,
    "message": "Success",
    "data": "oBtIaikeTQNrpF0Ag0j6NIEtbH4="
}**


# **Unit Test**

## **Execute functions unit test**

Inside $go-httpServer/src/app/functions: **go test -v --cover**

![](./img/test_functions.png)

## **Execute handlers unit test**

Inside $go-httpServer/src/app/handlers: **go test -v --cover**

![](./img/test_handlers.png)

# **DOCKER** 
## **On linux:**
### **To create the container:**
Inside $go-httpServer/src/app: **sudo docker build -t go-httpserver .**

### **To run the container:**
**sudo docker run -d -p 8443:8443 go-httpserver:latest**


