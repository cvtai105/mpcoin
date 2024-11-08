package handler

import (
	"mpc/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type BalanceHandler struct {
	blcUC usecase.BalanceUseCase
}

func NewBalanceHandler(blcUC usecase.BalanceUseCase) *BalanceHandler {
	return &BalanceHandler{blcUC: blcUC}
}


// GetBalances godoc
// @Summary Get balances by wallet id
// @Description Get a list of balances by wallet id
// @Tags balance
// @Accept json
// @Produce json
// @Success 200 {object} docs.GetBalancesResponse "Successful response"
// @Failure 400 {string} string "Bad request error due to invalid input"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /balances/{wallet_id} [get]
func (h *BalanceHandler) GetBalances(c *gin.Context) {
	//get user id from context
	userId := c.MustGet("userID").(uuid.UUID)
	if userId == uuid.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	//get balances
	balances, err := h.blcUC.GetBalancesByUserId(c.Request.Context(), userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balances": balances, "user": userId, "status": "200"})
}
