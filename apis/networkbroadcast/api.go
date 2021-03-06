package networkbroadcast

import (
	"github.com/nbit99/hx_sdk/caller"
	"github.com/nbit99/hx_sdk/hx"
)

type API struct {
	caller caller.Caller
	id     caller.APIID
}

func NewAPI(id caller.APIID, caller caller.Caller) *API {
	return &API{id: id, caller: caller}
}

func (api *API) call(method string, args []interface{}, reply interface{}) error {
	return api.caller.Call(api.id, method, args, reply)
}

// BroadcastTransaction broadcast a transaction to the network.
func (api *API) BroadcastTransaction(tx *hx.Transaction) error {
	return api.call("broadcast_transaction", []interface{}{tx}, nil)
}

func (api *API) BroadcastTransactionSynchronous(tx *hx.Transaction) (*BroadcastResponse, error) {
	response := BroadcastResponse{}
	err := api.call("broadcast_transaction_synchronous", []interface{}{tx}, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}
