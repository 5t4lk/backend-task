package news

type NewListInformation struct {
	ClubName            string              `xml:"ClubName"`
	ClubWebsiteURL      string              `xml:"ClubWebsiteURL"`
	NewsletterNewsItems NewsletterNewsItems `xml:"NewsletterNewsItems"`
}
type NewsletterNewsItem struct {
	ArticleURL        string `xml:"ArticleURL"`
	IsPublished       string `xml:"IsPublished"`
	LastUpdateDate    string `xml:"LastUpdateDate"`
	NewsArticleID     string `xml:"NewsArticleID"`
	OptaMatchID       string `xml:"OptaMatchId"`
	PublishDate       string `xml:"PublishDate"`
	Taxonomies        string `xml:"Taxonomies"`
	TeaserText        string `xml:"TeaserText"`
	ThumbnailImageURL string `xml:"ThumbnailImageURL"`
	Title             string `xml:"Title"`
}
type NewsletterNewsItems struct {
	NewsletterNewsItem []NewsletterNewsItem `xml:"NewsletterNewsItem"`
}

type MainData struct {
	ClubName       string      `json:"ClubName"`
	ClubWebsiteURL string      `json:"ClubWebsiteURL"`
	NewsArticle    NewsArticle `json:"NewsArticle"`
}

type NewsArticle struct {
	ArticleURL        string `json:"ArticleURL"`
	NewsArticleID     string `json:"NewsArticleID"`
	PublishDate       string `json:"PublishDate"`
	Taxonomies        string `json:"Taxonomies"`
	TeaserText        string `json:"TeaserText"`
	Subtitle          string `json:"Subtitle"`
	ThumbnailImageURL string `json:"ThumbnailImageURL"`
	Title             string `json:"Title"`
	GalleryImageURLs  string `json:"GalleryImageURLs"`
	VideoURL          string `json:"VideoURL"`
	OptaMatchID       string `json:"OptaMatchId"`
	LastUpdateDate    string `json:"LastUpdateDate"`
	IsPublished       string `json:"IsPublished"`
}
