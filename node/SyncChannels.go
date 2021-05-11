package node

import (
	"context"
	"encoding/json"

	"github.com/driemworks/mercury-blockchain/state"
	"github.com/libp2p/go-libp2p-core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// PendingTxBufSize is the number of incoming pending transactions to buffer for each epoch.
const PendingTxBufSize = 128

// Channel represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Channel.Publish, and received
// messages are pushed to the Messages channel.
type Channel struct {
	// A channel of signed transactions to send new pending transactions to peers
	PendingTransactions chan *state.SignedTx

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	roomName string
	self     peer.ID
	nick     string
}

// JoinChannel tries to subscribe to the PubSub topic for the room name, returning
// a ChatRoom on success.
func JoinChannel(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, nickname string, roomName string) (*Channel, error) {
	// join the pubsub topic
	topic, err := ps.Join(topicName(roomName))
	if err != nil {
		return nil, err
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}
	cr := &Channel{
		ctx:                 ctx,
		ps:                  ps,
		topic:               topic,
		sub:                 sub,
		self:                selfID,
		nick:                nickname,
		roomName:            roomName,
		PendingTransactions: make(chan *state.SignedTx, PendingTxBufSize),
	}
	// start reading messages from the subscription in a loop
	go cr.readLoop()
	return cr, nil
}

// Publish sends a message to the pubsub topic.
func (cr *Channel) Publish(tx *state.SignedTx) error {
	msgBytes, err := json.Marshal(tx)
	if err != nil {
		return err
	}
	return cr.topic.Publish(cr.ctx, msgBytes)
}

func (cr *Channel) ListPeers() []peer.ID {
	return cr.ps.ListPeers(topicName(cr.roomName))
}

// // readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (cr *Channel) readLoop() {
	for {
		msg, err := cr.sub.Next(cr.ctx)
		if err != nil {
			close(cr.PendingTransactions)
			return
		}
		// only forward messages delivered by others
		if msg.ReceivedFrom == cr.self {
			continue
		}
		cm := new(state.SignedTx)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}
		// send valid messages onto the Messages channel
		cr.PendingTransactions <- cm
	}
}

func topicName(roomName string) string {
	return "chat-room:" + roomName
}
