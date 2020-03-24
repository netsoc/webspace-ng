package webspace

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/netsoc/webspace-ng/webspaced/internal/config"
)

// Traefik manages webspace configuration for Traefik
type Traefik struct {
	config *config.Config
	redis  *redis.Client
}

// NewTraefik creates a new Traefik instance
func NewTraefik(cfg *config.Config) (*Traefik, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Traefik.Redis.Addr,
		DB:   cfg.Traefik.Redis.DB,
	})

	return &Traefik{
		cfg,
		client,
	}, nil
}

// UpdateConfig generates new Traefik configuration for a webspace
func (t *Traefik) UpdateConfig(ws *Webspace, running bool) error {
	// TODO: clear old config

	if !running {
		return nil
	}

	n := ws.InstanceName()

	// TODO: Wait and retry up to a maximum amount
	time.Sleep(2 * time.Second)
	addr, err := ws.GetIP()
	if err != nil {
		return fmt.Errorf("failed to get instance IP address: %w", err)
	}

	backend := fmt.Sprintf("http://%v:%v", addr, ws.Config.HTTPPort)

	rules := make([]string, len(ws.Domains))
	for i, d := range ws.Domains {
		rules[i] = fmt.Sprintf("Host(`%v`)", d)
	}
	rule := strings.Join(rules, " || ")

	// TODO: https (ssl termination _and_ passthrough)
	if _, err = t.redis.TxPipelined(func(pipe redis.Pipeliner) error {
		pipe.Set(fmt.Sprintf("traefik/http/services/%v/loadbalancer/servers/0/url", n), backend, 0)
		pipe.Set(fmt.Sprintf("traefik/http/routers/%v/service", n), n, 0)
		pipe.Set(fmt.Sprintf("traefik/http/routers/%v/rule", n), rule, 0)
		pipe.Set(fmt.Sprintf("traefik/http/routers/%v/entrypoints/0", n), t.config.Traefik.HTTPEntryPoint, 0)
		return nil
	}); err != nil {
		return fmt.Errorf("failed to set redis values: %w", err)
	}

	return nil
}
