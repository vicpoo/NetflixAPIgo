package entities

type Video struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
	UserID      int    `json:"user_id"`
}

// Getters
func (v *Video) GetID() int             { return v.ID }
func (v *Video) GetTitle() string       { return v.Title }
func (v *Video) GetDescription() string { return v.Description }
func (v *Video) GetURL() string         { return v.URL }
func (v *Video) GetUserID() int         { return v.UserID }

// Setters
func (v *Video) SetID(id int)               { v.ID = id }
func (v *Video) SetTitle(title string)      { v.Title = title }
func (v *Video) SetDescription(desc string) { v.Description = desc }
func (v *Video) SetURL(url string)          { v.URL = url }
func (v *Video) SetUserID(userID int)       { v.UserID = userID }
