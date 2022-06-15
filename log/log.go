package log

import (
	llog "log"

	"github.com/vishal1132/secured-layers/security"

	"go.uber.org/zap"
)

type log struct {
	security security.SecurityCompliance
	*zap.Logger
}

func (l *log) Println(k interface{}) {
	l.security.EncryptRegisteredFields(k)
	llog.Println(k)
}

type option func(l *log)

func New(opts ...option) *log {
	l := &log{
		Logger: zap.New(zap.L().Core()),
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func WithSecurity(security security.SecurityCompliance) option {
	return func(l *log) {
		l.security = security
	}
}
