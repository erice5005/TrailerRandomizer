package imdb

type KeywordSearchResponse struct {
	SearchType string `json:"searchType"`
	Expression string `json:"expression"`
	Results    []*KeywordResult `json:"results"`
	ErrorMessage string `json:"errorMessage"`
}

type KeywordResult struct {
	ID          string `json:"id"`
	ResultType  string `json:"resultType"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type KeywordElementResponse struct {
	Keyword string `json:"keyword"`
	Items   []*KeywordItem `json:"items"`
	ErrorMessage string `json:"errorMessage"`
}

type KeywordItem struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Year       string `json:"year"`
	Image      string `json:"image"`
	ImDbRating string `json:"imDbRating"`
}

type ReviewsResponse struct {
	ImDbID    string `json:"imDbId"`
	Title     string `json:"title"`
	FullTitle string `json:"fullTitle"`
	Type      string `json:"type"`
	Year      string `json:"year"`
	Items     []*Review`json:"items"`
	ErrorMessage string `json:"errorMessage"`
}

type Review struct {
	Username        string `json:"username"`
	UserURL         string `json:"userUrl"`
	ReviewLink      string `json:"reviewLink"`
	WarningSpoilers bool   `json:"warningSpoilers"`
	Date            string `json:"date"`
	Rate            string `json:"rate"`
	Helpful         string `json:"helpful"`
	Title           string `json:"title"`
	Content         string `json:"content"`
}

type TitleResponse struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	OriginalTitle  string `json:"originalTitle"`
	FullTitle      string `json:"fullTitle"`
	Type           string `json:"type"`
	Year           string `json:"year"`
	Image          string `json:"image"`
	ReleaseDate    string `json:"releaseDate"`
	RuntimeMins    string `json:"runtimeMins"`
	RuntimeStr     string `json:"runtimeStr"`
	Plot           string `json:"plot"`
	PlotLocal      string `json:"plotLocal"`
	PlotLocalIsRtl bool   `json:"plotLocalIsRtl"`
	Awards         string `json:"awards"`
	Directors      string `json:"directors"`
	DirectorList   []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"directorList"`
	Writers    string `json:"writers"`
	WriterList []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"writerList"`
	Stars    string `json:"stars"`
	StarList []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"starList"`
	ActorList []struct {
		ID          string `json:"id"`
		Image       string `json:"image"`
		Name        string `json:"name"`
		AsCharacter string `json:"asCharacter"`
	} `json:"actorList"`
	FullCast  interface{} `json:"fullCast"`
	Genres    string      `json:"genres"`
	GenreList []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"genreList"`
	Companies   string `json:"companies"`
	CompanyList []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"companyList"`
	Countries   string `json:"countries"`
	CountryList []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"countryList"`
	Languages    string `json:"languages"`
	LanguageList []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"languageList"`
	ContentRating    string      `json:"contentRating"`
	ImDbRating       string      `json:"imDbRating"`
	ImDbRatingVotes  string      `json:"imDbRatingVotes"`
	MetacriticRating string      `json:"metacriticRating"`
	Ratings          interface{} `json:"ratings"`
	Wikipedia        interface{} `json:"wikipedia"`
	Posters          interface{} `json:"posters"`
	Images           interface{} `json:"images"`
	Trailer          struct {
		ImDbID           string `json:"imDbId"`
		Title            string `json:"title"`
		FullTitle        string `json:"fullTitle"`
		Type             string `json:"type"`
		Year             string `json:"year"`
		VideoID          string `json:"videoId"`
		VideoTitle       string `json:"videoTitle"`
		VideoDescription string `json:"videoDescription"`
		ThumbnailURL     string `json:"thumbnailUrl"`
		UploadDate       string `json:"uploadDate"`
		Link             string `json:"link"`
		LinkEmbed        string `json:"linkEmbed"`
		ErrorMessage     string `json:"errorMessage"`
	} `json:"trailer"`
	BoxOffice struct {
		Budget                   string `json:"budget"`
		OpeningWeekendUSA        string `json:"openingWeekendUSA"`
		GrossUSA                 string `json:"grossUSA"`
		CumulativeWorldwideGross string `json:"cumulativeWorldwideGross"`
	} `json:"boxOffice"`
	Tagline     string   `json:"tagline"`
	Keywords    string   `json:"keywords"`
	KeywordList []string `json:"keywordList"`
	Similars    []struct {
		ID         string `json:"id"`
		Title      string `json:"title"`
		FullTitle  string `json:"fullTitle"`
		Year       string `json:"year"`
		Image      string `json:"image"`
		Plot       string `json:"plot"`
		Directors  string `json:"directors"`
		Stars      string `json:"stars"`
		Genres     string `json:"genres"`
		ImDbRating string `json:"imDbRating"`
	} `json:"similars"`
	TvSeriesInfo  interface{} `json:"tvSeriesInfo"`
	TvEpisodeInfo interface{} `json:"tvEpisodeInfo"`
	ErrorMessage  string      `json:"errorMessage"`
}