package types

type Config struct {
	Editor  string    `json:"editor"`
	Startup []Startup `json:"startup"`
	Monitor []Monitor `json:"monitor"`
	Daily   []Daily   `json:"daily"`
	Weekly  []Weekly  `json:"weekly"`
	Changed bool      `json:"-"`
}

type Startup struct {
	Name    string   `json:"name"`
	Enabled bool     `json:"enabled"`
	Cmd     string   `json:"cmd"`
	Args    []string `json:"args"`
	LastRun string   `json:"lastRun"`
}

type Monitor struct {
	Name    string   `json:"name"`
	Enabled bool     `json:"enabled"`
	Cmd     string   `json:"cmd"`
	Args    []string `json:"args"`
	LastRun string   `json:"lastRun"`
}

type Daily struct {
	Name    string   `json:"name"`
	Enabled bool     `json:"enabled"`
	Cmd     string   `json:"cmd"`
	Args    []string `json:"args"`
	LastRun string   `json:"lastRun"`
	Time    string   `json:"time"`
}

type Weekly struct {
	Name    string   `json:"name"`
	Enabled bool     `json:"enabled"`
	Cmd     string   `json:"cmd"`
	Args    []string `json:"args"`
	LastRun string   `json:"lastRun"`
	Time    string   `json:"time"`
	Day     string   `json:"day"`
}
