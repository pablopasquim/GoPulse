package routes

import (
	"log"
	"net/http"

	"github.com/pablopasquim/GoPulse/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
