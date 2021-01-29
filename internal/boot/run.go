package boot

import (
	"github.com/naumyegor/taxi-service/internal/services"
	"log"
	"net/http"
)

func Run() {
	r := services.InitRouter()

	http.Handle("/", r)

	err := http.ListenAndServe(":6006", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
