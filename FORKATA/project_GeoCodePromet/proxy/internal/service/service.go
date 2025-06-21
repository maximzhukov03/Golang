package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"task25/proxy/internal/repository"
	"time"

	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{
	service *database.UserRepositoryPostgres
}

func NewUserSevice(db *database.UserRepositoryPostgres) *UserService{
	return &UserService{
		service: db,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user_name, password string) error{
	if user_name == "" || password == ""{
		return fmt.Errorf("user_name or password are empty")
	}
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(1000000000)
	idstr := strconv.Itoa(id)
	err := s.service.Create(ctx, idstr, user_name, password)
	if err != nil{
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, user_name string) error{
	if user_name == ""{
		return fmt.Errorf("user_name are empty")
	}
	err := s.service.Delete(ctx, user_name)
	if err != nil{
		return err
	}
	return nil
}

func (s *UserService) FindForUserName(ctx context.Context, user_name string) (*database.User, error){
	user := &database.User{}
	if user_name == ""{
		return nil, fmt.Errorf("user_name are empty")
	}
	user, err := s.service.FindUser(ctx, user_name)
	if err != nil{
		return nil, err
	}
	return user, nil
}

type GeoService struct {
	api       *suggest.Api
	cache     *redis.Client
	apiKey    string
	secretKey string
}

type GeoProvider interface {
	AddressSearch(input string) ([]*Address, error)
	GeoCode(lat, lng string) ([]*Address, error)
}

func NewGeoService(apiKey, secretKey string) *GeoService {
	var err error
	endpointUrl, err := url.Parse("https://suggestions.dadata.ru/suggestions/api/4_1/rs/")
	if err != nil {
		return nil
	}

	creds := client.Credentials{
		ApiKeyValue:    apiKey,
		SecretKeyValue: secretKey,
	}

	api := suggest.Api{
		Client: client.NewClient(endpointUrl, client.WithCredentialProvider(&creds)),
	}

	return &GeoService{
		api:       &api,
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
	House  string `json:"house"`
	Lat    string `json:"lat"`
	Lon    string `json:"lon"`
}

func (g *GeoService) AddressSearch(input string) ([]*Address, error) {
	var res []*Address
	rawRes, err := g.api.Address(context.Background(), &suggest.RequestParams{Query: input})
	if err != nil {
		log.Println("Ошибка при попытке получить")
		return nil, err
	}

	for _, r := range rawRes {
		if r.Data.City == "" || r.Data.Street == "" {
			continue
		}
		res = append(res, &Address{City: r.Data.City, Street: r.Data.Street, House: r.Data.House, Lat: r.Data.GeoLat, Lon: r.Data.GeoLon})
	}

	return res, nil
}

func (g *GeoService) GeoCode(lat, lng string) ([]*Address, error) {
	httpClient := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"lat": %s, "lon": %s}`, lat, lng))
	req, err := http.NewRequest("POST", "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", g.apiKey))
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	var geoCode GeoCode

	err = json.NewDecoder(resp.Body).Decode(&geoCode)
	if err != nil {
		return nil, err
	}
	var res []*Address
	for _, r := range geoCode.Suggestions {
		var address Address
		address.City = string(r.Data.City)
		address.Street = string(r.Data.Street)
		address.House = r.Data.House
		address.Lat = r.Data.GeoLat
		address.Lon = r.Data.GeoLon

		res = append(res, &address)
	}

	return res, nil
}

func (s *UserService) Register(ctx context.Context, username, password string) (string, error) {
    if username == "" || password == "" {
        return "", fmt.Errorf("username or password empty")
    }
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    id := rand.Intn(1000000000)
	idstr := strconv.Itoa(id)
    if err := s.service.Create(ctx, idstr, username, string(hash)); err != nil {
        return "", err
    }
    return idstr, nil
}

func (s *UserService) Authenticate(ctx context.Context, username, password string) (string, error) {
    user, err := s.service.FindUser(ctx, username)
    if err != nil {
        return "", err
    }
    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
        return "", fmt.Errorf("invalid credentials")
    }
    return user.ID, nil
}