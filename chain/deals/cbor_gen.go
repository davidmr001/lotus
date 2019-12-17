package deals

import (
	"fmt"
	"io"

	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/libp2p/go-libp2p-core/peer"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

var _ = xerrors.Errorf

func (t *AskRequest) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *AskRequest) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *AskResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.Ask (types.SignedStorageAsk) (struct)
	if err := t.Ask.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *AskResponse) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Ask (types.SignedStorageAsk) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Ask = new(types.SignedStorageAsk)
			if err := t.Ask.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	return nil
}

func (t *Proposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.DealProposal (actors.StorageDealProposal) (struct)
	if err := t.DealProposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Piece (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Piece); err != nil {
		return xerrors.Errorf("failed to write cid field t.Piece: %w", err)
	}

	return nil
}

func (t *Proposal) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.DealProposal (actors.StorageDealProposal) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.DealProposal = new(actors.StorageDealProposal)
			if err := t.DealProposal.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	// t.Piece (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Piece: %w", err)
		}

		t.Piece = c

	}
	return nil
}

func (t *Response) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{132}); err != nil {
		return err
	}

	// t.State (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.State))); err != nil {
		return err
	}

	// t.Message (string) (string)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Message)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Message)); err != nil {
		return err
	}

	// t.Proposal (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Proposal); err != nil {
		return xerrors.Errorf("failed to write cid field t.Proposal: %w", err)
	}

	// t.PublishMessage (cid.Cid) (struct)

	if t.PublishMessage == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(w, *t.PublishMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.PublishMessage: %w", err)
		}
	}

	return nil
}

func (t *Response) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.State (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.State = uint64(extra)
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	// t.Proposal (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Proposal: %w", err)
		}

		t.Proposal = c

	}
	// t.PublishMessage (cid.Cid) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.PublishMessage: %w", err)
			}

			t.PublishMessage = &c
		}

	}
	return nil
}

func (t *SignedResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Response (deals.Response) (struct)
	if err := t.Response.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Signature (types.Signature) (struct)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SignedResponse) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Response (deals.Response) (struct)

	{

		if err := t.Response.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Signature (types.Signature) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Signature = new(types.Signature)
			if err := t.Signature.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	return nil
}

func (t *ClientDealProposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{136}); err != nil {
		return err
	}

	// t.Data (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Data); err != nil {
		return xerrors.Errorf("failed to write cid field t.Data: %w", err)
	}

	// t.PricePerEpoch (types.BigInt) (struct)
	if err := t.PricePerEpoch.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ProposalExpiration (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.ProposalExpiration))); err != nil {
		return err
	}

	// t.Duration (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Duration))); err != nil {
		return err
	}

	// t.ProviderAddress (address.Address) (struct)
	if err := t.ProviderAddress.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Client (address.Address) (struct)
	if err := t.Client.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinerWorker (address.Address) (struct)
	if err := t.MinerWorker.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinerID (peer.ID) (string)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.MinerID)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.MinerID)); err != nil {
		return err
	}
	return nil
}

func (t *ClientDealProposal) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 8 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Data (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Data: %w", err)
		}

		t.Data = c

	}
	// t.PricePerEpoch (types.BigInt) (struct)

	{

		if err := t.PricePerEpoch.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.ProposalExpiration (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.ProposalExpiration = uint64(extra)
	// t.Duration (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Duration = uint64(extra)
	// t.ProviderAddress (address.Address) (struct)

	{

		if err := t.ProviderAddress.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.Client (address.Address) (struct)

	{

		if err := t.Client.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.MinerWorker (address.Address) (struct)

	{

		if err := t.MinerWorker.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.MinerID (peer.ID) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.MinerID = peer.ID(sval)
	}
	return nil
}

func (t *ClientDeal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.ClientDeal (storagemarket.ClientDeal) (struct)
	if err := t.ClientDeal.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *ClientDeal) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.ClientDeal (storagemarket.ClientDeal) (struct)

	{

		if err := t.ClientDeal.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *MinerDeal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{135}); err != nil {
		return err
	}

	// t.Client (peer.ID) (string)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Client)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Client)); err != nil {
		return err
	}

	// t.Proposal (actors.StorageDealProposal) (struct)
	if err := t.Proposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ProposalCid (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.ProposalCid); err != nil {
		return xerrors.Errorf("failed to write cid field t.ProposalCid: %w", err)
	}

	// t.State (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.State))); err != nil {
		return err
	}

	// t.Ref (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Ref); err != nil {
		return xerrors.Errorf("failed to write cid field t.Ref: %w", err)
	}

	// t.DealID (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.DealID))); err != nil {
		return err
	}

	// t.SectorID (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SectorID))); err != nil {
		return err
	}
	return nil
}

func (t *MinerDeal) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 7 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Client (peer.ID) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Client = peer.ID(sval)
	}
	// t.Proposal (actors.StorageDealProposal) (struct)

	{

		if err := t.Proposal.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.ProposalCid (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ProposalCid: %w", err)
		}

		t.ProposalCid = c

	}
	// t.State (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.State = uint64(extra)
	// t.Ref (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Ref: %w", err)
		}

		t.Ref = c

	}
	// t.DealID (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.DealID = uint64(extra)
	// t.SectorID (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.SectorID = uint64(extra)
	return nil
}

func (t *StorageDataTransferVoucher) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.Proposal (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Proposal); err != nil {
		return xerrors.Errorf("failed to write cid field t.Proposal: %w", err)
	}

	// t.DealID (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.DealID))); err != nil {
		return err
	}
	return nil
}

func (t *StorageDataTransferVoucher) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Proposal (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Proposal: %w", err)
		}

		t.Proposal = c

	}
	// t.DealID (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.DealID = uint64(extra)
	return nil
}
