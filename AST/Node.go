package AST

type NodeType string

const (
	NodeTypeBlock NodeType = "block"
	NodeTypeText  NodeType = "text"
)

type Node struct {
	NodeType NodeType `json:"type"`

	// BlockNode
	BlockType string                 `json:"block_type,omitempty"`
	Children  *[]Node                `json:"children,omitempty"`
	Arguments map[string]interface{} `json:"arguments,omitempty"`

	// BlockText
	Value string `json:"value,omitempty"`
}
