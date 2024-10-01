package services

import (
	"github.com/paghapour/golang-clean-web-api/api/dto"
	"github.com/paghapour/golang-clean-web-api/common"
	"github.com/paghapour/golang-clean-web-api/config"
	"github.com/paghapour/golang-clean-web-api/constants"
	"github.com/paghapour/golang-clean-web-api/data/db"
	"github.com/paghapour/golang-clean-web-api/data/models"
	"github.com/paghapour/golang-clean-web-api/pkg/logging"
	"github.com/paghapour/golang-clean-web-api/pkg/logging/service_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	OtpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		cfg:        cfg,
		database:   database,
		logger:     logger,
		OtpService: NewOtpService(cfg),
	}
}

// Register/Login By mobileNumber
// login by username

// Register By username
func (s *UserService) RegisterByUsername(req dto.RegisterUserByUsernameRequest) error{
	u := models.User{Username: req.UserName, FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Password: req.Password}

	exists, err := s.existsByEmail(req.Email)
	if err != nil{
		return err
	}
	if exists{
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = s.existsByUsername(req.UserName)
	if err != nil{
		return err
	}
	if exists{
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil{
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hp)
	roleId, err := s.getDefaultRole()
	if err != nil{
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}

	tx := s.database.Begin()
	err = tx.Create(&u).Error
	if err != nil{
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error
	if err != nil{
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil
}


func (s *UserService) SendOtp(req *dto.GetOtpRequest) error{
	otp := common.GenerateOtp()
	err := s.OtpService.SetOtp(req.MobileNumber, otp)
	if err != nil{
		return err
	}
	return nil
}

func (s *UserService) existsByEmail(email string) (bool, error){
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil{
			s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
			return false, err
		}
	return exists, nil
}

func (s *UserService) existsByUsername(username string) (bool, error){
	var exists bool

	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil{
			s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
			return false, err
		}
	return exists, nil
}

func (s *UserService) existsByMobileNumber(mobileNumber string) (bool, error){
	var exists bool

	if err := s.database.Model(&models.User{}).
	Select("count(*) > 0").
	Where("mobile_number = ?", mobileNumber).
	Find(&exists).
	Error; err != nil{
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error){
	if err = s.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil{
			return 0, err
		}
	return roleId, nil
}