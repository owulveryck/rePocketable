package main

import "time"

type v1Payload struct {
	Archive      string   `json:"archive"`
	URL          string   `json:"url"`
	WebURL       string   `json:"web_url"`
	NaturalKey   []string `json:"natural_key"`
	FullPath     string   `json:"full_path"`
	AssetBaseURL string   `json:"asset_base_url"`
	LastPosition struct {
		TopPercentage    float64   `json:"top_percentage"`
		LastModifiedTime time.Time `json:"last_modified_time"`
	} `json:"last_position"`
	AllowScripts    bool        `json:"allow_scripts"`
	Videoclip       interface{} `json:"videoclip"`
	PreviousChapter struct {
		URL    string `json:"url"`
		WebURL string `json:"web_url"`
		Title  string `json:"title"`
	} `json:"previous_chapter"`
	NextChapter struct {
		URL    string `json:"url"`
		WebURL string `json:"web_url"`
		Title  string `json:"title"`
	} `json:"next_chapter"`
	Subjects []interface{} `json:"subjects"`
	Authors  []struct {
		Name string `json:"name"`
	} `json:"authors"`
	Cover           string        `json:"cover"`
	BookTitle       string        `json:"book_title"`
	Updated         time.Time     `json:"updated"`
	SiteStyles      []string      `json:"site_styles"`
	EpubProperties  []interface{} `json:"epub_properties"`
	Title           string        `json:"title"`
	Filename        string        `json:"filename"`
	Path            string        `json:"path"`
	Content         string        `json:"content"`
	MinutesRequired float64       `json:"minutes_required"`
	Stylesheets     []struct {
		FullPath    string `json:"full_path"`
		URL         string `json:"url"`
		OriginalURL string `json:"original_url"`
	} `json:"stylesheets"`
	Images               []interface{} `json:"images"`
	Videoclips           []interface{} `json:"videoclips"`
	AcademicExcluded     bool          `json:"academic_excluded"`
	CreatedTime          time.Time     `json:"created_time"`
	LastModifiedTime     time.Time     `json:"last_modified_time"`
	VirtualPages         int           `json:"virtual_pages"`
	HeadExtra            interface{}   `json:"head_extra"`
	HasVideo             bool          `json:"has_video"`
	Description          string        `json:"description"`
	PublisherScripts     string        `json:"publisher_scripts"`
	PublisherScriptFiles []interface{} `json:"publisher_script_files"`
}
