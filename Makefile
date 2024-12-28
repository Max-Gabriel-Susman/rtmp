
deps: 
	go get "github.com/golang-jwt/jwt" "github.com/AgustinSRG/go-tls-certificate-loader"
	go mod tidy 
	go mod vendor 