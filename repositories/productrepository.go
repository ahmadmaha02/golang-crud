package repositories

import (
	"github.com/golang-crud/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{db}
}

func (repo *ProductRepository) GetAllProducts() ([]models.Product, error) {
    var products []models.Product
    if err := repo.db.Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}

func (repo *ProductRepository) GetProductByID(id string) (models.Product, error) {
    var product models.Product
    if err := repo.db.First(&product, id).Error; err != nil {
        return models.Product{}, err
    }
    return product, nil
}

func (repo *ProductRepository) CreateProduct(product *models.Product) error {
    if err := repo.db.Create(product).Error; err != nil {
        return err
    }
    return nil
}

func (repo *ProductRepository) UpdateProduct(product *models.Product) error {
    if err := repo.db.Save(product).Error; err != nil {
        return err
    }
    return nil
}

func (repo *ProductRepository) DeleteProduct(product *models.Product) error {
    if err := repo.db.Delete(product).Error; err != nil {
        return err
    }
    return nil
}