package posts

import (
	"net/http"
	"encoding/json"
)

func List(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetPosts())
}

func GetOne(w http.ResponseWriter, r *http.Request, id int) {
	json.NewEncoder(w).Encode(GetPost(id))
}
