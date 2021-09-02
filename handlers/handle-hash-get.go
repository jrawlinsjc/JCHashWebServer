package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jmrawlins/JCHashWebServer/datastore/hashdatastore"
	"github.com/jmrawlins/JCHashWebServer/hash"
)

type HashGetHandler struct {
	Ds hashdatastore.HashDataStore
}

func (handler HashGetHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// Get the id from the uri
	strUri := strings.TrimLeft(req.URL.Path, "/")
	log.Println("Received request at:", strUri)

	// Get endpoint as an int, if possible
	if len(strUri) == 0 {
		fmt.Fprintf(resp, "Up and running...")
	} else if hashId, err := hash.HashIdFromString(strUri); err != nil {
		log.Println("404")
	} else {
		var hashValue string
		var err error
		if hashValue, err = handler.Ds.GetHash(hashId); err != nil {
			fmt.Fprint(resp, "404 hash not defined for ", hashId)
		} else {
			log.Printf("{ \"id\":%d, \"hash\":\"%s\"\n", hashId, hashValue)
			fmt.Fprint(resp, hashValue)
		}
	}
}
