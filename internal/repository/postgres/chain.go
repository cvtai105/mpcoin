package postgres

import (
	"context"
	"mpc/internal/domain"
	sqlc "mpc/internal/infrastructure/db/sqlc"
	"mpc/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type chainRepository struct {
	repository.BaseRepository
}

// GetChains implements repository.ChainRepository.
func (c *chainRepository) GetChains(ctx context.Context) ([]domain.Chain, error) {
	q := sqlc.New(c.DB())
	chains, err := q.GetChains(ctx)
	if err != nil {
		return nil, err
	}

	var result []domain.Chain
	for _, c := range chains {
		result = append(result, domain.Chain{
			ID:          c.ID.Bytes,
			Name:        c.Name,
			ChainID:     c.ChainID,
			NativeCurrency: c.NativeCurrency.String,
			NativeTokenID: c.NativeTokenID.Bytes,
			ExplorerURL: c.ExplorerUrl.String,
			RPCURL: 	c.RpcUrl,
			WSURL: 		c.WsUrl.String,
			LastScanBlock: c.LastScanBlockNumber.Int64,
			CreatedAt: 	c.CreatedAt.Time,
			UpdatedAt: 	c.UpdatedAt.Time,
		})
	}
	return result, nil
}

func NewChainRepo(dbPool *pgxpool.Pool) repository.ChainRepository {
	return &chainRepository{
		BaseRepository: repository.NewBaseRepo(dbPool),
	}
}

var _ repository.ChainRepository = (*chainRepository)(nil)
