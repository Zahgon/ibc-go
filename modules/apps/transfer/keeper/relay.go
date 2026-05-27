package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/ibc-go/v11/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v11/modules/core/04-channel/types"
)

// SendTransfer handles transfer sending logic. There are 2 possible cases:
//
// 1. Sender chain is acting as the source zone. The coins are transferred
// to an escrow address (i.e locked) on the sender chain and then transferred
// to the receiving chain through IBC TAO logic. It is expected that the
// receiving chain will mint vouchers to the receiving address.
//
// 2. Sender chain is acting as the sink zone. The coins (vouchers) are burned
// on the sender chain and then transferred to the receiving chain though IBC
// TAO logic. It is expected that the receiving chain, which had previously
// sent the original denomination, will unescrow the fungible token and send
// it to the receiving address.
//
// Another way of thinking of source and sink zones is through the token's
// timeline. Each send to any chain other than the one it was previously
// received from is a movement forwards in the token's timeline. This causes
// trace to be added to the token's history and the destination port and
// destination channel to be prefixed to the denomination. In these instances
// the sender chain is acting as the source zone. When the token is sent back
// to the chain it previously received from, the prefix is removed. This is
// a backwards movement in the token's timeline and the sender chain
// is acting as the sink zone.
//
// Example:
// These steps of transfer occur: A -> B -> C -> A -> C -> B -> A
//
// 1. A -> B : sender chain is source zone. Denom upon receiving: 'B/denom'
// 2. B -> C : sender chain is source zone. Denom upon receiving: 'C/B/denom'
// 3. C -> A : sender chain is source zone. Denom upon receiving: 'A/C/B/denom'
// 4. A -> C : sender chain is sink zone. Denom upon receiving: 'C/B/denom'
// 5. C -> B : sender chain is sink zone. Denom upon receiving: 'B/denom'
// 6. B -> A : sender chain is sink zone. Denom upon receiving: 'denom'
func (k *Keeper) SendTransfer(
	ctx sdk.Context,
	sourcePort,
	sourceChannel string,
	token types.Token,
	sender sdk.AccAddress,
) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: SendTransfer simply sends the denomination as it exists on its own
// chain inside the packet data. The receiving chain will perform denom
// prefixing as necessary.

// if the denom is prefixed by the port and channel on which we are sending
// the token, then we must be returning the token back to the chain they originated from

// transfer the coins to the module account and burn them

// NOTE: should not happen as the module account was
// retrieved on the step above and it has enough balance
// to burn.

// obtain the escrow address for the source channel end

// OnRecvPacket processes a cross chain fungible token transfer.
//
// If the sender chain is the source of minted tokens then vouchers will be minted
// and sent to the receiving address. Otherwise if the sender chain is sending
// back tokens this chain originally transferred to it, the tokens are
// unescrowed and sent to the receiving address.
func (k *Keeper) OnRecvPacket(
	ctx sdk.Context,
	data types.InternalTransferRepresentation,
	sourcePort string,
	sourceChannel string,
	destPort string,
	destChannel string,
) error {
	_ = "STUB: not implemented"
	// validate packet data upon receiving
	return nil
}

// parse the transfer amount

// This is the prefix that would have been prefixed to the denomination
// on sender chain IF and only if the token originally came from the
// receiving chain.
//
// NOTE: We use SourcePort and SourceChannel here, because the counterparty
// chain would have prefixed with DestPort and DestChannel when originally
// receiving this token.

// sender chain is not the source, unescrow tokens

// remove prefix added by sender chain

// sender chain is the source, mint vouchers

// since SendPacket did not prefix the denomination, we must add the destination port and channel to the trace

// mint new tokens if the source of the transfer is the same chain

// send to receiver

// The ibc_module.go module will return the proper ack.

