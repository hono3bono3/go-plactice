package handler

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world")
	log.Info().Msg("receive hello world request")
}
