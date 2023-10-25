package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>This is contact page</h1> "+
		"<p>For further info please contact me at "+
		"<a href=\"mailto:nanasuryana335@gmail.com\">nanasuryana335@gmail.com </a>"+
		"</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<div>"+
		"<p>Q: Is there a free version? </p>"+
		"<p>A: Yes! We offer a free trial version for 30 days on any paid plans</p>"+
		"</div>"+
		"<div>"+
		"<p>Q: What are your support hours?</p>"+
		"<p>A: We have support staff answering emails 24/7, though response times may be a bit slower on weekends</p>"+
		"</div>"+
		"<div>"+
		"<p>Q: How do i contact support?"+
		"<p>A: Email us - <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a></p>"+
		"</div>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
		break
	case "/contact":
		contactHandler(w, r)
		break
	case "/faq":
		faqHandler(w, r)
	default:
		//TODO:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

//type Router struct{}
//
//func (rt Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//		break
//	case "/contact":
//		contactHandler(w, r)
//		break
//	default:
//		//TODO:
//		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//	}
//}

func main() {
	http.Handle("/", http.HandlerFunc(pathHandler))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		return
	}
}
