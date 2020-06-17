package httpclient

import (
	"time"

	"github.com/go-kit/kit/log"
)

// Builder ...
type Builder interface {
	Build() Client
	Host(host string) Builder
	Path(path string) Builder
	Logger(l log.Logger) Builder
	TimeOut(timeout time.Duration) Builder
}

type builder struct {
	host string
	path string
	logger log.Logger
	timeOut time.Duration
}

// 	Build ...
func (b *builder) Build() Client {
	return &client{
		host:       b.host,
		path:       b.path,
		logger:     b.logger,
		timeOut:    b.timeOut,
	}
}

// Host ...
func (b *builder) Host(host string) Builder {
	b.host = host
	return b
}

// Path ...
func (b *builder) Path(path string) Builder {
	b.path = path
	return b
}

// Logger ...
func (b *builder) Logger(l log.Logger) Builder {
	b.logger = l
	return b
}

// Timeout
func (b *builder) TimeOut(timeout time.Duration) Builder {
	b.timeOut = timeout
	return b
}

// New ...
func New() Builder {
	return &builder{}
}



