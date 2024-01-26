// Code generated by DIGEN; DO NOT EDIT.
// This file was generated by Dependency Injection Container Generator 0.1.0 (built at 2023-10-21T19:51:59Z).
// See docs at https://github.com/strider2038/digen

package di

import (
	"cmd/app/config"
	meeting_repository "cmd/app/entities/meeting/repository"
	"cmd/di/internal"
	"context"
	"net/http"
	"sync"
)

type Container struct {
	mu *sync.Mutex
	c  *internal.Container
}

type Injector func(c *Container) error

func NewContainer(
	config config.Params,
	injectors ...Injector,
) (*Container, error) {
	c := &Container{
		mu: &sync.Mutex{},
		c:  internal.NewContainer(),
	}

	c.c.SetConfig(config)

	for _, inject := range injectors {
		err := inject(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Container) Server(ctx context.Context) (*http.Server, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	s := c.c.Server(ctx)
	err := c.c.Error()
	if err != nil {
		return nil, err
	}

	return s, err
}

func SetMeetingRepository(s meeting_repository.MeetingsRepository) Injector {
	return func(c *Container) error {
		c.c.Repositories().(*internal.RepositoryContainer).SetMeetingRepository(s)

		return nil
	}
}

func (c *Container) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.c.Close()
}
