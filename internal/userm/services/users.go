package services

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/glu/shopvui/idl/pb"
// 	"github.com/glu/shopvui/internal/userm/constants"
// 	"github.com/glu/shopvui/internal/userm/entities"
// 	"github.com/glu/shopvui/internal/userm/golibs/database"
// 	"github.com/glu/shopvui/internal/userm/models"
// 	"github.com/glu/shopvui/internal/userm/repositories"
// 	"github.com/glu/shopvui/utils"
// 	"github.com/google/uuid"
// 	"github.com/jackc/pgtype"
// 	"google.golang.org/protobuf/types/known/timestamppb"
// )

// type UserService struct {
// 	config utils.Config
// 	pb.UnimplementedUserServiceServer
// 	DB database.Ext

// 	UserRepo interface {
// 		GetUser(ctx context.Context, db database.Ext, email pgtype.Text) (*entities.User, error)
// 		CreateUser(ctx context.Context, db database.Ext, u *entities.User) (*entities.User, error)
// 		AddRoles(ctx context.Context, db database.Ext, roles *entities.Role) error
// 		GetRole(ctx context.Context, db database.Ext, roleName pgtype.Text) (*entities.Role, error)
// 		GetUserByID(ctx context.Context, db database.Ext, userID pgtype.Text) (*entities.User, error)
// 		UpdateRole(ctx context.Context, db database.Ext, e *entities.UserRole) (*entities.UserRole, error)
// 	}
// }

// func NewUserService(config utils.Config, db database.Ext) pb.UserServiceServer {
// 	return &UserService{
// 		config:   config,
// 		DB:       db,
// 		UserRepo: new(repositories.UserRepo),
// 	}
// }

// // gRPC service
// // func (u *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

// // 	user, err := u.UserRepo.GetUser(ctx, u.DB, database.Text(req.GetEmail()))
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	err = utils.CheckPassword(req.Password, user.Password.String)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	authenticateMaker, err := authenticate.NewPasetoMaker(u.config.authenticateSymmetricKey)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("cannot create authenticate maker: %w", err)
// // 	}

// // 	accessauthenticate, _, err := authenticateMaker.Createauthenticate(
// // 		user.ID.String,
// // 		u.config.AccessauthenticateDuration,
// // 	)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("cannot create access authenticate: %w", err)
// // 	}
// // 	return &pb.LoginResponse{
// // 		User: &pb.User{
// // 			Id:         user.ID.String,
// // 			Email:      user.Email.String,
// // 			FirstName:  user.FirstName.String,
// // 			LastName:   user.LastName.String,
// // 			CreateDate: timestamppb.New(user.InsertedAt.Time),
// // 		},
// // 		Accessauthenticate: accessauthenticate,
// // 	}, nil
// // }

// // func validateRegisterRequest(req *pb.RegisterRequest) error {
// // 	if req.Email == "" {
// // 		return fmt.Errorf("email is required")
// // 	}
// // 	if req.Password == "" {
// // 		return fmt.Errorf("password is required")
// // 	}
// // 	return nil
// // }

// // func toGrpcUser(req *pb.RegisterRequest) (*entities.User, error) {
// // 	hashPassword, err := utils.HashPassword(req.Password)
// // 	if err != nil {
// // 		return &entities.User{}, err
// // 	}
// // 	return &entities.User{
// // 		ID:        database.Text(uuid.NewString()),
// // 		Email:     database.Text(req.Email),
// // 		FirstName: database.Text(req.FirstName),
// // 		LastName:  database.Text(req.LastName),
// // 		Password:  database.Text(hashPassword),
// // 	}, nil
// // }
// // func (u *UserService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

// // 	if err := validateRegisterRequest(req); err != nil {
// // 		return nil, err
// // 	}
// // 	_, err := u.UserRepo.GetUser(ctx, u.DB, database.Text(req.GetEmail()))
// // 	if err == nil {
// // 		return nil, fmt.Errorf("user already exists")
// // 	}

