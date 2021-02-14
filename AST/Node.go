package AST

type NodeType string

const (
	NodeTypeBlock NodeType = "block"
	NodeTypeText  NodeType = "text"
	NodeTypeList  NodeType = "list"
)

type Node struct {
	NodeType NodeType `json:"type"`

	// BlockNode
	BlockType string                  `json:"block_type,omitempty"`
	Arguments *map[string]interface{} `json:"arguments,omitempty"`
	Children  *[]Node                 `json:"children,omitempty"`

	// BlockText
	Value string `json:"value,omitempty"`
}

func FromText(text string) *Node {
	return &Node{
		NodeType: NodeTypeText,
		Value:    text,
	}
}

func FromBlockType(blockType string, arguments *map[string]interface{}, children *[]Node) *Node {
	return &Node{
		NodeType:  NodeTypeBlock,
		BlockType: blockType,
		Arguments: arguments,
		Children:  children,
	}
}
