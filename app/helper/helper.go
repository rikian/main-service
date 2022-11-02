package helper

import (
	"context"
	"crypto"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"main/golang/app/api/auth"
	"main/golang/app/entities"
	redisCli "main/golang/app/redis"

	pb "main/golang/app/grpc/protos"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var secretJwt = []byte("S4n94t_R4h4S14_BRO...")
var redisClient = redisCli.InitRedis()
var grpcClient = auth.NewApiAuth()

type id string

var userIdentity id

type utils struct {
	indentity *entities.Identity
	product   *pb.DataInsertProduct

	// update
	dataUpdateProduct *pb.DataUpdateProduct

	// errors
	errorSession *entities.ErrorSession
}

type UtilsImpl interface {
	JwtEncryptSession(id string, expired int) (string, error)
	JwtDecryptSession(tokenStr string) (jwt.MapClaims, error)
	IsValidEmail(email string) bool
	SHA256(text string) string
	DateFormat(layout string, d float64) string
	CheckDirIfNotExistAndCreateIt(filename string) error
	CheckDir(filename string) bool
	CookieParser(ctx *gin.Context) (*entities.Identity, error)
	CheckSession(ctx *gin.Context, r *redis.Client, data *entities.Identity) bool
	StrToBool(str string) bool
	GetContextId(ctx context.Context) (*entities.Identity, bool)
	StrToUint32(str string) (uint32, error)
	GUImageName(fileName, ext string) (string, error)
	ParseMultiPartData(f *gin.Context, id string) (*pb.DataInsertProduct, error)
	ParseMultiPartDataUpdate(f *gin.Context, id string) (*pb.DataUpdateProduct, error)
}

func InitHelper() UtilsImpl {
	return &utils{}
}

func (u *utils) GUImageName(fileName, ext string) (string, error) {
	imageName := uuid.New().String()

	switch strings.Split(ext, "/")[1] {
	case "jpeg":
		imageName += ".jpeg"
	case "jpg":
		imageName += ".jpg"
	case "png":
		imageName += ".png"
	default:
		return "", errors.New("invalid content type")
	}

	return imageName, nil
}

func (u *utils) JwtEncryptSession(id string, expired int) (string, error) {
	encrypt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Second * time.Duration(expired)).Unix(),
	})

	token, err := encrypt.SignedString(secretJwt)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *utils) JwtDecryptSession(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretJwt, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("not ok")
	}
}

func (u *utils) IsValidEmail(email string) bool {
	var regex = regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)

	return regex.MatchString(email)
}

func (u *utils) SHA256(text string) string {
	algorithm := crypto.SHA256.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

func (u *utils) DateFormat(layout string, d float64) string {
	intTime := int64(d)
	t := time.Unix(intTime, 0)
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return t.Format(layout)
}

func (u *utils) CheckDirIfNotExistAndCreateIt(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *utils) CheckDir(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func (u *utils) CookieParser(ctx *gin.Context) (*entities.Identity, error) {
	cookieUser, err := ctx.Request.Cookie("session")
	if err != nil {
		return nil, errors.New("invalid cookie session")
	}

	id, err := ctx.Request.Cookie("id")

	if err != nil {
		return nil, errors.New("invalid cookie id")
	}

	if cookieUser.Value == "" || id.Value == "" {
		return nil, errors.New("invalid cookie")
	}

	u.indentity = &entities.Identity{
		Id:       id.Value,
		JwtToken: cookieUser.Value,
	}

	return u.indentity, nil
}

func (u *utils) CheckSession(ctx *gin.Context, r *redis.Client, data *entities.Identity) bool {
	// get session user from redis if exist
	// sessionUserFromRedis, err := redisClient.GetSessionUser(r, data.Id)
	// if err == nil || sessionUserFromRedis != "" {
	// 	if !u.ParseJwtToken(sessionUserFromRedis, data) {
	// 		return false
	// 	}
	// 	ginctx := context.WithValue(ctx.Request.Context(), userIdentity, data)
	// 	ctx.Request = ctx.Request.WithContext(ginctx)
	// 	return true
	// }

	// get session user from db if session user in redis not exits
	sessionUserFromDB, err := grpcClient.SelectSessionUserById(data.Id)
	if err != nil {
		log.Print(err.Error())
		return false
	}

	if u.ParseJwtToken(sessionUserFromDB.UserSession, data).Status != 200 {
		return false
	}

	ginctx := context.WithValue(ctx.Request.Context(), userIdentity, data)
	ctx.Request = ctx.Request.WithContext(ginctx)
	return true
}

func (u *utils) ParseJwtToken(token string, data *entities.Identity) *entities.ErrorSession {
	tokenParts := strings.Split(token, ".")
	u.errorSession = &entities.ErrorSession{}

	if len(tokenParts) != 3 {
		u.errorSession.Status = 400
		u.errorSession.Message = "token not have 3 part"
		return u.errorSession
	}

	if tokenParts[2] != data.JwtToken {
		u.errorSession.Status = 400
		u.errorSession.Message = "token part 2 not valid"
		return u.errorSession
	}

	decryptToken, err := u.JwtDecryptSession(token)

	// if token expired
	if err != nil {
		log.Print(err.Error())
		u.errorSession.Status = 401
		u.errorSession.Message = "token is expired"
		return u.errorSession
	}

	if decryptToken["user_id"] == nil || decryptToken["user_id"] != data.Id {
		u.errorSession.Status = 400
		u.errorSession.Message = "token not have valid id"
		return u.errorSession
	}

	u.errorSession.Status = 200
	u.errorSession.Message = "ok"
	return u.errorSession
}

func (u *utils) StrToBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		log.Print(err.Error())
		return false
	}

	return b
}

