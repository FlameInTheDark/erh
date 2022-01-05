package erh

type Motd struct {
	Msg string `json:"msg"`
	Url string `json:"url"`
}

type Symbol struct {
	Description string `json:"description"`
	Code        string `json:"code"`
}

type SymbolsResponse struct {
	Motd    Motd              `json:"motd"`
	Success bool              `json:"success"`
	Symbols map[string]Symbol `json:"symbols"`
}

type ConvertQuery struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type ConvertInfo struct {
	Rate float64 `json:"rate"`
}

type ConvertResponse struct {
	Motd       Motd         `json:"motd"`
	Success    bool         `json:"success"`
	Query      ConvertQuery `json:"query"`
	Info       ConvertInfo  `json:"info"`
	Historical bool         `json:"historical"`
	Date       string       `json:"date"`
	Result     float64      `json:"result"`
}

type TimeSeriesResponse struct {
	Motd       Motd                          `json:"motd"`
	Success    bool                          `json:"success"`
	TimeSeries bool                          `json:"timeseries"`
	Base       string                        `json:"base"`
	StartDate  string                        `json:"start_date"`
	EndDate    string                        `json:"end_date"`
	Rates      map[string]map[string]float64 `json:"rates"`
}

type HistoricalResponse struct {
	Motd       Motd               `json:"Motd"`
	Success    bool               `json:"success"`
	Historical bool               `json:"historical"`
	Base       string             `json:"base"`
	Date       string             `json:"date"`
	Rates      map[string]float64 `json:"rates"`
}

type LatestResponse struct {
	Motd    Motd               `json:"motd"`
	Success bool               `json:"success"`
	Base    string             `json:"base"`
	Date    string             `json:"date"`
	Rates   map[string]float64 `json:"rates"`
}
