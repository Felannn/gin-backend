package handlers

import (
	"net/http"

	"github.com/Felannn/gin-backend.git/models"
	"github.com/Felannn/gin-backend.git/services"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	txService *services.TransactionService
}

func NewTransactionHandler() *TransactionHandler {
	return &TransactionHandler{txService: services.NewTransactionService()}
}

func (h *TransactionHandler) Checkout(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak ditemukan"})
		return
	}

	var req models.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Metode pembayaran wajib diisi"})
		return
	}

	tx, err := h.txService.Checkout(userID, req.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Checkout berhasil",
		"data":    tx,
	})
}

func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "User ID tidak ditemukan"})
		return
	}

	transactions, err := h.txService.GetTransactions(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Gagal mengambil riwayat transaksi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    transactions,
	})
}