func (u *utils) StrToUint32(str string) (uint32, error) {
	b, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		log.Print(err.Error())
		return 0, errors.New("invalid string. string must containe number. receive string : " + str)
	}

	return uint32(b), nil
}

func (u *utils) GetContextId(ctx context.Context) (*entities.Identity, bool) {
	value, ok := ctx.Value(userIdentity).(*entities.Identity)

	if !ok {
		return nil, ok
	}

	return value, ok
}

func (u *utils) ParseMultiPartData(f *gin.Context, id string) (*pb.DataInsertProduct, error) {
	pPrice, err := u.StrToUint32(f.PostForm("product_price"))

	if err != nil || pPrice == 0 {
		return nil, err
	}

	pSell, err := u.StrToUint32(f.PostForm("product_sell"))

	if err != nil || pSell == 0 {
		return nil, err
	}

	pStock, err := u.StrToUint32(f.PostForm("product_stock"))

	if err != nil || pStock == 0 {
		return nil, err
	}

	if f.PostForm("product_name") == "" || f.PostForm("product_stock") == "" {
		return nil, errors.New("invalid data product")
	}

	if f.PostForm("product_id") == "" {
		return nil, errors.New("invalid product id")
	}

	u.product = &pb.DataInsertProduct{
		UserId:       id,
		ProductId:    f.PostForm("product_id"),
		ProductName:  f.PostForm("product_name"),
		ProductPrice: pPrice,
		ProductSell:  pSell,
		ProductStock: pStock,
		ProductInfo:  f.PostForm("product_info"),
	}

	return u.product, nil
}

func (u *utils) ParseMultiPartDataUpdate(f *gin.Context, id string) (*pb.DataUpdateProduct, error) {
	pPrice, err := u.StrToUint32(f.PostForm("product_price"))

	if err != nil || pPrice == 0 {
		return nil, err
	}

	pSell, err := u.StrToUint32(f.PostForm("product_sell"))

	if err != nil || pSell == 0 {
		return nil, err
	}

	pStock, err := u.StrToUint32(f.PostForm("product_stock"))

	if err != nil || pStock == 0 {
		return nil, err
	}

	if f.PostForm("product_image") == "" || f.PostForm("product_id") == "" || f.PostForm("product_name") == "" || f.PostForm("product_stock") == "" || f.PostForm("product_info") == "" || f.PostForm("created_date") == "" {
		return nil, errors.New("invalid data product")
	}

	u.dataUpdateProduct = &pb.DataUpdateProduct{
		UserId:       id,
		ProductId:    f.PostForm("product_id"),
		ProductName:  f.PostForm("product_name"),
		ProductImage: f.PostForm("product_image"),
		ProductPrice: pPrice,
		ProductSell:  pSell,
		ProductStock: pStock,
		ProductInfo:  f.PostForm("product_info"),
		CreatedDate:  f.PostForm("created_date"),
	}

	return u.dataUpdateProduct, nil
}
