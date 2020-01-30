// Copyright (c) 2017 The Eacred developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wallet

import (
	"context"

	"github.com/Eacred/eacrd/hdkeychain"
	"github.com/Eacred/eacrwallet/errors"
	"github.com/Eacred/eacrwallet/wallet/udb"
	"github.com/Eacred/eacrwallet/wallet/walletdb"
)

// UpgradeToSLIP0044CoinType upgrades the wallet from the legacy BIP0044 coin
// type to one of the coin types assigned to Eacred in SLIP0044.  This should be
// called after a new wallet is created with a random (not imported) seed.
//
// This function does not register addresses from the new account 0 with the
// wallet's network backend.  This is intentional as it allows offline
// activities, such as wallet creation, to perform this upgrade.
func (w *Wallet) UpgradeToSLIP0044CoinType(ctx context.Context) error {
	const op errors.Op = "wallet.UpgradeToSLIP0044CoinType"

	var acctXpub, extBranchXpub, intBranchXpub *hdkeychain.ExtendedKey

	err := walletdb.Update(ctx, w.db, func(dbtx walletdb.ReadWriteTx) error {
		err := w.Manager.UpgradeToSLIP0044CoinType(dbtx)
		if err != nil {
			return err
		}

		acctXpub, err = w.Manager.AccountExtendedPubKey(dbtx, 0)
		if err != nil {
			return err
		}

		extBranchXpub, err = w.Manager.AccountBranchExtendedPubKey(dbtx, 0,
			udb.ExternalBranch)
		if err != nil {
			return err
		}
		intBranchXpub, err = w.Manager.AccountBranchExtendedPubKey(dbtx, 0,
			udb.InternalBranch)
		return err
	})
	if err != nil {
		return errors.E(op, err)
	}

	w.addressBuffersMu.Lock()
	w.addressBuffers[0] = &bip0044AccountData{
		xpub:        acctXpub,
		albExternal: addressBuffer{branchXpub: extBranchXpub, lastUsed: ^uint32(0)},
		albInternal: addressBuffer{branchXpub: intBranchXpub, lastUsed: ^uint32(0)},
	}
	w.addressBuffersMu.Unlock()

	return nil
}
