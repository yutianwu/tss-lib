package tss

import (
	"errors"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"

	"github.com/binance-chain/tss-lib/protob"
)

const (
	ProtoNamePrefix = "binance.tss-lib.ecdsa."
)

func ParseMessageFromProtoB(wire *protob.Message, from *PartyID, to []*PartyID) (ParsedMessage, error) {
	var any ptypes.DynamicAny
	meta := MessageMetadata{
		From: from,
		To:   to,
	}
	err := ptypes.UnmarshalAny(wire.Message, &any)
	if err != nil {
		return nil, err
	}
	if content, ok := any.Message.(MessageContent); ok {
		return NewMessage(meta, content, wire), nil
	}
	return nil, errors.New("ParseMessage: the message contained unknown content")
}

// Used externally to update a LocalParty with a valid ParsedMessage
func ParseMessage(wireBytes []byte, from *PartyID, to []*PartyID) (ParsedMessage, error) {
	wire := new(protob.Message)
	if err := proto.Unmarshal(wireBytes, wire); err != nil {
		return nil, err
	}
	return ParseMessageFromProtoB(wire, from, to)
}