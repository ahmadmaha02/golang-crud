package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-crud/models"
	"github.com/golang-crud/repositories"
)

type ProductController struct {
    repo *repositories.ProductRepository
}

func NewProductController(repo *repositories.ProductRepository) *ProductController {
    return &ProductController{repo}
}

func (ctrl *ProductController) Index(c *gin.Context) {
    products, err := ctrl.repo.GetAllProducts()
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"products": products})
}

func (ctrl *ProductController) Show(c *gin.Context) {
    id := c.Param("id")
    product, err := ctrl.repo.GetProductByID(id)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"product": product})
}

func (ctrl *ProductController) Create(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    if err := ctrl.repo.CreateProduct(&product); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"product": product})
}

func (ctrl *ProductController) Update(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }
    existingProduct, err := ctrl.repo.GetProductByID(id)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
        return
    }

    product.Id = existingProduct.Id
    if err := ctrl.repo.UpdateProduct(&product); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}
func (ctrl *ProductController) Delete(c *gin.Context) {
    id := c.Param("id")
    product, err := ctrl.repo.GetProductByID(id)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
        return
    }
    if err := ctrl.repo.DeleteProduct(&product); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
