# epam

### How to run

1) ```$ docker-compose up```
2) build each service by Dockerfile with changing params
3) local run each service

### How to use

returns all product:  
```$ curl -X GET http://127.0.0.1:8081/products```  

reserve product by id:  
 ```$ curl -X GET http://127.0.0.1:8082/products/1/reserve```
 
 accepted reservation by hash:    
```
$ curl -X POST \
  http://127.0.0.1:8082/reserves/accept \
  -H 'content-type: application/json' \
  -d '{
    "hash": "4debf6c1-a031-4eae-a4b3-2e47c600128e"
}'
```  



additional info, how to update images:  

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

docker reserver  
```
docker build ./ -t 'reserver:epam'  
docker tag reserver:epam dubrovin/epam-reserver
docker push dubrovin/epam-reserver
```