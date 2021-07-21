package freqtrade

type Status struct {
	DryRun        bool
	BotName       string `json:"bot_name"` // name
	State         string // stopped | stop_buy | running | unknown
	RunMode       string // dry_run | live | unknown
	Strategy      string
	MaxOpenTrades int `json:"max_open_trades"`
}

func (s *Status) StateStr() string {
	if s.MaxOpenTrades == 0 && s.State != "unknown" {
		return "stop_buy"
	}
	return s.State
}

func (s *Status) StateInt() int {
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

func (s *Status) ModeInt() int {
	switch s.RunMode {
	case "dry_run":
		return 0
	case "live":
		return 1
	default:
		return -1
	}
}

func EmptyStatus() *Status {
	return &Status{
		DryRun:  true,
		BotName: "unknown",
		State:   "unknown",
		RunMode: "unknown",
	}
}

func NewStatus(conn *Connection) *Status {
	var name, expireAt, query = Connector(conn, API_STATUS)
	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = new(Status)
		err := conn.GET(name, query, &target)
		return target, err
	}); err == nil && data != nil {
		return data.(*Status)
	}

	return EmptyStatus()
}
