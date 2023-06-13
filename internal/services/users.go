package services

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glu/shopvui/internal/entities"
	"github.com/glu/shopvui/internal/golibs/database"
	"github.com/glu/shopvui/internal/models"
	"github.com/glu/shopvui/util"
)

func (server *Server) loginUser(ctx *gin.Context) {
	var req models.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "error")
		return
	}
	email := database.Text(req.Email)
	user, err := server.UserRepo.GetUser(ctx, server.db, email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "error")
			return
		}
		ctx.JSON(http.StatusInternalServerError, "error")
		return
	}
	if user.Active.Bool {
		ctx.JSON(http.StatusInternalServerError, "user war detected")
	}

	err = util.CheckPassword(req.Password, user.Password.String)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "error")
		return
	}

	// accessToken, accessPayload, err := server.tokenMaker.CreateToken(
	// 	user.Username,
	// 	server.config.AccessTokenDuration,
	// )
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "error")
		return
	}

	// refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
	// 	user.Username,
	// 	server.config.RefreshTokenDuration,
	// )
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "error")
		return
	}

	rsp := models.LoginUserResponse{
		//AccessToken: accessToken,
		//AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		//RefreshToken:          refreshToken,
		//RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User: models.UserResponse{
			Email:     user.Email.String,
			FirstName: user.FirstName.String,
			LastName:  user.LastName.String,
			CreatedAt: user.Inserted_at.Time,
			UpdatedAt: user.Updated_at.Time,
		},
	}
	ctx.JSON(http.StatusOK, rsp)
}

func toUser(req models.CreateUserRequest) (*entities.User, error) {
	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return &entities.User{}, err
	}
	return &entities.User{
		Email:     database.Text(req.Email),
		FirstName: database.Text(req.FirstName),
		LastName:  database.Text(req.LastName),
		Password:  database.Text(hashPassword),
	}, nil
}

func (s *Server) register(ctx *gin.Context) {
	var req models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "error")
		return
	}
	u, err := toUser(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "can not create user")
		return
	}
	user, err := s.UserRepo.CreateUser(ctx, s.db, u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	rsp := models.UserResponse{
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Email:     user.Email.String,
		CreatedAt: user.Inserted_at.Time,
		UpdatedAt: user.Updated_at.Time,
	}
	ctx.JSON(http.StatusOK, rsp)
}
