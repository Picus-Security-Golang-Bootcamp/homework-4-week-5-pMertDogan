package admin

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AdminHandler(r *mux.Router) {
	// create mux
	// r.Use(AdminMiddleware)			 
	http.HandleFunc("/admin/drop", DropTables)
	http.HandleFunc("/admin/init", InitDatabase)

}

