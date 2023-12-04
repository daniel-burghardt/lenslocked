package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>This is my Go majestic app! Wohoo!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:daniel@email.com\">daniel@email.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<h1>FAQ</h1>
		<div>
			<div>Q: Is there a free version?</div>
			<div>A: Yes! We offer a free trial for 30 days on any paid plans.</div>
		</div>
		</br>
		<div>
			<div>Q: What are your support hours?</div>
			<div>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</div>
		</div>
		</br>
		<div>
			<div>Q: How do I contact support?</div>
			<div>A: Email us - support@lenslocked.com</div>
		</div>
	`)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func main() {
	router := Router{}
	fmt.Println("Starting the server on port 3000...")
	http.ListenAndServe("localhost:3000", router)
}
