package entity

// LoginRequest 登录请求
type LoginRequest struct {
	Email      string `json:"email" binding:"required,email" msg:"邮箱格式不正确"`
	Password   string `json:"password" binding:"required,min=6" msg:"密码至少6位"`
	RememberMe bool   `json:"rememberMe"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username        string `json:"username" binding:"required,min=3,max=32" msg:"用户名长度3-32位"`
	Email           string `json:"email" binding:"required,email" msg:"邮箱格式不正确"`
	Password        string `json:"password" binding:"required,min=6" msg:"密码至少6位"`
	ConfirmPassword string `json:"confirmPassword" binding:"required" msg:"确认密码不能为空"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	User         *UserInfo `json:"user"`
	Token        string    `json:"token"`        // JWT访问令牌
	RefreshToken string    `json:"refreshToken"` // 刷新令牌
	ExpiresIn    int64     `json:"expiresIn"`    // 过期时间（秒）
}

// UserInfo 用户信息
type UserInfo struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Role        string `json:"role"` // "admin" | "user"
	Nickname    string `json:"nickname,omitempty"`
	Website     string `json:"website,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// RefreshTokenRequest 刷新令牌请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// RefreshTokenResponse 刷新令牌响应
type RefreshTokenResponse struct {
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expiresIn"`
}

// UpdateProfileRequest 更新用户资料请求
type UpdateProfileRequest struct {
	Nickname    string `json:"nickname,omitempty"`
	Website     string `json:"website,omitempty"`
	Description string `json:"description,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
}

// ChangePasswordRequest 修改密码请求
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}
