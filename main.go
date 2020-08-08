package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	var dumpPath, query string
	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	log.Println("Starting simplefts")

	start := time.Now()
	docs, err := loadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(index)
	idx.add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchedIDs := idx.search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
