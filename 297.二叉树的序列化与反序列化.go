/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const SEP = ","
const NULL = "#"

type Codec struct {
	List []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	this.serializeList(root)
	return strings.Join(this.List, SEP)
}

func (this *Codec) serializeList(root *TreeNode) {
	if root == nil {
		this.List = append(this.List, NULL)
		return
	}
	this.List = append(this.List, strconv.Itoa(root.Val))
	this.serializeList(root.Left)
	this.serializeList(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.List = []string{}
	for _, v := range strings.Split(data, SEP) {
		this.List = append(this.List, v)
	}
	return this.deserializeList()
}

func (this *Codec) deserializeList() *TreeNode {
	if len(this.List) == 0 {
		return nil
	}

	first := this.List[0]
	this.List = this.List[1:]
	if first == NULL {
		return nil
	}
	root := new(TreeNode)
	root.Val, _ = strconv.Atoi(first)
	root.Left = this.deserializeList()
	root.Right = this.deserializeList()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

