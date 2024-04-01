package product_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	productcontroller "github.com/golang-crud/controllers"
	"github.com/golang-crud/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setup() (*gin.Engine, *gorm.DB, error) {
	db, err := models.ConnectDatabase()
	if err != nil {
		return nil, nil, err
	}

	db.Exec("DELETE FROM products")

	r := gin.Default()
	return r, db, nil
}

func teardown(db *gorm.DB)  {
	

}

func TestIndex(t *testing.T) {
	r, db, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer teardown(db)

	products := []models.Product{
		{
			NamaProduct: "Product 1",
			Deskripsi:   "Deskripsi Product 1",
		},
		{
			NamaProduct: "Product 2",
			Deskripsi:   "Deskripsi Product 2",
		},
	}
	for _, product := range products {
		db.Create(&product)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/products", nil)
	r.GET("/api/products", productcontroller.Index)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Product
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, products, response)
}

func TestShow(t *testing.T) {
	r, db, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer teardown(db)

	product := models.Product{
		NamaProduct: "Product 1",
		Deskripsi:   "Deskripsi Product 1",
	}
	db.Create(&product)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/product/1", nil)
	r.GET("/api/product/:id", productcontroller.Show)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Product
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	assert.Equal(t, product, response)
}

func TestCreate(t *testing.T) {
	r, db, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer teardown(db)

	newProduct := models.Product{
		NamaProduct: "Product Baru",
		Deskripsi:   "Deskripsi Product Baru",
	}

	body, _ := json.Marshal(newProduct)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/product", bytes.NewReader(body))
	r.POST("/api/product", productcontroller.Create)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestUpdate(t *testing.T) {
	r, db, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer teardown(db)

	product := models.Product{
		NamaProduct: "Product 1",
		Deskripsi:   "Deskripsi Product 1",
	}
	db.Create(&product)

	updatedProduct := models.Product{
		NamaProduct: "Product 1 Diubah",
		Deskripsi:   "Deskripsi Product 1 Diubah",
	}

	body, _ := json.Marshal(updatedProduct)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/product/1", bytes.NewReader(body))
	r.PUT("/api/product/:id", productcontroller.Update)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var updatedProductDB models.Product
	db.First(&updatedProductDB, 1)
	assert.Equal(t, updatedProduct.NamaProduct, updatedProductDB.NamaProduct)
	assert.Equal(t, updatedProduct.Deskripsi, updatedProductDB.Deskripsi)
}

func TestDelete(t *testing.T) {
	r, db, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer teardown(db)

	product := models.Product{
		NamaProduct: "Product 1",
		Deskripsi:   "Deskripsi Product 1",
	}
	db.Create(&product)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/product/1", nil)
	r.DELETE("/api/product/:id", productcontroller.Delete)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var deletedProduct models.Product
	err = db.First(&deletedProduct, 1).Error
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
