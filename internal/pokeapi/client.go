package pokeapi

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gv/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	revalidationTime, err := time.ParseDuration("10s")
	if err != nil {
		log.Fatalf("failed to parse revalidation time: %v", err)
		panic(err)
	}

	cache := pokecache.NewCache(revalidationTime)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}

func (c *Client) get(url string) ([]byte, error) {
	cached, exists := c.cache.Get(url)
	if exists {
		return cached, nil
	}
	res, err := c.httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return nil, errors.New(res.Status)
	}

	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)

	return body, err
}
