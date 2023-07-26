package vss

import "go.dedis.ch/kyber/v3/share"

type DealerOption func(*Dealer) error

func WithFPoly(f *share.PriPoly) DealerOption {
	return func(d *Dealer) error {
		d.f = f
		return nil
	}
}

func WithGPoly(g *share.PriPoly) DealerOption {
	return func(d *Dealer) error {
		d.g = g
		return nil
	}
}

func defaultPriPolys() DealerOption {
	return func(d *Dealer) error {
		d.f = share.NewPriPoly(d.suite, d.t, d.secret, d.suite.RandomStream())
		d.g = share.NewPriPoly(d.suite, d.t, nil, d.suite.RandomStream())
		return nil
	}
}

func DefaultOptions() []DealerOption {
	return []DealerOption{
		defaultPriPolys(),
	}
}
