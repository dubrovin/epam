# epam

```curl -X GET http://127.0.0.1:8081/products```
```curl -X GET http://127.0.0.1:8082/products/1/reserve```
```
curl -X POST \
  http://127.0.0.1:8082/reserves/accept \
  -H 'content-type: application/json' \
  -d '{
    "hash": "4debf6c1-a031-4eae-a4b3-2e47c600128e"
}'
```
docker db  
```
docker build ./ -t 'db:epam'  
docker tag db:epam dubrovin/epam-db
docker push dubrovin/epam-db 
```

docker watcher  
```
docker build ./ -t 'watcher:epam'  
docker tag watcher:epam dubrovin/epam-watcher
docker push dubrovin/epam-watcher
```
