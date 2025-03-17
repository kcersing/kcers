package jufenbei

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) Do(ctx context.Context, req *protocol.Request, res *protocol.Response) error {
	args := m.Called(ctx, req, res)
	return args.Error(0)
}
