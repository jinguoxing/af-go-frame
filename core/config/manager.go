package config

const (
	//ProjectEnvKey recognize env of project, dev, release, production
	ProjectEnvKey = "PROJECT_ENV"
	//ProjectPrefix project env key prefix
	ProjectPrefix = "PROJECT_PREFIX"
	//DefaultPrefix  default config file name prefix
	DefaultPrefix = "config"
)

var manager *Manager

//Manager config manager, hold the config
type Manager struct {
	Config
}

//Init config source, inconvenient for user, because of circular reference
func Init(sources ...Source) {
	c := New(WithSource(sources...))
	if err := c.Load(); err != nil {
		panic(err)
	}
	manager = &Manager{Config: c}
}

func Load() {
	if err := manager.Config.Load(); err != nil {
		panic(err)
	}
}

//Scan any type
func Scan[T any](keys ...string) T {
	key := ""
	if len(keys) > 0 && keys[0] != "" {
		key = keys[0]
	}
	var data T
	if key == "" {
		if err := manager.Config.Scan(&data); err != nil {
			panic(err)
		}
		return data
	}
	value := manager.Config.Value(key)
	if err := value.Scan(&data); err != nil {
		panic(err)
	}
	return data
}

func GetValue(key string) Value {
	return manager.Config.Value(key)
}

func Watch(key string, o Observer) error {
	return manager.Config.Watch(key, o)
}

func Close() error {
	return manager.Config.Close()
}
