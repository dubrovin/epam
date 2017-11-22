package services

import (
	"errors"
	"github.com/dubrovin/epam/models"
	"github.com/satori/go.uuid"
	"log"
	"sync"
	"time"
)

type DataBase struct {
	currentID         int64
	availableProducts map[int64]*models.Product
	reservedProducts  map[string]*models.Product
	acceptedProducts  map[string]*models.Product
	mu                sync.RWMutex
}

func NewDataBase() *DataBase {
	return &DataBase{
		availableProducts: make(map[int64]*models.Product, 1000),
		reservedProducts:  make(map[string]*models.Product),
		acceptedProducts:  make(map[string]*models.Product),
		currentID:         1,
	}
}

func (db *DataBase) GetProducts() map[int64]*models.Product {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.availableProducts
}

func (db *DataBase) AddProduct(product *models.Product) (int64, error) {
	err := product.SetID(db.currentID)
	if err != nil {
		return 0, err
	}
	db.currentID++

	db.mu.Lock()
	defer db.mu.Unlock()
	db.availableProducts[product.GetID()] = product

	return product.GetID(), nil
}

func (db *DataBase) ReserveProduct(productID int64, ttl time.Duration) (string, error) {

	if _, ok := db.availableProducts[productID]; ok {
		//do something here
		hash := uuid.NewV4().String()
		db.mu.Lock()
		defer db.mu.Unlock()
		db.reservedProducts[hash] = db.availableProducts[productID]
		db.reservedProducts[hash].SetHash(hash)
		db.reservedProducts[hash].SetTTL(ttl)
		delete(db.availableProducts, productID)
		log.Printf("Reserved product with hash %s", hash)
		return hash, nil
	} else {
		return "", errors.New("product not found")
	}

}

func (db *DataBase) AcceptReserve(hash string) (bool, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.reservedProducts[hash] = db.acceptedProducts[hash]
	delete(db.acceptedProducts, hash)
	log.Printf("Accepted reserve for product with hash %s", hash)
	return true, nil
}

func (db *DataBase) cancelReserve(hash string) {
	product := db.reservedProducts[hash]
	db.availableProducts[product.GetID()] = product
	delete(db.reservedProducts, hash)
	log.Printf("Canceled reserve for product with hash %s", hash)
}

func (db *DataBase) Checker(sleepTime time.Duration) {
	for {
		log.Print("Start check")
		db.mu.Lock()
		for _, product := range db.reservedProducts {
			if time.Now().Add(product.GetTTL()).Unix() > time.Now().Unix() {
				db.cancelReserve(product.GetHash())
			}

		}
		db.mu.Unlock()
		log.Print("Stop check")
		time.Sleep(sleepTime)
	}
}
