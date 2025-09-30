package main

import (
	"fmt"
	"strings"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_javascript "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
)

func main() {
	code := []byte("const foo = 1 + 2")

	parser := tree_sitter.NewParser()
	defer parser.Close()
	parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_javascript.Language()))

	tree := parser.Parse(code, nil)
	defer tree.Close()

	root := tree.RootNode()

	fmt.Println("🌳 Tree-Sitter Analysis")
	fmt.Println("=" + strings.Repeat("=", 40))
	fmt.Printf("📝 Source Code: %s\n", string(code))
	fmt.Printf("🔍 Language: JavaScript\n")
	fmt.Printf("📊 Node Count: %d\n", root.NamedChildCount())
	fmt.Println()

	fmt.Println("🔧 Raw S-Expression:")
	fmt.Println(root.ToSexp())
	fmt.Println()

	fmt.Println("🌲 Pretty Tree Structure:")
	printTree(root, "", true, string(code))
	fmt.Println()

	fmt.Println("📋 Node Details:")
	printNodeDetails(root, string(code), 0)
}

func printTree(node *tree_sitter.Node, prefix string, isLast bool, source string) {
	nodeType := node.Kind()
	text := source[node.StartByte():node.EndByte()]
	//
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	if node.IsNamed() {
		fmt.Printf("%s%s%s", prefix, connector, nodeType)
		if len(strings.TrimSpace(text)) > 0 && len(text) < 50 {
			fmt.Printf(" 📄 \"%s\"", strings.ReplaceAll(text, "\n", "\\n"))
		}
		fmt.Printf(" [%d:%d-%d:%d]\n",
			node.StartPosition().Row+1, node.StartPosition().Column+1,
			node.EndPosition().Row+1, node.EndPosition().Column+1)
	}

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	childCount := int(node.NamedChildCount())
	for i := 0; i < childCount; i++ {
		child := node.NamedChild(uint(i))
		isLastChild := i == childCount-1
		printTree(child, childPrefix, isLastChild, source)
	}
}

func printNodeDetails(node *tree_sitter.Node, source string, depth int) {
	indent := strings.Repeat("  ", depth)
	nodeType := node.Kind()
	text := source[node.StartByte():node.EndByte()]

	if node.IsNamed() {
		fmt.Printf("%s🔸 %s\n", indent, nodeType)
		fmt.Printf("%s   📍 Position: %d:%d to %d:%d\n", indent,
			node.StartPosition().Row+1, node.StartPosition().Column+1,
			node.EndPosition().Row+1, node.EndPosition().Column+1)
		fmt.Printf("%s   📏 Bytes: %d-%d\n", indent, node.StartByte(), node.EndByte())

		if len(strings.TrimSpace(text)) > 0 {
			displayText := text
			if len(text) > 100 {
				displayText = text[:97] + "..."
			}
			fmt.Printf("%s   📝 Text: \"%s\"\n", indent, strings.ReplaceAll(displayText, "\n", "\\n"))
		}

		if node.HasError() {
			fmt.Printf("%s   ❌ Has Error: true\n", indent)
		}

		if node.IsMissing() {
			fmt.Printf("%s   ⚠️  Is Missing: true\n", indent)
		}

		fmt.Println()
	}

	childCount := int(node.NamedChildCount())
	for i := 0; i < childCount; i++ {
		child := node.NamedChild(uint(i))
		printNodeDetails(child, source, depth+1)
	}
}
