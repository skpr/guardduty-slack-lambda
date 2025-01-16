package slack

// Message which is sent to Slack.
type Message struct {
	Blocks []Block `json:"blocks"`
}

// BlockType for formatting.
type BlockType string

const (
	// BlockTypeContext for formatting.
	BlockTypeContext BlockType = "context"
	// BlockTypeSection for formatting.
	BlockTypeSection BlockType = "section"
	// BlockTypeHeader for formatting.
	BlockTypeHeader BlockType = "header"
)

// Block section in Slack message.
type Block struct {
	Type     BlockType      `json:"type"`
	Elements []BlockElement `json:"elements,omitempty"`
	Text     *BlockText     `json:"text,omitempty"`
}

// BlockElementType for formatting.
type BlockElementType string

const (
	// BlockElementTypeMarkdown used to declare a block as markdown.
	BlockElementTypeMarkdown BlockElementType = "mrkdwn"
)

// BlockElement for context sections.
type BlockElement struct {
	Type BlockElementType `json:"type"`
	Text *string          `json:"text,omitempty"`
}

// BlockTextType for text formatting.
type BlockTextType string

const (
	// BlockTextTypeMarkdown used to declare a block as markdown.
	BlockTextTypeMarkdown BlockTextType = "mrkdwn"
	// BlockTextTypePlainText used to declare a block as plain text.
	BlockTextTypePlainText BlockTextType = "plain_text"
)

// BlockText provides a block section.
type BlockText struct {
	Type BlockTextType `json:"type"`
	Text string        `json:"text"`
}
