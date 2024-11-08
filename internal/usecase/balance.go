package usecase

import (
	"context"
	"fmt"
	"mpc/internal/domain"
	"mpc/internal/repository"

	"github.com/google/uuid"
)

type BalanceUseCase interface {
	GetBalancesByUserId(ctx context.Context, userId uuid.UUID) ([]domain.GetBalanceResponse, error)
}

type balanceUseCase struct {
	balanceRepo repository.BalanceRepository
	walletRepo repository.WalletRepository
}

func NewBalanceUC(balanceRepo repository.BalanceRepository, walletRepo repository.WalletRepository) BalanceUseCase {
	return &balanceUseCase{
		balanceRepo: balanceRepo,
		walletRepo: walletRepo,	
	}
}

var _ BalanceUseCase = (*balanceUseCase)(nil)

func (uc *balanceUseCase) GetBalancesByUserId(ctx context.Context, userId uuid.UUID) ([]domain.GetBalanceResponse, error) {
	// get balances
	result, err := uc.balanceRepo.GetBalancesByUserId(ctx, userId)
	if err != nil {
		fmt.Println("balanceUseCase.GetBalancesByUserId error 1: ", err.Error()) 
		return []domain.GetBalanceResponse{} , err
	}

	return result, nil	
}