// // 	ur, err := toGrpcUser(req)
// // 	if err != nil {
// // 		return nil, err
// // 	}
// // 	user, err := u.UserRepo.CreateUser(ctx, u.DB, ur)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("can't create user: %v", err)
// // 	}

// // 	authenticateMaker, err := authenticate.NewPasetoMaker(u.config.authenticateSymmetricKey)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("cannot create authenticate maker: %w", err)
// // 	}

// // 	accessauthenticate, _, err := authenticateMaker.Createauthenticate(
// // 		user.ID.String,
// // 		u.config.AccessauthenticateDuration,
// // 	)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("cannot create access authenticate: %w", err)
// // 	}

// // 	return &pb.RegisterResponse{
// // 		User: &pb.User{
// // 			Id:         user.ID.String,
// // 			Email:      user.Email.String,
// // 			FirstName:  user.FirstName.String,
// // 			LastName:   user.LastName.String,
// // 			CreateDate: timestamppb.New(user.InsertedAt.Time),
// // 		},
// // 		Accessauthenticate: accessauthenticate,
// // 	}, nil
// // }

// func (u *UserService) GetMe(ctx context.Context, req *pb.GetMeRequest) (*pb.GetMeResponse, error) {
// 	pay, err := u.authorizeUser(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user, err := u.UserRepo.GetUserByID(ctx, u.DB, database.Text(pay.UserID))
// 	if err != nil {
// 		return nil, fmt.Errorf("GetUser: %v", err)
// 	}

// 	return &pb.GetMeResponse{
// 		User: &pb.User{
// 			Id:         user.ID.String,
// 			Email:      user.Email.String,
// 			FirstName:  user.FirstName.String,
// 			LastName:   user.LastName.String,
// 			CreateDate: timestamppb.New(user.InsertedAt.Time),
// 		},
// 	}, nil
// }

// // RestFul api service
// func (s *Server) loginUser(ctx *gin.Context) {
// 	var req models.LoginUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, "cannot bin JSON")
// 		return
// 	}
// 	email := database.Text(req.Email)
// 	user, err := s.UserRepo.GetUser(ctx, s.db, email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, "User not found")
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, "User not found")
// 		return
// 	}
// 	if !user.Active.Bool {
// 		ctx.JSON(http.StatusInternalServerError, "user war detected")
// 	}

// 	err = utils.CheckPassword(req.Password, user.Password.String)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, err)
// 		return
// 	}

// 	accessauthenticate, _, err := s.authenticateMaker.Createauthenticate(
// 		user.ID.String,
// 		s.config.AccessauthenticateDuration,
// 	)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, "error")
// 		return
// 	}

// 	// refreshauthenticate, refreshPayload, err := server.authenticateMaker.Createauthenticate(
// 	// 	user.Username,
// 	// 	server.config.RefreshauthenticateDuration,
// 	// )
// 	// if err != nil {
// 	// 	ctx.JSON(http.StatusInternalServerError, "error")
// 	// 	return
// 	// }

// // 	rsp := models.LoginUserResponse{
// // 		Accessauthenticate: accessauthenticate,
// // 		//AccessauthenticateExpiresAt:  accessPayload.ExpiredAt,
// // 		//Refreshauthenticate:          refreshauthenticate,
// // 		//RefreshauthenticateExpiresAt: refreshPayload.ExpiredAt,
// // 		User: models.UserResponse{
// // 			Email:     user.Email.String,
// // 			FirstName: user.FirstName.String,
// // 			LastName:  user.LastName.String,
// // 			CreatedAt: user.InsertedAt.Time,
// // 			UpdatedAt: user.InsertedAt.Time,
// // 		},
// // 	}
// // 	ctx.JSON(http.StatusOK, rsp)
// // }

