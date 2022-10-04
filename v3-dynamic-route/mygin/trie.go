package mygin

import "strings"

type node struct {
	pattern  string  //待匹配的路由 例如 /p/:la
	part     string  //路由中的一部分，例如 :la
	children []*node //子节点
	isWild   bool    //是否精确匹配
}

// 匹配子节点，返回第一个匹配成功的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		// 1.如果子节点的part==传入的part 或者 子节点是模糊匹配例如 :la 即为匹配成功
		// 2.同时成立
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 遍历下一层
// 找到所有匹配成功的节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 构建路由树
func (n *node) insert(pattern string, parts []string, height int) {
	//终止条件：height匹配完，到了最下层
	if len(parts) == height {
		//匹配完给node的pattern赋值
		n.pattern = pattern
		return
	}

	part := parts[height]
	// 匹配出一个子节点
	child := n.matchChild(part)

	// 没匹配到节点
	if child == nil {
		// 如果当前part的第一个字符是":"或者"*"就为模糊匹配
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		// 增加当前节点的子节点
		n.children = append(n.children, child)
	}

	// 去下一层匹配判断是否要插入节点
	child.insert(pattern, parts, height+1)
}

// 递归搜索匹配的node
func (n *node) search(parts []string, height int) *node {
	// 如果匹配完成 或者 匹配到了 *
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 节点的pattern值为空
		if n.pattern == "" {
			// 匹配失败
			return nil
		}
		return n
	}

	part := parts[height]
	// 找到匹配的和part匹配的子节点
	children := n.matchChildren(part)

	// 遍历找到的所有子节点
	for _, child := range children {
		// 递归搜索
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
