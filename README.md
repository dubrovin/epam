# epam

```curl -X GET http://127.0.0.1:8080/products```
```curl -X GET http://127.0.0.1:8080/products/1/reserve```
```
curl -X POST \
  http://127.0.0.1:8080/reserves/accept \
  -H 'content-type: application/json' \
  -d '{
    "hash": "4debf6c1-a031-4eae-a4b3-2e47c600128e"
}'
```