// // func toUser(req models.CreateUserRequest) (*entities.User, error) {
// // 	hashPassword, err := utils.HashPassword(req.Password)
// // 	if err != nil {
// // 		return &entities.User{}, err
// // 	}
// // 	return &entities.User{
// // 		ID:        database.Text(uuid.NewString()),
// // 		Email:     database.Text(req.Email),
// // 		FirstName: database.Text(req.FirstName),
// // 		LastName:  database.Text(req.LastName),
// // 		Password:  database.Text(hashPassword),
// // 	}, nil
// // }

// // func (s *Server) register(ctx *gin.Context) {
// // 	var req models.CreateUserRequest
// // 	if err := ctx.ShouldBindJSON(&req); err != nil {
// // 		ctx.JSON(http.StatusBadRequest, err)
// // 		return
// // 	}
// // 	email := database.Text(req.Email)
// // 	_, err := s.UserRepo.GetUser(ctx, s.db, email)
// // 	if err == nil {
// // 		ctx.JSON(http.StatusInternalServerError, "User already exists")
// // 		return
// // 	}
// // 	u, err := toUser(req)
// // 	if err != nil {
// // 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// // 		return
// // 	}

// // 	user, err := s.UserRepo.CreateUser(ctx, s.db, u)
// // 	if err != nil {
// // 		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("can't create user: %v", err))
// // 		return
// // 	}
// // 	rsp := models.UserResponse{
// // 		ID:        user.ID.String,
// // 		FirstName: user.FirstName.String,
// // 		LastName:  user.LastName.String,
// // 		Email:     user.Email.String,
// // 		CreatedAt: user.InsertedAt.Time,
// // 		UpdatedAt: user.UpdatedAt.Time,
// // 	}
// // 	ctx.JSON(http.StatusOK, rsp)
// // }

// // func toRole(req models.AddRoleRequest) *entities.Role {
// // 	return &entities.Role{
// // 		ID:   database.Text(uuid.NewString()),
// // 		Name: database.Text(req.Name),
// // 	}
// // }

// // func (s *Server) addRole(ctx *gin.Context) {
// // 	var req models.AddRoleRequest
// // 	if err := ctx.ShouldBindJSON(&req); err != nil {
// // 		ctx.JSON(http.StatusBadRequest, "can't bind JSON")
// // 		return
// // 	}
// // 	roles := toRole(req)
// // 	err := s.UserRepo.AddRoles(ctx, s.db, roles)
// // 	if err != nil {
// // 		fmt.Println(err.Error())
// // 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// // 		return
// // 	}
// // 	ctx.JSON(http.StatusOK, "success")
// // }

// // // need to update for scaling
// // func validateRole(role string) error {
// // 	if role == constants.RoleAdmin {
// // 		return fmt.Errorf("permission denied")
// // 	}
// // 	if role != constants.RoleUser && role != constants.RoleGuest {
// // 		return fmt.Errorf("invalid role")
// // 	}
// // 	return nil
// // }

// // func toUpdateRole(req models.UpdateRoleRequest, roleID string) *entities.UserRole {
// // 	return &entities.UserRole{
// // 		UserID: database.Text(req.UserID),
// // 		RoleID: database.Text(roleID),
// // 	}
// // }

// // func (s *Server) updateRole(ctx *gin.Context) {
// // 	var req models.UpdateRoleRequest
// // 	if err := ctx.ShouldBindJSON(&req); err != nil {
// // 		ctx.JSON(http.StatusBadRequest, "can't bind JSON")
// // 		return
// // 	}
// // 	if err := validateRole(req.RoleName); err != nil {
// // 		ctx.JSON(http.StatusBadRequest, err)
// // 		return
// // 	}
// // 	role, err := s.UserRepo.GetRole(ctx, s.db, database.Text(req.RoleName))
// // 	if err != nil {
// // 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// // 		return
// // 	}
// // 	user := toUpdateRole(req, role.ID.String)
// // 	_, err = s.UserRepo.UpdateRole(ctx, s.db, user)
// // 	if err != nil {
// // 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// // 		return
// // 	}
// // 	ctx.JSON(http.StatusOK, "success")
// // }
