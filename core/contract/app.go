package contract


const AppKey = "af:app"

type AppInfo interface {
    ID() string
    Name() string

    Endpoint() []string
}

