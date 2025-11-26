package service

import (
	"errors"
	"job-platform-go/internal/model/dto"
	"job-platform-go/internal/model/entity"
	"job-platform-go/internal/repository"
	"job-platform-go/pkg/utils"
	"strings"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repository.NewUserRepository(),
	}
}

// Login登录逻辑：查用户 -> 校验状态 -> 校验密码 -> 签发 Token -> 更新登录时间
func (s *AuthService) Login(req dto.LoginReq) (*dto.LoginResp, error) {
	//1.查找用户
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("邮箱未注册")
	}
	//2.校验状态
	if user.Status == entity.StatusDisabled {
		return nil, errors.New("账号已禁用")
	}
	if user.Status == entity.StatusPending {
		return nil, errors.New("账号审核中")
	}
	//3.校验密码
	if err := utils.CheckPasswordHash(req.Password, user.PasswordHash); err != nil {
		return nil, errors.New("密码错误")
	}
	//4.签发Token
	roleStr := "student"
	if user.Role == entity.RoleHR {
		roleStr = "hr"
	}
	token, err := utils.GenerateToken(user.ID, user.Email, roleStr)
	if err != nil {
		return nil, errors.New("生成Token失败")
	}
	//5.更新登录时间
	go s.userRepo.UpdateLastLogin(user.ID)
	//6.返回登录信息
	statusStr := "active"
	return &dto.LoginResp{
		Token: token,
		UserInfo: dto.UserInfo{
			ID:     user.ID,
			Email:  user.Email,
			Role:   roleStr,
			Status: statusStr,
		},
	}, nil

}

// Register 注册逻辑：校验验证码 -> 补全学生邮箱后缀 -> 检查邮箱唯一性 -> 密码加密 -> 调用 Repo 存库
func (s *AuthService) Register(req dto.RegisterReq) error {
	//1.校验验证码(占位逻辑)
	if req.VerificationCode != "123456" {
		return errors.New("验证码错误")
	}
	//2.补全学生邮箱后缀
	email := strings.TrimSpace(req.Email)
	if req.Role == "student" && !strings.Contains(email, "@") {
		email += "@mail2.sysu.edu.cn"
	}
	//3.检查邮箱是否存在
	existingUser, _ := s.userRepo.FindByEmail(email)
	if existingUser != nil {
		return errors.New("邮箱已被注册")

	}
	//4. 密码加密
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return errors.New("密码加密失败")
	}
	//5.调用 Repo 存库
	roleInt := entity.RoleStudent
	statusInt := entity.StatusActive // 学生默认激活
	if req.Role == "hr" {
		roleInt = entity.RoleHR
		statusInt = entity.StatusPending // HR 需要审核
	}
	user := &entity.User{
		Email:        email,
		PasswordHash: hash,
		Role:         roleInt,
		Status:       statusInt,
	}
	return s.userRepo.CreateUserWithRole(user, req.Role)
}
