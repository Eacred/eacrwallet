// Copyright (c) 2018 The Eacred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package udb

import (
	"github.com/Eacred/eacrd/chaincfg/chainhash"
	"github.com/Eacred/eacrd/gcs"
	"github.com/Eacred/eacrd/gcs/blockcf"
	"github.com/Eacred/eacrwallet/wallet/walletdb"
)

// CFilter returns the saved regular compact filter for a block.
func (s *Store) CFilter(dbtx walletdb.ReadTx, blockHash *chainhash.Hash) (*gcs.FilterV1, error) {
	ns := dbtx.ReadBucket(wtxmgrBucketKey)
	v, err := fetchRawCFilter(ns, blockHash[:])
	if err != nil {
		return nil, err
	}
	vc := make([]byte, len(v)) // Copy for FromfilterNData which stores passed slice
	copy(vc, v)
	return gcs.FromBytesV1(blockcf.P, vc)
}
