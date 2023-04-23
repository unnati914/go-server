package main

import(
	"fmt"
	"log"
	"net/http"
)
func formHandeler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w,"POST request Succesful")
	name := r.FormValue("name")
	fmt.Fprintf(w,"Name = %s\n",name)
}
func helloHandeler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"method is not supported",http.StatusNotFound)
	}
	fmt.Fprintf(w,"hello!")

}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandeler)
	http.HandleFunc("/hello",helloHandeler)

	fmt.Printf("Starting server")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}

}