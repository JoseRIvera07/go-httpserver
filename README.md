

Build
Inside $go-httpServer/src/app: go build -o main .

Run
./main
Take care of configurate your client with the certificate and key, because the server was configure to use https method. 
Inside $go-httpServer/src/app you're going to find cer.pem and key.pem, certificate and key correspondingly.
![](./img/postman_ca.png)

Unit Test

Execute functions unit test
Inside $go-httpServer/src/app/functions: go test -v --cover

Execute handlers unit test
Inside $go-httpServer/src/app/handlers: go test -v --cover
