package types

import (
	"github.com/pkg/errors"
	"github.com/nbit99/hx_sdk/encoding/transaction"
)

/*
{
	"previous": "0130b67ef5be56dbda6fbc364dccf453977a8812",
	"timestamp": "2022-03-22T06:57:55",
	"trxfee": 100,
	"miner": "1.6.6",
	"transaction_merkle_root": "8882f7eaf6cd29f6be79f0a33f88127dcedc5de9",
	"extensions": [],
	"next_secret_hash": "29fed79db3b14619e1764c1c858d8ddaec8bf470",
	"previous_secret": "0b72e9b2152a46d69373ecf99ffcb8af1ba1da6b",
	"miner_signature": "200a6cba7c66060826d0547040a69512473e09cc315213bde50ed4fef6584402b9105634b4a713da0ed03a6aced5c48cb14677e33e535d0a3d8a3b7ce1fded6b81",
	"transactions": [{
		"ref_block_num": 46718,
		"ref_block_prefix": 3679895285,
		"expiration": "2022-03-22T07:07:50",
		"operations": [
			[0, {
				"fee": {
					"amount": 100,
					"asset_id": "1.3.0"
				},
				"from": "1.2.0",
				"to": "1.2.0",
				"from_addr": "HXNUxdXWHhNakfHRx71EQprtpfZaFweYzNAt",
				"to_addr": "HXNX8E5W3HCjcA2CnPPGX5zgsHzf1LwihvDw",
				"amount": {
					"amount": 41387040,
					"asset_id": "1.3.0"
				},
				"memo": {
					"from": "HX1111111111111111111111111111111114T1Anm",
					"to": "HX1111111111111111111111111111111114T1Anm",
					"nonce": 0,
					"message": "000000006175746f"
				},
				"extensions": []
			}]
		],
		"extensions": [],
		"signatures": ["2079b1609edfffd252d7aef8f4242a9c10933dcd755ec4e6f51775dd6dbe5721a502bf5ed311d9de3c24c560fc3ce6b72d06521d6029bec80ee4bc677f61c64d2e"],
		"operation_results": [
			[0, {}]
		]
	}]
}
 */

type Transaction struct {
	RefBlockNum    uint16     `json:"ref_block_num"`
	RefBlockPrefix uint32     `json:"ref_block_prefix"`
	Expiration     Time       `json:"expiration"`
	Operations     Operations `json:"operations"`
	Signatures     []string   `json:"signatures"`
}

// MarshalTransaction implements transaction.Marshaller interface.
func (tx *Transaction) MarshalTransaction(encoder *transaction.Encoder) error {
	if len(tx.Operations) == 0 {
		return errors.New("no operation specified")
	}

	enc := transaction.NewRollingEncoder(encoder)

	enc.Encode(tx.RefBlockNum)
	enc.Encode(tx.RefBlockPrefix)
	enc.Encode(tx.Expiration)

	enc.EncodeUVarint(uint64(len(tx.Operations)))
	for _, op := range tx.Operations {
		enc.Encode(op)
	}

	// Extensions are not supported yet.
	enc.EncodeUVarint(0)
	return enc.Err()
}

// PushOperation can be used to add an operation into the transaction.
func (tx *Transaction) PushOperation(op Operation) {
	tx.Operations = append(tx.Operations, op)
}
