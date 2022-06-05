package domain

// ※全部のカラムは今ない（だるい）
// enumはconstで作りたい
// https://developers.google.com/assistant/conversational/reference/rest/v1/TopLevel/fulfill#request-body
// omitemptyついてるカラムはオプション無いのは必須
// RequestPayloadのomitemptyは意味ないが一応
type RequestPayloadGoogleAssistant struct {
	*Handler `json:"handler"`
	*Intent  `json:"intent"`
	*Scene   `json:"scene,omitempty"`
	*Session `json:"session"`
	*User    `json:"user"`
	*Home    `json:"home,omitempty"`
	*Device  `json:"device"`
	*Context `json:"context,omitempty"`
}

type ResponsePayloadGoogleAssistant struct {
	*Prompt   `json:"prompt,omitempty"`
	*Scene    `json:"scene,omitempty"`
	*Session  `json:"session,omitempty"`
	*User     `json:"user,omitempty"`
	*Home     `json:"home,omitempty"`
	*Device   `json:"device,omitempty"`
	*Expected `json:"expected,omitempty"`
}

type Handler struct {
	Name string `json:"name"`
}

type Intent struct {
	Name    string `json:"name"`
	*Params `json:"params"`
	Query   string `json:"query"`
}

type Scene struct {
	Name              string `json:"name"`
	SlotFillingStatus string `json:"slotFillingStatus"`
	*Slots            `json:"slots"`
	*Next             `json:"next"`
}

type Session struct {
	Id            string `json:"id"`
	Params        `json:"params"`
	TypeOverrides *[]TypeOverride `json:"typeOverrides"`
	LanguageCode  string          `json:"languageCode"`
}

type User struct {
}

type Home struct {
}

type Device struct {
}

type Context struct {
}

type Prompt struct {
	Override     bool    `json:"override,omitempty"` //trueのとき書き換える
	FirstSimple  *Simple `json:"firstSimple,omitempty"`
	*Content     `json:"content,omitempty"`
	LastSimple   *Simple       `json:"lastSimple,omitempty"`
	Suggestions  *[]Suggestion `json:"suggestions,omitempty"`
	*Link        `json:"link,omitempty"`
	*Canvas      `json:"canvas,omitempty"`
	*OrderUpdate `json:"orderUpdate,omitempty"`
}

type Simple struct {
	Speech string `json:"speech,omitempty"`
	Text   string `json:"text,omitempty"`
}

type Content struct {
	*Card             `json:"card,omitempty"`
	*Image            `json:"image,omitempty"`
	*Table            `json:"table,omitempty"`
	*Media            `json:"media,omitempty"`
	*Collection       `json:"collection,omitempty"`
	*List             `json:"list,omitempty"`
	*CollectionBrowse `json:"collectionBrowse,omitempty"`
}

type Card struct {
	Title     string `json:"title,omitempty"`
	Subtitle  string `json:"subtitle,omitempty"`
	Text      string `json:"text"` // imageがない限り必須
	*Image    `json:"image,omitempty"`
	ImageFill string `json:"imageFill,omitempty"`
	Button    *Link  `json:"button,omitempty"`
}

type Image struct {
	Url    string `json:"url"`
	Alt    string `json:"alt"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

type Table struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	*Image   `json:"image"`
	Columns  *[]TableColumn `json:"columns"`
	Rows     *[]TableRow    `json:"rows"`
	Button   *[]Link        `json:"button"`
}

type TableColumn struct {
	Header string `json:"header"`
	Align  string `json:"align,omitempty"`
}

type TableRow struct {
	Cells   *[]TableCell `json:"cells"`
	Divider bool         `json:"divider"`
}

type TableCell struct {
	Text string `json:"text"`
}

type Media struct {
	MediaType             string         `json:"mediaType"`
	StartOffset           string         `json:"startOffset"`
	OptionalMediaControls *[]string      `json:"optionalMediaControls"`
	MediaObjects          *[]MediaObject `json:"mediaObjects"`
	RepeatMode            string         `json:"repeatMode"`
	FirstMediaObjectIndex int            `json:"firstMediaObjectIndex"`
}

type MediaObject struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Url         string      `json:"url"`
	Image       *MediaImage `json:"image"`
}

type MediaImage struct {
	Large *Image `json:"large"`
	Icon  *Image `json:"icon"`
}

type Collection struct {
	Title     string            `json:"title"`
	Subtitle  string            `json:"subtitle"`
	Items     *[]CollectionItem `json:"items"`
	ImageFill string            `json:"imageFill"`
}

type CollectionItem struct {
	Key string `json:"key"`
}

type List struct {
	Title    string      `json:"title,omitempty"`
	Subtitle string      `json:"subtitle,omitempty"`
	Items    *[]ListItem `json:"items"`
}

type ListItem struct {
	Key string `json:"key"`
}

type CollectionBrowse struct {
	Items     *[]Item `json:"items"`
	ImageFill string  `json:"imageFill"`
}

type Item struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Footer        string `json:"footer"`
	Image         `json:"image"`
	OpenUriAction *OpenUri `json:"openUriAction"`
}

type OpenUri struct {
	Url  string `json:"url"`
	Hint string `json:"hint"`
}

type Suggestion struct {
	Title string `json:"title"`
}

type Link struct {
	Name string   `json:"name"`
	Open *OpenUri `json:"open"`
}

type Canvas struct {
	Url                    string         `json:"url"`
	Data                   *[]interface{} `json:"data,omitempty"`
	SuppressMic            bool           `json:"suppressMic,omitempty"`
	*ContinuousMatchConfig `json:"continuousMatchConfig"`
}

type ContinuousMatchConfig struct {
	ExpectedPhtases *[]ExpectedPhtase `json:"expectedPhrases"`
	DurationSeconds int               `json:"durationSeconds"`
}

type ExpectedPhtase struct {
	Phrase             string    `json:"phrase"`
	AlternativePhrases *[]string `json:"alternativePhrases"`
}

type OrderUpdate struct {
	Type string `json:"type"` //非推奨
}

type Expected struct {
}

type Params struct {
	Original string      `json:"original"`
	Resolved interface{} `json:"resolved"`
}

type Slots struct {
	Mode    string      `json:"mode"`
	Status  string      `json:"status"`
	Value   interface{} `json:"value"`
	Updated bool        `json:"updated"`
	Prompt  `json:"prompt"`
}

type Next struct {
	Name string `json:"name"`
}

type TypeOverride struct {
	Name    string `json:"name"`
	Mode    string `json:"mode"`
	Synonym `json:"synonym"`
}

type Synonym struct {
	Entries *[]Entrie `json:"entries"`
}

type Entrie struct {
}
