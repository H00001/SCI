package service

import (
	"context"
	"gunplan.top/SCI/protocol"
)

type PageImpl struct {
}

func (p *PageImpl) DoDown(ctx context.Context, inbound *protocol.PageInbound) (*protocol.PageOutbound, error) {
	return &protocol.PageOutbound{Content: inbound.Content}, nil
}
