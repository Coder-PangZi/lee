package types

type ASTree interface {
	Child(i int) ASTree
	ChildrenCnt() int
	Children() []ASTree
	Type() ASTreeType
}

type ASTBase struct {
	asType   byte
	children []ASTree
	token    *Token
}

func (ast *ASTBase) Token() *Token {
	return ast.token
}

func (ast *ASTBase) SetToken(token *Token) {
	ast.token = token
}

func (ast *ASTBase) Child(i int) ASTree {
	return ast.children[i]
}

func (ast *ASTBase) Children() []ASTree {
	return ast.children
}

func (ast *ASTBase) ChildrenCnt() int {
	return len(ast.children)
}

func (ast *ASTBase) Type() byte {
	return ast.asType
}

func (ast *ASTBase) SetType(asType byte) {
	ast.asType = asType
}
