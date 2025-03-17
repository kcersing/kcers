package douyin

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) Do(ctx context.Context, req *protocol.Request, res *protocol.Response) error {
	args := m.Called(ctx, req, res)
	return args.Error(0)
}

func TestGetAccessToken(t *testing.T) {

	//// 创建模拟客户端
	//mockClient := new(MockClient)
	//
	//// 模拟客户端的行为
	//mockClient.On("Do", context.Background(), mock.Anything, mock.Anything).Return(nil)

	params := map[string]string{
		"account_id": "7218069838395082791",
		"page":       "1",
		"size":       "100",
	}
	GetCrmQuery(params)

	// 调用目标函数
	//_, err := GetAccessToken()

	//// 断言错误为 nil
	//assert.NoError(t, err)
	//
	//// 验证模拟客户端的调用
	//mockClient.AssertExpectations(t)
}

//func TestGetAccessTokenTimeout(t *testing.T) {
//	// 创建模拟客户端
//	mockClient := new(MockClient)
//
//	// 模拟客户端超时错误
//	mockClient.On("Do", context.Background(), mock.Anything, mock.Anything).Return(context.DeadlineExceeded)
//
//	// 调用目标函数
//	_, err := GetAccessToken()
//
//	// 断言错误为超时错误
//	assert.Equal(t, context.DeadlineExceeded, err)
//
//	// 验证模拟客户端的调用
//	mockClient.AssertExpectations(t)
//}
