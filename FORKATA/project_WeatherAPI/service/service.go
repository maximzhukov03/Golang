package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

type Coins struct {
    Bitcoin  PriceData `json:"bitcoin,omitempty"`
    Ethereum PriceData `json:"ethereum,omitempty"`
}

type PriceData struct {
    USD float64 `json:"usd"`
}

type CryptoService struct{
	client *http.Client
	cache *redis.Client
	url string
}

type Service interface{
	GetBitcoin(context.Context) (Coins, error)
	GetEthereum(context.Context) (Coins, error)
}

type ProxyService struct{
	service Service
}


func NewService(r *redis.Client) *CryptoService{
	return &CryptoService{
		client: &http.Client{},
		cache: r,
		url: "https://api.coingecko.com/api/v3/simple/price?ids=",
	}
}

func (s *CryptoService) GetBitcoin(ctx context.Context) (Coins, error){
	var coin Coins
	key := "bitcoin:usd"
	res, err := s.cache.Get(key).Result()
	if err == nil{
		log.Println("Лезет в кэш")
		err := json.Unmarshal([]byte(res), &coin)
		if err != nil{
			return Coins{}, err
		}
		return coin, nil
	}
	result, err := http.NewRequestWithContext(ctx, http.MethodGet , s.url+"bitcoin&vs_currencies=usd", nil)
	if err != nil{
		return Coins{}, err
	}
	resp, err := s.client.Do(result)
    if err != nil {
        return Coins{}, err
    }
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&coin)
	if err != nil{
		return Coins{}, err
	}

	req, err := json.Marshal(coin)
	if err != nil{
		return Coins{}, nil
	}
	_ = s.cache.Set(key, req, 2 * time.Minute)
	return coin, nil
}

func (s *CryptoService) GetEthereum(ctx context.Context) (Coins, error){
	var coin Coins
	key := "ethereum:usd"
	res, err := s.cache.Get(key).Result()
	if err == nil{
		log.Println("Лезет в кэш")
		err := json.Unmarshal([]byte(res), &coin)
		if err != nil{
			return Coins{}, err
		}
		return coin, nil
	}
	result, err := http.NewRequestWithContext(ctx, http.MethodGet , s.url+"ethereum&vs_currencies=usd", nil)
	if err != nil{
		return Coins{}, err
	}
	resp, err := s.client.Do(result)
    if err != nil {
        return Coins{}, err
    }
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&coin)
	if err != nil{
		return Coins{}, err
	}

	req, err := json.Marshal(coin)
	if err != nil{
		return Coins{}, nil
	}
	_ = s.cache.Set(key, req, 2 * time.Minute)
	return coin, nil

}	
	// var coin Coins
	// result, err := http.NewRequestWithContext(ctx, http.MethodGet ,"https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd", nil)
	// if err != nil{
	// 	return Coins{}, err
	// }
	// resp, err := s.client.Do(result)
    // if err != nil {
    //     return Coins{}, err
    // }
	// err = json.NewDecoder(resp.Body).Decode(&coin)
	// if err != nil{
	// 	return Coins{}, err
	// }
	// return coin, nil