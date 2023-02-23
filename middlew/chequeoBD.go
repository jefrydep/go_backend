package middlew

import (
	"net/http"
	"github.com/jefrydep/go_backend/bd"
)

func CheckeoBD(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckedConnection() == 0{
			http.Error(w,"conexion perdida con la base de dato",500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
