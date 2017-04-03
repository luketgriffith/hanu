package hanu

import (
	"fmt"
	// "encoding/json"
	"golang.org/x/net/websocket"
	"github.com/redventures/fuse-server/models"
	"github.com/sbstjn/allot"
)

// ConversationInterface is the interface for a conversation
type ConversationInterface interface {
	Integer(name string) (int, error)
	String(name string) (string, error)
	Reply(text string, a ...interface{})
	Match(position int) (string, error)
	SpecialReply(res models.Pool, a ...interface{})
	SetConnection(connection Connection)

	send(msg MessageInterface)
}

// Connection is the needed interface for a connection
type Connection interface {
	Send(ws *websocket.Conn, v interface{}) (err error)
}

// Conversation stores message, command and socket information and is passed
// to the handler function
type Conversation struct {
	message Message
	specialMessage SpecialMessage
	match   allot.MatchInterface
	socket  *websocket.Conn

	connection Connection
}

func (c *Conversation) send(msg MessageInterface) {
	if c.socket != nil {
		c.connection.Send(c.socket, msg)
	}
}

// SetConnection sets the conversation connection
func (c *Conversation) SetConnection(connection Connection) {
	c.connection = connection
}

// Reply sends message using the socket to Slack
func (c *Conversation) Reply(text string, a ...interface{}) {
	prefix := ""

	if !c.message.IsDirectMessage() {
		prefix = "<@" + c.message.User() + ">: "
	}

	msg := c.message
	msg.SetText(prefix + fmt.Sprintf(text, a...))

	c.send(msg)
}

// Reply sends message using the socket to Slack
func (c *Conversation) SpecialReply(res models.Pool, a ...interface{}) {
	// prefix := ""
	fmt.Printf("terrible mate")
	fmt.Printf("%+v\n", res)
	// nums, numErr := json.Marshal(res)
	// if numErr != nil {
	// } else {
	//
	// 	fmt.Printf("%+v\n", nums)
	// }
	// if !c.specialMessage.IsDirectMessage() {
	// 	prefix = "<@" + c.specialMessage.User() + ">: "
	// }

	// msg := c.specialMessage
	// msg.SetText(prefix + fmt.Sprintf(text, a...))
	//
	// c.send(msg)
}

// String return string paramter
func (c Conversation) String(name string) (string, error) {
	return c.match.String(name)
}

// Integer returns integer parameter
func (c Conversation) Integer(name string) (int, error) {
	return c.match.Integer(name)
}

// Match returns the parameter at the position
func (c Conversation) Match(position int) (string, error) {
	return c.match.Match(position)
}

// NewConversation returns a Conversation struct
func NewConversation(match allot.MatchInterface, msg Message, socket *websocket.Conn) ConversationInterface {
	conv := &Conversation{
		message: msg,
		match:   match,
		socket:  socket,
	}
	// fmt.Printf("%+v\n", conv)
	conv.SetConnection(websocket.JSON)

	return conv
}
