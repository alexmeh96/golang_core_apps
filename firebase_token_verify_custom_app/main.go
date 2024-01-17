package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"log"
	"net/http"
	"time"
)

const token = ""

// публичные ключи
var publicKeys = map[string]*rsa.PublicKey{}

// ресурс, который возвращает сертификаты
var CertsAPIEndpoint = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Mount("/private", privateRouter())
	r.Mount("/public", publicRouter())

	println("server started")
	http.ListenAndServe(":8085", r)

}

// получение сертификатов из удалённого сервиса и считытвание из них публичных ключей
func getCertificates() (certs map[string]*rsa.PublicKey, err error) {
	res, err := http.Get(CertsAPIEndpoint)
	if err != nil {
		return
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	var c map[string]string
	json.Unmarshal(data, &c)

	certs = make(map[string]*rsa.PublicKey)

	for key := range c {
		certs[key], _ = readPublicKey(c[key])
	}

	return
}

// получение публичного ключа по kid из jwt. Если его нет у нас локально,
// то происходит запрос на удалённый сервис, чтобы получить сертификаты и вытащить из них публичные ключи
func getPublicKey(kid string) (cert *rsa.PublicKey, err error) {
	if publicKeys[kid] == nil {
		certs, e := getCertificates()
		if e != nil {
			return
		}
		publicKeys = certs
	}

	cert = publicKeys[kid]
	if cert == nil {
		err = errors.New("key not found")
	}

	return
}

// получение публичное ключа из сертифика
func readPublicKey(cert string) (*rsa.PublicKey, error) {
	publicKeyBlock, _ := pem.Decode([]byte(cert))
	if publicKeyBlock == nil {
		return nil, errors.New("invalid public key data")
	}
	if publicKeyBlock.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("invalid public key type: %s", publicKeyBlock.Type)
	}
	c, err := x509.ParseCertificate(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := c.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not RSA public key")
	}
	return publicKey, nil
}

// проверка токена
func verifyToken(token string) (*jwt.Token, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		kid, ok := jwtToken.Header["kid"]
		if !ok {
			return []byte{}, errors.New("kid not found")
		}
		kidString, ok := kid.(string)
		if !ok {
			return []byte{}, errors.New("kid cast error to string")
		}
		return getPublicKey(kidString)
	})

	return tok, err
}

func privateRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(AuthMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value("token").(string)
		w.Write([]byte(token))
	})
	return r
}

func publicRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(Logging)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("publicRouter"))
	})
	return r
}

// проверка токена
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if tok, err := verifyToken(token); err == nil {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "token", tok.Raw)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		}
	})
}

// собственный middleware, для оценки длительности выполнения запроса
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
