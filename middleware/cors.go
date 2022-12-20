package middleware

import "net/http"

type CORSMiddleware struct {
	Next http.Handler
}

func (cm *CORSMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cm.Next == nil {
		cm.Next = http.DefaultServeMux
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cm.Next.ServeHTTP(w, r)
}
