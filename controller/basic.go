package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

func Basic() {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["name"] = "Batman"
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := make(map[string]interface{})
		data["name"] = "Superman"

		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/vulncmd", func(w http.ResponseWriter, r *http.Request) {
		keys := r.URL.Query()["key"][0]

		awsKey := "${{ secrets.AWS_KEY }}"
		awsSecret := "${{ secrets.AWS_SECRET }}"
		fmt.Println(awsKey, awsSecret)

		 if !ok || len(keys) < 1 {
		 	fmt.Println("Url Param 'key' is missing")
		 	return
		 }

		key := keys[0]

		fmt.Println("Url Param 'key' is: " + string(key[0]))
		cmd := exec.Command("/bin/sh", "-c", string(key[0]))
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Fprintf(w, string(stdout))
	})
}
