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
// @Param wallet_id path string true "Wallet ID"
// @Success 200 {object} docs.GetBalancesResponse "Successful response"
// @Failure 400 {string} string "Bad request error due to invalid input"
// @Failure 500 {string} string "Internal server error"
// @Security ApiKeyAuth
// @Router /balances/{wallet_id} [get]
func (h *BalanceHandler) GetBalances(c *gin.Context) {
	//get wallet id from query parameter
	walletId := uuid.MustParse(c.Param("wallet_id"))
	//check if wallet id is valid
	if walletId == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet id"})
		return
	}

	//get user id from context
	userId := c.MustGet("user_id").(uuid.UUID)

	//get balances
	balances, err := h.blcUC.GetBalancesByWalletId(c.Request.Context(), walletId, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balances": balances})
}
