package erh

//Motd message of the day
type Motd struct {
	Msg string `json:"msg"`
	Url string `json:"url"`
}

//Symbol contains currency code and description
type Symbol struct {
	Description string `json:"description"`
	Code        string `json:"code"`
}

//SymbolsResponse contains response of the symbols method
type SymbolsResponse struct {
	Motd    Motd              `json:"motd"`
	Success bool              `json:"success"`
	Symbols map[string]Symbol `json:"symbols"`
}

//ConvertQuery contains query arguments of convert method
type ConvertQuery struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

//ConvertInfo contains converted value
type ConvertInfo struct {
	Rate float64 `json:"rate"`
}

//ConvertResponse contains response of convert method
type ConvertResponse struct {
	Motd       Motd         `json:"motd"`
	Success    bool         `json:"success"`
	Query      ConvertQuery `json:"query"`
	Info       ConvertInfo  `json:"info"`
	Historical bool         `json:"historical"`
	Date       string       `json:"date"`
	Result     float64      `json:"result"`
}

//TimeSeriesResponse contains rime series response
type TimeSeriesResponse struct {
	Motd       Motd                          `json:"motd"`
	Success    bool                          `json:"success"`
	TimeSeries bool                          `json:"timeseries"`
	Base       string                        `json:"base"`
	StartDate  string                        `json:"start_date"`
	EndDate    string                        `json:"end_date"`
	Rates      map[string]map[string]float64 `json:"rates"`
}

//HistoricalResponse contains historical response
type HistoricalResponse struct {
	Motd       Motd               `json:"Motd"`
	Success    bool               `json:"success"`
	Historical bool               `json:"historical"`
	Base       string             `json:"base"`
	Date       string             `json:"date"`
	Rates      map[string]float64 `json:"rates"`
}

//LatestResponse contains latest rates response
type LatestResponse struct {
	Motd    Motd               `json:"motd"`
	Success bool               `json:"success"`
	Base    string             `json:"base"`
	Date    string             `json:"date"`
	Rates   map[string]float64 `json:"rates"`
}
