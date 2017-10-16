package models

// VacancySearchPage is for search page from HeadHunter API
type VacancySearchPage struct {
	Clusters interface{} `json:"clusters"`
	Items    []struct {
		Salary struct {
			To       *int   `json:"to"`
			Gross    bool   `json:"gross"`
			From     *int   `json:"from"`
			Currency string `json:"currency"`
		} `json:"salary"`
		Snippet struct {
			Requirement    interface{} `json:"requirement"`
			Responsibility string      `json:"responsibility"`
		} `json:"snippet"`
		Archived bool   `json:"archived"`
		Premium  bool   `json:"premium"`
		Name     string `json:"name"`
		Area     struct {
			URL  string `json:"url"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		URL               string        `json:"url"`
		CreatedAt         string        `json:"created_at"`
		AlternateURL      string        `json:"alternate_url"`
		ApplyAlternateURL string        `json:"apply_alternate_url"`
		Relations         []interface{} `json:"relations"`
		Employer          struct {
			LogoUrls struct {
				Num90    string `json:"90"`
				Num240   string `json:"240"`
				Original string `json:"original"`
			} `json:"logo_urls"`
			VacanciesURL string `json:"vacancies_url"`
			Name         string `json:"name"`
			URL          string `json:"url"`
			AlternateURL string `json:"alternate_url"`
			ID           string `json:"id"`
			Trusted      bool   `json:"trusted"`
		} `json:"employer"`
		ResponseLetterRequired bool   `json:"response_letter_required"`
		PublishedAt            string `json:"published_at"`
		Address                struct {
			Building      string        `json:"building"`
			City          string        `json:"city"`
			Description   interface{}   `json:"description"`
			Metro         interface{}   `json:"metro"`
			MetroStations []interface{} `json:"metro_stations"`
			Raw           interface{}   `json:"raw"`
			Street        string        `json:"street"`
			Lat           float64       `json:"lat"`
			Lng           float64       `json:"lng"`
			ID            string        `json:"id"`
		} `json:"address"`
		Department        interface{} `json:"department"`
		SortPointDistance interface{} `json:"sort_point_distance"`
		Type              struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		ID string `json:"id"`
	} `json:"items"`
	Pages        int         `json:"pages"`
	Arguments    interface{} `json:"arguments"`
	Found        int         `json:"found"`
	AlternateURL string      `json:"alternate_url"`
	PerPage      int         `json:"per_page"`
	Page         int         `json:"page"`
	Errors       []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"errors"`
}
