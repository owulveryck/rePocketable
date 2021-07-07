package pocket

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// RetrieveOption is the options for retrieve API.
type RetrieveOption struct {
	ConsumerKey string         `json:"consumer_key"`
	AccessToken string         `json:"access_token"`
	State       State          `json:"state,omitempty"`
	Favorite    FavoriteFilter `json:"favorite,omitempty"`
	Tag         string         `json:"tag,omitempty"`
	ContentType ContentType    `json:"contentType,omitempty"`
	Sort        Sort           `json:"sort,omitempty"`
	DetailType  DetailType     `json:"detailType,omitempty"`
	Search      string         `json:"search,omitempty"`
	Domain      string         `json:"domain,omitempty"`
	Since       int            `json:"since,omitempty"`
	Count       int            `json:"count,omitempty"`
	Offset      int            `json:"offset,omitempty"`
}

type State string

const (
	StateUnread  State = "unread"
	StateArchive State = "archive"
	StateAll     State = "all"
)

type ContentType string

const (
	ContentTypeArticle ContentType = "article"
	ContentTypeVideo   ContentType = "video"
	ContentTypeImage   ContentType = "image"
)

type Sort string

const (
	SortNewest Sort = "newest"
	SortOldest Sort = "oldest"
	SortTitle  Sort = "title"
	SortSite   Sort = "site"
)

type DetailType string

const (
	DetailTypeSimple   DetailType = "simple"
	DetailTypeComplete DetailType = "complete"
)

type FavoriteFilter string

const (
	FavoriteFilterUnspecified FavoriteFilter = ""
	FavoriteFilterUnfavorited FavoriteFilter = "0"
	FavoriteFilterFavorited   FavoriteFilter = "1"
)

type RetrieveResult struct {
	List     map[string]Item
	Status   int
	Complete int
	Since    int
}

type ItemStatus int

const (
	ItemStatusUnread   ItemStatus = 0
	ItemStatusArchived ItemStatus = 1
	ItemStatusDeleted  ItemStatus = 2
)

type ItemMediaAttachment int

const (
	ItemMediaAttachmentNoMedia  ItemMediaAttachment = 0
	ItemMediaAttachmentHasMedia ItemMediaAttachment = 1
	ItemMediaAttachmentIsMedia  ItemMediaAttachment = 2
)

type Item struct {
	ItemID        int        `json:"item_id,string"`
	ResolvedId    int        `json:"resolved_id,string"`
	GivenURL      string     `json:"given_url"`
	ResolvedURL   string     `json:"resolved_url"`
	GivenTitle    string     `json:"given_title"`
	ResolvedTitle string     `json:"resolved_title"`
	Favorite      int        `json:",string"`
	Status        ItemStatus `json:",string"`
	Excerpt       string
	IsArticle     int                 `json:"is_article,string"`
	HasImage      ItemMediaAttachment `json:"has_image,string"`
	HasVideo      ItemMediaAttachment `json:"has_video,string"`
	WordCount     int                 `json:"word_count,string"`

	// Fields for detailed response
	Tags    map[string]map[string]interface{}
	Authors map[string]map[string]interface{}
	Images  map[string]map[string]interface{}
	Videos  map[string]map[string]interface{}

	// Fields that are not documented but exist
	SortId        int  `json:"sort_id"`
	TimeAdded     Time `json:"time_added"`
	TimeUpdated   Time `json:"time_updated"`
	TimeRead      Time `json:"time_read"`
	TimeFavorited Time `json:"time_favorited"`
}

type Time time.Time

func (t Time) MarshalBinary() (data []byte, err error) {
	tim := time.Time(t)
	return tim.MarshalBinary()
}

func (t *Time) UnmarshalBinary(data []byte) error {
	var tim time.Time
	err := tim.UnmarshalBinary(data)
	if err != nil {
		return err
	}
	*t = Time(tim)
	return nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(bytes.Trim(b, `"`)), 10, 64)
	if err != nil {
		return err
	}

	*t = Time(time.Unix(i, 0))

	return nil
}

// URL returns ResolvedURL or GivenURL
func (item Item) URL() string {
	url := item.ResolvedURL
	if url == "" {
		url = item.GivenURL
	}
	return url
}

// Title returns ResolvedTitle or GivenTitle
func (item Item) Title() string {
	title := item.ResolvedTitle
	if title == "" {
		title = item.GivenTitle
	}
	return title
}

func (p *Pocket) Get(ctx context.Context) error {
	request := RetrieveOption{
		ConsumerKey: p.ConsumerKey,
		AccessToken: p.auth.AccessToken,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.URL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("X-Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := p.downloader.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("got response %d; X-Error=[%s]", resp.StatusCode, resp.Header.Get("X-Error"))
	}

	defer resp.Body.Close()
	var res RetrieveResult
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}
	for _, item := range res.List {
		select {
		case <-ctx.Done():
			return nil
		case p.ItemsC <- item:
		}
	}
	return nil
}
