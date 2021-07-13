// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package transaction

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var errNotCached = errors.New("sender not stored")

// QuorumSigner is a struct that stores data related to signing
// Stored to avoid additional request
type QuorumSigner struct {
	addr      common.Address
	blockhash common.Hash
}

func setSenderFromServer(tx *types.Transaction, addr common.Address, block common.Hash) {
	types.Sender(&QuorumSigner{addr, block}, tx)
}

func (q *QuorumSigner) Equal(other types.Signer) bool {
	os, ok := other.(*QuorumSigner)
	return ok && os.blockhash == q.blockhash
}

func (q *QuorumSigner) Sender(tx *types.Transaction) (common.Address, error) {
	if q.blockhash == (common.Hash{}) {
		return common.Address{}, errNotCached
	}
	return q.addr, nil
}

func (s *QuorumSigner) Hash(tx *types.Transaction) common.Hash {
	panic("can't sign with senderFromServer")
}

func (s *QuorumSigner) SignatureValues(tx *types.Transaction, sig []byte) (R, S, V *big.Int, err error) {
	panic("can't sign with senderFromServer")
}
