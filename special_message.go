package hanu

import (
	//"regexp"
  //"encoding/json"
  //"github.com/guregu/null"
 "strconv"
	"strings"
  "fmt"
  "github.com/redventures/fuse-server/models"
)

type SpecialMessageInterface interface {
	IsMessage() bool
	IsFrom(user string) bool
	IsDirectMessage() bool
	IsRelevantFor(user string) bool
	User() string
}

type SpecialMessage struct {
	ID      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	UserID  string `json:"user"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
  Fallback string `json:"fallback"`
  Color string `json:"color"`
  Pretext string `json:"pretext"`
  Title string `json:"title"`
  Text string `json:"text"`
  Fields []Field `json:"fields"`
}

type Field struct {
  Title string `json:"title"`
  Value string `json:"value"`
  Short bool  `json:"short"`
}

func (f *Field) SetFieldTitle(s string) {
  f.Title = s
}

func (f *Field) SetFieldValue(s int64, n int64) {
  m := strconv.FormatInt(s, 10)
  v := strconv.FormatInt(n, 10)
  newLine := "Phone numbers remaining: " + m + " Phone Numbers Total: " + v
  f.Value = newLine
}

func (f *Field) SetFieldShort(s bool) {
  f.Short = s
}

// User returns the message's User ID
func (m SpecialMessage) User() string {
	return m.UserID
}

//IsMessage checks if it is a Message or some other kind of processing information
func (m SpecialMessage) IsMessage() bool {
	return m.Type == "message"
}
//
// IsFrom checks the sender of the message
func (m SpecialMessage) IsFrom(user string) bool {
	return m.UserID == user
}


func (m *SpecialMessage) SetSpecialMessage(sm models.Pool) {
  field := Field{}
  field.SetFieldTitle(sm.Name)
  field.SetFieldValue(sm.PhoneNumbersAvailable.Int64, sm.PhoneNumbersTotal.Int64)
  field.SetFieldShort(true)

  slice := []Field{field}

  attachment := Attachment{
    Fallback: "Fuse Status",
    Color: "#7CD197",
    Pretext: "Fuse Status",
    Title: "Fuse Status",
    Text: "Fuse Status",
    Fields: slice,
  }
  fmt.Printf("%+v\n", attachment)

  attach := []Attachment{attachment}
  m.Attachments = attach;
}


// // IsDirectMessage checks if the message is received using a direct messaging channel
func (m SpecialMessage) IsDirectMessage() bool {
	return strings.HasPrefix(m.Channel, "D")
}
//
// IsRelevantFor checks if the message is relevant for a user
func (m SpecialMessage) IsRelevantFor(user string) bool {
	return m.IsMessage() && !m.IsFrom(user) && (m.IsDirectMessage())
}
