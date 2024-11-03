package postgres

import (
	"context"
	"mpc/internal/domain"
	sqlc "mpc/internal/infrastructure/db/sqlc"
	"mpc/internal/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type balanceRepository struct {
	repository.BaseRepository
}

// GetBalance implements repository.BalanceRepository.
func (b *balanceRepository) GetBalancesByWalletId(ctx context.Context, walletID uuid.UUID) ([]domain.GetBalanceResponse, error) {
	q := sqlc.New(b.DB())
	balances, err := q.GetBalancesByWalletId(ctx, pgtype.UUID{Bytes: walletID, Valid: true})
	if err != nil {
		return []domain.GetBalanceResponse{}, err
	}

	var result []domain.GetBalanceResponse
	for _, b := range balances {
		result = append(result, domain.GetBalanceResponse{
			TokenID:   b.TokenID.Bytes,
			Balance: func() float64 {
				val, _ := b.Balance.Float64Value()
				return val.Float64
			}(),
			UpdatedAt: b.UpdatedAt.Time,
			TokenName: b.TokenName,
			TokenSymbol: b.TokenSymbol,
			Decimals: int64(b.Decimals),
		})

	}

	return result, nil

}

func NewBalanceRepo(dbPool *pgxpool.Pool) repository.BalanceRepository {
	return &balanceRepository{
		BaseRepository: repository.NewBaseRepo(dbPool),
	}
}

var _ repository.BalanceRepository = (*balanceRepository)(nil)