// OnAcknowledgementPacket responds to the success or failure of a packet acknowledgment
// written on the receiving chain.
//
// If the acknowledgement was a success then nothing occurs. Otherwise,
// if the acknowledgement failed, then the sender is refunded their tokens.
func (k *Keeper) OnAcknowledgementPacket(
	ctx sdk.Context,
	sourcePort string,
	sourceChannel string,
	data types.InternalTransferRepresentation,
	ack channeltypes.Acknowledgement,
) error {
	_ = "STUB: not implemented"
	return nil
}

// the acknowledgement succeeded on the receiving chain so nothing
// needs to be executed and no error needs to be returned

// OnTimeoutPacket processes a transfer packet timeout by refunding the tokens to the sender
func (k *Keeper) OnTimeoutPacket(
	ctx sdk.Context,
	sourcePort string,
	sourceChannel string,
	data types.InternalTransferRepresentation,
) error {
	_ = "STUB: not implemented"
	return nil
}

// refundPacketTokens will unescrow and send back the token back to sender
// if the sending chain was the source chain. Otherwise, the sent token
// were burnt in the original send so new tokens are minted and sent to
// the sending address.
func (k *Keeper) refundPacketTokens(
	ctx sdk.Context,
	sourcePort string,
	sourceChannel string,
	data types.InternalTransferRepresentation,
) error {
	_ = "STUB: not implemented"
	// NOTE: packet data type already checked in handler.go
	return nil
}

// escrow address for unescrowing tokens back to sender

// if the token we must refund is prefixed by the source port and channel
// then the tokens were burnt when the packet was sent and we must mint new tokens

// mint vouchers back to sender

// EscrowCoin will send the given coin from the provided sender to the escrow address. It will also
// update the total escrowed amount by adding the escrowed coin's amount to the current total escrow.
func (k *Keeper) EscrowCoin(ctx sdk.Context, sender, escrowAddress sdk.AccAddress, coin sdk.Coin) error {
	_ = "STUB: not implemented"
	return nil
}

// failure is expected for insufficient balances

// track the total amount in escrow keyed by denomination to allow for efficient iteration

// UnescrowCoin will send the given coin from the escrow address to the provided receiver. It will also
// update the total escrow by deducting the unescrowed coin's amount from the current total escrow.
func (k *Keeper) UnescrowCoin(ctx sdk.Context, escrowAddress, receiver sdk.AccAddress, coin sdk.Coin) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: this error is only expected to occur given an unexpected bug or a malicious
// counterparty module. The bug may occur in bank or any part of the code that allows
// the escrow address to be drained. A malicious counterparty module could drain the
// escrow address by allowing more tokens to be sent back then were escrowed.

// track the total amount in escrow keyed by denomination to allow for efficient iteration

// tokenFromCoin constructs an IBC token given an SDK coin.
func (k *Keeper) TokenFromCoin(ctx sdk.Context, coin sdk.Coin) (types.Token, error) {
	_ = "STUB: not implemented"
	// if the coin does not have an IBC denom, return as is
	return *new(types.Token), nil
}

// NOTE: denomination and hex hash correctness checked during msg.ValidateBasic

// GetDenomFromIBCDenom returns the `Denom` given the IBC Denom (ibc/{hex hash}) of the denomination.
// The ibcDenom is the hex hash of the denomination prefixed by "ibc/", often referred to as the IBC denom.
func (k *Keeper) GetDenomFromIBCDenom(ctx sdk.Context, ibcDenom string) (types.Denom, error) {
	_ = "STUB: not implemented"
	return *new(types.Denom), nil
}

// Deprecated: usage of this function should be replaced by `Keeper.GetDenomFromIBCDenom`
// DenomPathFromHash returns the full denomination path prefix from an ibc denom with a hash
// component.
func (k *Keeper) DenomPathFromHash(ctx sdk.Context, ibcDenom string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// createPacketDataBytesFromVersion creates the packet data bytes to be sent based on the application version.
func createPacketDataBytesFromVersion(appVersion, sender, receiver, memo string, token types.Token) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
