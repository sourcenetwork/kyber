package dkg

import vss "go.dedis.ch/kyber/v3/share/vss/rabin"

type DKGOption func(*DistKeyGenerator) error

func WithDealer(dealer *vss.Dealer) DKGOption {
	return func(dkg *DistKeyGenerator) error {
		dkg.dealer = dealer
		return nil
	}
}

func defaultDealerOpt() DKGOption {
	return func(dkg *DistKeyGenerator) error {
		var err error
		ownSec := dkg.suite.Scalar().Pick(dkg.suite.RandomStream())
		dkg.dealer, err = vss.NewDealer(dkg.suite, dkg.long, ownSec, dkg.participants, dkg.t)
		return err
	}
}

func DefaultOptions() []DKGOption {
	return []DKGOption{
		defaultDealerOpt(),
	}
}
