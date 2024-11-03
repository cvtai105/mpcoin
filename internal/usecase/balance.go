package usecase

import (
	"context"
	"errors"
	"mpc/internal/domain"
	"mpc/internal/repository"

	"github.com/google/uuid"
)

type BalanceUseCase interface {
	GetBalancesByWalletId(ctx context.Context, walletID uuid.UUID, userId uuid.UUID) ([]domain.GetBalanceResponse, error)
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

func (uc *balanceUseCase) GetBalancesByWalletId(ctx context.Context, walletID uuid.UUID, userId uuid.UUID) ([]domain.GetBalanceResponse, error) {
	// Check if wallet belongs to user
	wallet, err := uc.walletRepo.GetWallet(ctx, walletID)
	// return error if wallet does not exist
	if err != nil {
		return nil, errors.New("invalid wallet id")
	}

	// return error if wallet does not belong to 
	if wallet.UserID != userId {
		return nil, errors.New("Forbidden")
	}

	// get balances
	return uc.balanceRepo.GetBalancesByWalletId(ctx, walletID)
}

