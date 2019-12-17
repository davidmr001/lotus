package sharedutils

import (
	"bytes"
	sharedaddress "github.com/filecoin-project/go-fil-components/shared/address"
	sharedamount "github.com/filecoin-project/go-fil-components/shared/tokenamount"
	sharedtypes "github.com/filecoin-project/go-fil-components/shared/types"
	"github.com/filecoin-project/lotus/chain/address"
	"github.com/filecoin-project/lotus/chain/types"
)

func FromSharedAddress(in sharedaddress.Address) address.Address {
	addr, err := address.NewFromString(in.String())
	if err != nil {
		panic(err)
	}
	return addr
}

func ToSharedAddress(in address.Address) sharedaddress.Address {
	addr, err := sharedaddress.NewFromString(in.String())
	if err != nil {
		panic(err)
	}
	return addr
}

func FromSharedTokenAmount(in sharedamount.TokenAmount) types.BigInt {
	return types.BigInt{Int: in.Int}
}

func ToSharedTokenAmount(in types.BigInt) sharedamount.TokenAmount {
	return sharedamount.TokenAmount{Int: in.Int}
}

func ToSharedSignedVoucher(in *types.SignedVoucher) (*sharedtypes.SignedVoucher, error) {
	var encoded bytes.Buffer
	err := in.MarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	var out sharedtypes.SignedVoucher
	err = out.UnmarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func FromSharedSignedVoucher(in *sharedtypes.SignedVoucher) (*types.SignedVoucher, error) {
	var encoded bytes.Buffer
	err := in.MarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	var out types.SignedVoucher
	err = out.UnmarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func ToSharedSignature(in *types.Signature) (*sharedtypes.Signature, error) {
	var encoded bytes.Buffer
	err := in.MarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	var out sharedtypes.Signature
	err = out.UnmarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func FromSharedSignature(in *sharedtypes.Signature) (*types.Signature, error) {
	var encoded bytes.Buffer
	err := in.MarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	var out types.Signature
	err = out.UnmarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func ToSharedStorageAsk(in *types.SignedStorageAsk) (*sharedtypes.SignedStorageAsk, error) {
	var encoded bytes.Buffer
	err := in.MarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	var out sharedtypes.SignedStorageAsk
	err = out.UnmarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func FromSignedStorageAsk(in *sharedtypes.SignedStorageAsk) (*types.SignedStorageAsk, error) {
	var encoded bytes.Buffer
	err := in.MarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	var out types.SignedStorageAsk
	err = out.UnmarshalCBOR(&encoded)
	if err != nil {
		return nil, err
	}
	return &out, nil
}