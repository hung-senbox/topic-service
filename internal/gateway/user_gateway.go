package gateway

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/consul/api"
)

// User struct là dữ liệu trả về từ service user
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Thêm các field khác nếu cần
}

// UserGateway là interface để tương tác với service user
type UserGateway interface {
	GetAuthorInfo(ctx context.Context, userID string) (*User, error)
}

// userGatewayImpl là implementation của UserGateway
type userGatewayImpl struct {
	serviceName string
	consul      *api.Client
}

// NewUserGateway khởi tạo UserGateway
func NewUserGateway(serviceName string, consulClient *api.Client) UserGateway {
	return &userGatewayImpl{
		serviceName: serviceName,
		consul:      consulClient,
	}
}

// GetAuthorInfo lấy thông tin user từ service user
func (g *userGatewayImpl) GetAuthorInfo(ctx context.Context, userID string) (*User, error) {
	token, ok := ctx.Value("token").(string) // hoặc dùng constants.TokenKey
	if !ok || token == "" {
		return nil, fmt.Errorf("token không tồn tại trong context")
	}

	client, err := NewGatewayClient(g.serviceName, token, g.consul, nil)
	if err != nil {
		return nil, fmt.Errorf("khởi tạo GatewayClient thất bại: %w", err)
	}

	resp, err := client.Call("GET", "/v1/user/"+userID, nil)
	if err != nil {
		return nil, fmt.Errorf("gọi API user thất bại: %w", err)
	}

	var user User
	if err := json.Unmarshal(resp, &user); err != nil {
		return nil, fmt.Errorf("giải mã response thất bại: %w", err)
	}

	return &user, nil
}
