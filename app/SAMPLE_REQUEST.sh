curl --location 'http://localhost:8080/v1/jwt' \
--header 'Content-Type: application/json' \
--data '{
  "client-secret": "my-secret",
  "client-id": "my-id"
}'