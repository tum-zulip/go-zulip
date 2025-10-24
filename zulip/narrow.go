package zulip

import (
	"encoding/json"
	"fmt"

	"github.com/tum-zulip/go-zulip/zulip/internal/utils"
)

// Narrow represents a query constraint for filtering messages in the Zulip API.
// It consists of one or more NarrowTerms that can be combined to create complex queries.
// Use NewNarrow() to create a new Narrow and chain Where() or WhereNot() to add constraints.
type Narrow struct {
	terms []NarrowTerm
}

// NarrowTerm represents a single constraint within a Narrow query.
// It specifies an operator, an operand, and whether the term should be negated.
type NarrowTerm struct {
	Operator NarrowOperator `json:"operator"`
	Operand  NarrowOperand  `json:"operand"`
	Negated  bool           `json:"negated,omitempty"`
}

// NarrowOperand represents the value part of a narrow constraint.
// It can hold one of multiple types: string, int64, list of ints, or list of strings.
type NarrowOperand struct {
	String    *string
	Int       *int64
	ListOfInt *[]int64
}

// NarrowOperator represents a filter operator used in narrow queries.
type NarrowOperator string

// NewNarrow creates and returns a new empty Narrow query builder.
func NewNarrow() *Narrow {
	return &Narrow{}
}

// And appends a positive narrow term to the query.
func (n *Narrow) And(term NarrowTerm) *Narrow {
	n.terms = append(n.terms, term)
	return n
}

// AndNot adds and negates a positive term to the query
func (n *Narrow) AndNot(term NarrowTerm) *Narrow {
	term.Negated = true
	return n.And(term)
}

// Constructs a narrow with a single term
func Where(term NarrowTerm) *Narrow {
	return NewNarrow().And(term)
}

// ChannelNameIs returns a NarrowTerm that filters messages by channel name.
func ChannelNameIs(name string) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorChannel,
		Operand:  NewNarrowStringOperand(name),
	}
}

// ChannelIs returns a NarrowTerm that filters messages by channel ID.
func ChannelIs(channelId int64) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorChannel,
		Operand:  NewNarrowIntOperand(channelId),
	}
}

// TopicIs returns a NarrowTerm that filters messages by topic name.
func TopicIs(topic string) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorTopic,
		Operand:  NewNarrowStringOperand(topic),
	}
}

// SenderEmailIs returns a NarrowTerm that filters messages sent by a specific email address.
func SenderEmailIs(email string) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorSender,
		Operand:  NewNarrowStringOperand(email),
	}
}

// SenderIs returns a NarrowTerm that filters messages sent by a specific user ID.
func SenderIs(userId int64) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorSender,
		Operand:  NewNarrowIntOperand(userId),
	}
}

// HasReactions returns a NarrowTerm that filters messages that have reactions.
func HasReactions() NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorHas,
		Operand:  NewNarrowStringOperand("reactions"), // TODO: make this an enum
	}
}

// IsMuted returns a NarrowTerm that filters muted messages.
func IsMuted() NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorIs,
		Operand:  NewNarrowStringOperand("muted"),
	}
}

// IsDirectMessage returns a NarrowTerm that filters private messages.
func IsDirectMessage() NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorIs,
		Operand:  NewNarrowStringOperand("private"),
	}
}

// IsStarred returns a NarrowTerm that filters starred messages.
func IsStarred() NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorIs,
		Operand:  NewNarrowStringOperand("starred"),
	}
}

// IsAlerted returns a NarrowTerm that filters alerted messages.
func IsAlerted() NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorIs,
		Operand:  NewNarrowStringOperand("alerted"),
	}
}

// MessageContains returns a NarrowTerm that filters messages using full-text search.
func MessageContains(search string) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorSearch,
		Operand:  NewNarrowStringOperand(search),
	}
}

// StreamIs returns a NarrowTerm that filters messages by stream name (legacy alias for ChannelNameIs).
func StreamIs(name string) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorStream,
		Operand:  NewNarrowStringOperand(name),
	}
}

// MessageID returns a NarrowTerm that filters a specific message by its ID.
func MessageID(id int64) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorId,
		Operand:  NewNarrowIntOperand(id),
	}
}

// DirectMessage returns a NarrowTerm that filters direct messages between two specific user IDs.
func DirectMessage(userIds []int64) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorDm,
		Operand:  NewNarrowListOfIntOperand(userIds),
	}
}

// DirectMessageIncluding returns a NarrowTerm that filters group direct messages including specific user IDs.
func DirectMessageIncluding(userIds []int64) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorDmIncluding,
		Operand:  NewNarrowListOfIntOperand(userIds),
	}
}

// WithUser returns a NarrowTerm that filters direct messages with a specific user ID.
func WithUser(userId int64) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorWith,
		Operand:  NewNarrowIntOperand(userId),
	}
}

// WithUserEmail returns a NarrowTerm that filters direct messages with a specific user email.
func WithUserEmail(email string) NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorWith,
		Operand:  NewNarrowStringOperand(email),
	}
}

// HasFiles returns a NarrowTerm that filters messages that have files attached.
func HasFiles() NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorHas,
		Operand:  NewNarrowStringOperand("files"),
	}
}

// HasImages returns a NarrowTerm that filters messages that have images embedded.
func HasImages() NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorHas,
		Operand:  NewNarrowStringOperand("images"),
	}
}

// HasLinks returns a NarrowTerm that filters messages that have links.
func HasLinks() NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorHas,
		Operand:  NewNarrowStringOperand("links"),
	}
}

// HasEmbeds returns a NarrowTerm that filters messages that have embedded content.
func HasEmbeds() NarrowTerm {
	return NarrowTerm{
		Operator: NarrowOperatorHas,
		Operand:  NewNarrowStringOperand("embeds"),
	}
}

func (o NarrowOperand) MarshalJSON() ([]byte, error) {
	return utils.MarshalUnionType(o)
}

func (o *NarrowOperand) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalUnionType(data, o)
}

// NewNarrowStringOperand creates a NarrowOperand with a string value.
func NewNarrowStringOperand(value string) NarrowOperand {
	return NarrowOperand{
		String: &value,
	}
}

// NewNarrowIntOperand creates a NarrowOperand with an int64 value.
func NewNarrowIntOperand(value int64) NarrowOperand {
	return NarrowOperand{
		Int: &value,
	}
}

// NewNarrowListOfIntOperand creates a NarrowOperand with a list of int64 values.
func NewNarrowListOfIntOperand(value []int64) NarrowOperand {
	return NarrowOperand{
		ListOfInt: &value,
	}
}

func (o Narrow) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.terms)
}

func (o *Narrow) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &o.terms)
}

// TODO: make not exported
func (n *Narrow) ToArray() ([][]string, error) {
	if n == nil {
		return nil, nil
	}
	terms := make([][]string, len(n.terms))
	for i, t := range n.terms {
		if t.Negated {
			return nil, fmt.Errorf("negated narrows not supported with this endpoint")
		}
		term := make([]string, 2)
		term[0] = string(t.Operator)

		if t.Operand.ListOfInt != nil {
			return nil, fmt.Errorf("negated narrows not supported with this endpoint")
		}
		if t.Operand.Int != nil && t.Operand.String != nil {
			return nil, fmt.Errorf("only one type of operand can be valid at a time")
		}

		if t.Operand.String != nil {
			term[1] = *t.Operand.String
		} else if t.Operand.Int != nil {
			term[1] = fmt.Sprintf("%d", *t.Operand.Int)
		} else {
			return nil, fmt.Errorf("operand cannot me empty")
		}
		terms[i] = term

	}
	return terms, nil
}
