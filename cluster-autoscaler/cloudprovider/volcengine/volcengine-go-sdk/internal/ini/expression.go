/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ini

// Copy from https://github.com/aws/aws-sdk-go
// May have been modified by Beijing Volcanoengine Technology Ltd.

// newExpression will return an expression AST.
// Expr represents an expression
//
//	grammar:
//	expr -> string | number
func newExpression(tok Token) AST {
	return newASTWithRootToken(ASTKindExpr, tok)
}

func newEqualExpr(left AST, tok Token) AST {
	return newASTWithRootToken(ASTKindEqualExpr, tok, left)
}

// EqualExprKey will return a LHS value in the equal expr
func EqualExprKey(ast AST) string {
	children := ast.GetChildren()
	if len(children) == 0 || ast.Kind != ASTKindEqualExpr {
		return ""
	}

	return string(children[0].Root.Raw())
}
