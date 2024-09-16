package full_node

import (
	"github.com/ghia-xch/ghia/pkg/node/protocol"
	"github.com/ghia-xch/ghia/pkg/node/protocol/message"
)

type NewUnfinishedBlock struct {
	UnfinishedRewardHash protocol.Hash
}

func (u *NewUnfinishedBlock) Type() message.Type { return protocol.NewUnfinishedBlock }

type RequestUnfinishedBlock struct {
	UnfinishedRewardHash protocol.Hash
}

func (u *RequestUnfinishedBlock) Type() message.Type { return protocol.RequestUnfinishedBlock }

type RespondUnfinishedBlock struct {
	UnfinishedBlock UnfinishedBlock
}

func (u *RespondUnfinishedBlock) Type() message.Type { return protocol.RespondUnfinishedBlock }

type NewUnfinishedBlock2 struct {
	UnfinishedRewardHash protocol.Hash
	FoliageHash          *protocol.Hash `cenc:"optional"`
}

func (u *NewUnfinishedBlock2) Type() message.Type { return protocol.NewUnfinishedBlock2 }

type RequestUnfinishedBlock2 struct {
	UnfinishedRewardHash protocol.Hash
	FoliageHash          *protocol.Hash `cenc:"optional"`
}

func (u *RequestUnfinishedBlock2) Type() message.Type { return protocol.RequestUnfinishedBlock2 }

///
///

type UnfinishedBlock struct{}
