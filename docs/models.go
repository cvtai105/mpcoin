package docs

import "mpc/internal/domain"

type LoginResponse struct {
	Payload	  	struct {
		User         domain.LoginResponse `json:"user"`
		AccessToken  string               `json:"access_token"`
		RefreshToken string               `json:"refresh_token"`
	}
}

type RefreshResponse struct {
	Payload	  	struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
}

type SignupResponse struct {
	Payload	  	struct {
	User         domain.SignupResponse `json:"user"`
	Wallet       domain.CreateWalletResponse `json:"wallet"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	}
}

type CreateTxnResponse struct {
	Payload	  	struct {
	TransactionHash   string `json:"tx_hash"`
	Message    		string `json:"message"`
	}
}

type SubmitTnxResponse struct {
	Payload	  	struct {
		TransactionId   string `json:"tnx_hash"`
		Message    		string `json:"message"`
	}
}

type GetTxnResponse struct {
	Payload	  	struct {
		Transactions []domain.Transaction `json:"transactions"`
		Page 	   int                  `json:"page"`
		PerPage   int                  `json:"per_page"`
	}
}

type GetBalancesResponse struct {
	Payload	  	struct {
		GetBalanceResponse []domain.GetBalanceResponse `json:"balances"`
	}
}

type GetUserWalletResponse struct {
	Payload	  	struct {
		User domain.UserWithWallet `json:"user"`
	}
}
