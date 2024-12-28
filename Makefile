
deps: 
	go get "github.com/golang-jwt/jwt"
	go mod tidy 
	go mod vendor 

clean: 
	rm -r vendor
	rm go.sum