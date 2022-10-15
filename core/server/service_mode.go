package server

const (
    NotSetMod = "not-set"
    // DevMode means development mode.
    DevMode = "dev"
    // TestMode means test mode.
    TestMode = "test"
    // RtMode means regression test mode.
    RtMode = "rt"
    // PreMode means pre-release mode.
    PreMode = "pre"
    // ProMode means production mode.
    ProMode = "pro"
)

var (
    currentMode = NotSetMod
)

// Mode
func Mode() string {
    if currentMode == NotSetMod {

    }

    return currentMode
}

// IsDevelop
func IsDevelop() bool {
    return Mode() == DevMode
}

// IsTesting
func IsTesting() bool {
    return Mode() == TestMode
}
