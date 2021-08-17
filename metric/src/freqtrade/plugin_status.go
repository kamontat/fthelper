package freqtrade

import (
	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
)

const STATUS_CONST = "show_config"

func NewStatus() *status {
	return &status{
		DryRun:   true,
		BotName:  "unknown",
		State:    "unknown",
		RunMode:  "unknown",
		Strategy: "unknown",
	}
}

type status struct {
	DryRun        bool
	BotName       string `json:"bot_name"` // name
	State         string // stopped | stop_buy | running | unknown
	RunMode       string // dry_run | live | unknown
	Strategy      string
	MaxOpenTrades int `json:"max_open_trades"`
}

func (s *status) StateStr() string {
	if s.MaxOpenTrades == 0 && s.State != "unknown" {
		return "stop_buy"
	}
	return s.State
}

func (s *status) StateInt() int {
	switch s.StateStr() {
	case "stopped":
		return 0
	case "running":
		return 1
	case "stop_buy":
		return 2
	default:
		return -1
	}
}

func (s *status) ModeInt() int {
	switch s.RunMode {
	case "dry_run":
		return 0
	case "live":
		return 1
	default:
		return -1
	}
}

func (s *status) Name() string {
	return STATUS_CONST
}

func (s *status) Build(connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	err := connection.Http.GET(s.Name(), s)
	return s, err
}

func (s *status) Empty() bool {
	return s.RunMode == "unknown"
}

func ToStatus(connector connection.Connector) (*status, error) {
	raw, err := connector.Connect(STATUS_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*status), nil
}
