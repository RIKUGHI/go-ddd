package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/rikughi/go-ddd/aggregate"
	"github.com/rikughi/go-ddd/domain/product"
)

type MemoryRepositoryProduct struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryRepositoryProduct {
	return &MemoryRepositoryProduct{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

func (mpr *MemoryRepositoryProduct) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

func (mpr *MemoryRepositoryProduct) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[uuid.UUID(id)]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryRepositoryProduct) Add(newprod aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	mpr.products[newprod.GetID()] = newprod

	return nil
}

func (mpr *MemoryRepositoryProduct) Update(upprod aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[upprod.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

func (mpr *MemoryRepositoryProduct) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
