package rest_txs

import (
	"net/http"

	"github.com/gorilla/mux"
	wire "github.com/tendermint/go-wire"
	rpcclient "github.com/tendermint/tendermint/rpc/client"

	sdk "github.com/cosmos/cosmos-sdk"
)

// ServiceTxs exposes a REST API service for sendings txs.
// It wraps a Tendermint RPC client.
type ServiceTxs struct {
	node rpcclient.Client
}

func NewServiceTxs(c rpcclient.Client) *ServiceTxs {
	return &ServiceTxs{
		node: c,
	}
}

func (s *ServiceTxs) PostTx(w http.ResponseWriter, r *http.Request) {
	tx := new(sdk.Tx)
	if err := sdk.ParseRequestAndValidateJSON(r, tx); err != nil {
		sdk.WriteError(w, err)
		return
	}
	
	packet := wire.BinaryBytes(*tx)
	// commit, err := s.node.BroadcastTxCommit(packet)
	commit, err := s.node.BroadcastTxSync(packet)
	if err != nil {
		sdk.WriteError(w, err)
		return
	}
	
	sdk.WriteSuccess(w, commit)
}

// mux.Router registrars

// RegisterPostTx is a mux.Router handler that exposes POST
// method access to post a transaction to the blockchain.
func (s *ServiceTxs) RegisterPostTx(r *mux.Router) error {
	r.HandleFunc("/tx", s.PostTx).Methods("POST")
	return nil
}

// End of mux.Router registrars

