package env

import (
	"os"
	"strings"

	"github.com/jinguoxing/af-go-frame/core/config"
)

type env struct {
	prefixs []string
}

func NewSource(prefixs ...string) config.Source {
	prefix := os.Getenv(config.ProjectPrefix)
	if prefix != "" && len(prefixs) <= 0 {
		prefixs = strings.Split(prefix, ";")
	}
	return &env{prefixs: prefixs}
}

func (e *env) Load() (kv []*config.KeyValue, err error) {
	return e.load(os.Environ()), nil
}

func (e *env) load(envStrings []string) []*config.KeyValue {
	var kv []*config.KeyValue
	for _, envstr := range envStrings {
		var k, v string
		subs := strings.SplitN(envstr, "=", 2) //nolint:gomnd
		k = subs[0]
		if len(subs) > 1 {
			v = subs[1]
		}

		if len(e.prefixs) > 0 {
			p, ok := matchPrefix(e.prefixs, k)
			if !ok || len(p) == len(k) {
				continue
			}
			// trim prefix
			k = strings.TrimPrefix(k, p)
			k = strings.TrimPrefix(k, "_")
		}

		if len(k) != 0 {
			kv = append(kv, &config.KeyValue{
				Key:   k,
				Value: []byte(v),
			})
		}
	}
	return kv
}

func (e *env) Watch() (config.Watcher, error) {
	w, err := NewWatcher()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func matchPrefix(prefixes []string, s string) (string, bool) {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return p, true
		}
	}
	return "", false
}
