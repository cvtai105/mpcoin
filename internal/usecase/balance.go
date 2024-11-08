package usecase

import (
	"context"
	"fmt"
	"mpc/internal/domain"
	"mpc/internal/repository"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type BalanceUseCase interface {
	GetBalancesByUserId(ctx context.Context, userId uuid.UUID) ([]domain.GetBalanceResponse, error)
	UpdateBalanceRPC(ctx context.Context, address common.Address, tokenId uuid.UUID) error
}

type balanceUseCase struct {
	balanceRepo repository.BalanceRepository
	walletRepo repository.WalletRepository
	chainRepo repository.ChainRepository
	ethRepo   repository.EthereumRepository
}

func NewBalanceUC(balanceRepo repository.BalanceRepository, walletRepo repository.WalletRepository, chainRepo repository.ChainRepository, ethRepo repository.EthereumRepository) BalanceUseCase {
	return &balanceUseCase{
		balanceRepo: balanceRepo,
		walletRepo:  walletRepo,
		chainRepo:   chainRepo,
		ethRepo:     ethRepo,
	}
}

var _ BalanceUseCase = (*balanceUseCase)(nil)

func (uc *balanceUseCase) GetBalancesByUserId(ctx context.Context, userId uuid.UUID) ([]domain.GetBalanceResponse, error) {
	// get balances
	result, err := uc.balanceRepo.GetBalancesByUserId(ctx, userId)
	if err != nil {
		fmt.Println("balanceUseCase.GetBalancesByUserId error 1: ", err.Error())
		return []domain.GetBalanceResponse{}, err
	}

	return result, nil
}

// UpdateBalance implements BalanceUseCase.
func (uc *balanceUseCase) UpdateBalanceRPC(ctx context.Context, address common.Address, tokenId uuid.UUID) error {

	//get chain by token id
	//todo: refactor: create a new method to get chain by token id
	chains, err := uc.chainRepo.GetChains(ctx)
	if err != nil {
		return err
	}
	foundIndex := -1
	for i, chain := range chains {
		if chain.NativeTokenID == tokenId {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return fmt.Errorf("chain not found for token id: %s", tokenId)
	}
	chain := chains[foundIndex]

	ethRepo, err2 := uc.ethRepo.NewInstance(chain.RPCURL)
	if err2 != nil {
		return err2
	}

	//get balance by address
	balance, err2 := ethRepo.GetBalance(address)
	if err2 != nil {
		return err2
	}

	if balance == nil {
		fmt.Println("address is: ", address)
		return fmt.Errorf("balance is nil")
	}else{
		fmt.Println("balance is: ", balance)
	}
	

	param := domain.UpdateBalanceParams{
		Address: address,
		TokenID: tokenId,
		Balance: balance,
	}

	fmt.Println("param: ", param)

	//update balance
	err = uc.balanceRepo.UpdateBalance(ctx, param)
	if err != nil {
		return err
	}

	return nil
}
