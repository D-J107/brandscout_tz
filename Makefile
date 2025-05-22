.PHONY: curl_post
curl_post:
	curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'

.PHONY: curl_get_all
curl_get_all:
	curl http://localhost:8080/quotes
	
.PHONY: curl_get_random
curl_get_random:
	curl http://localhost:8080/quotes/random

.PHONY: curl_get_by_author
curl_get_by_author:
	curl http://localhost:8080/quotes?author=Confucius

.PHONY: curl_delete
curl_delete:
	curl -X DELETE http://localhost:8080/quotes/1

.PHONY: run_curl_tests
run_curl_tests: curl_post curl_get_all curl_get_random curl_get_by_author curl_delete

