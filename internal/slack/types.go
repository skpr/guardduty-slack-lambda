package slack

type Message struct {
	Blocks []Block `json:"blocks"`
}

type BlockType string

const (
	BlockTypeContext BlockType = "context"
	BlockTypeSection BlockType = "section"
	BlockTypeHeader  BlockType = "header"
)

type Block struct {
	Type     BlockType      `json:"type"`
	Elements []BlockElement `json:"elements,omitempty"`
	Text     *BlockText     `json:"text,omitempty"`
}

type BlockElementType string

const (
	// BlockElementTypeMarkdown used to declare a block as markdown.
	BlockElementTypeMarkdown BlockElementType = "mrkdwn"
)

type BlockElement struct {
	Type BlockElementType `json:"type"`
	Text *string          `json:"text,omitempty"`
}

type BlockTextType string

const (
	// BlockTextTypeMarkdown used to declare a block as markdown.
	BlockTextTypeMarkdown BlockTextType = "mrkdwn"
	// BlockTextTypePlainText used to declare a block as plain text.
	BlockTextTypePlainText BlockTextType = "plain_text"
)

type BlockText struct {
	Type BlockTextType `json:"type"`
	Text string        `json:"text"`
}
