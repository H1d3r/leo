package leo

import (
	"errors"

	"github.com/zan8in/leo/pkg/ssh"
)

type Execute struct {
	options *Options
}

func NewExecute(options *Options) *Execute {
	return &Execute{options: options}
}

func (e *Execute) start(host, username, password string, m any) error {
	service := e.options.Service
	if service == SSH_NAME {
		sshclient := m.(*ssh.SSH)
		return sshclient.AuthSSHRtries(host, username, password)
	}
	return nil
}

func (e *Execute) validateService(host string) (any, error) {
	service := e.options.Service
	if service == SSH_NAME {
		m, err := ssh.NewSSH(host, e.options.Port, e.options.Retries, e.options.Timeout)
		if err != nil {
			return m, err
		}
		return m, nil
	}

	return nil, errors.New("error")
}