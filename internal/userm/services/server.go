package services

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/glu/shopvui/internal/userm/entities"
	"github.com/glu/shopvui/internal/userm/golibs/database"
	"github.com/glu/shopvui/internal/userm/repositories"

	"github.com/glu/shopvui/token"
	"github.com/glu/shopvui/util"
	"github.com/jackc/pgtype"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	//store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	db         database.Ext

	UserRepo interface {
		GetUser(ctx context.Context, db database.Ext, email pgtype.Text) (*entities.User, error)
		CreateUser(ctx context.Context, db database.Ext, u *entities.User) (*entities.User, error)
		AddRoles(ctx context.Context, db database.Ext, roles *entities.Role) error
		GetRole(ctx context.Context, db database.Ext, roleName pgtype.Text) (*entities.Role, error)
		UpdateRole(ctx context.Context, db database.Ext, e *entities.UserRole) (*entities.UserRole, error)
	}
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, db database.Ext) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		db:         db,
		tokenMaker: tokenMaker,
		UserRepo:   new(repositories.UserRepo),
	}

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("currency", validCurrency)
	// }

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users/register", server.register)
	router.POST("/users/login", server.loginUser)
	router.POST("/users/role", server.addRole)
	router.POST("users/role/update", server.updateRole)
	// router.POST("/tokens/renew_access", server.renewAccessToken)

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccounts)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
