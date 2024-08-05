package scarlet

import (
	"log"
	"net/http"
)

func (s *Scarlet) Listen(port string) {
	createRoutesTree(s.routes, s.config)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
