// generated by gocc; DO NOT EDIT.

package parser

import "github.com/awalterschulze/gographviz/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : DotGraph	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `DotGraph : graphx "{" "}"	<< ast.NewGraph(ast.GRAPH, ast.FALSE, nil, nil) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      1,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.GRAPH, ast.FALSE, nil, nil)
		},
	},
	ProdTabEntry{
		String: `DotGraph : strict graphx "{" "}"	<< ast.NewGraph(ast.GRAPH, ast.TRUE, nil, nil) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      2,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.GRAPH, ast.TRUE, nil, nil)
		},
	},
	ProdTabEntry{
		String: `DotGraph : graphx Id "{" "}"	<< ast.NewGraph(ast.GRAPH, ast.FALSE, X[1], nil) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      3,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.GRAPH, ast.FALSE, X[1], nil)
		},
	},
	ProdTabEntry{
		String: `DotGraph : strict graphx Id "{" "}"	<< ast.NewGraph(ast.GRAPH, ast.TRUE, X[2], nil) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      4,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.GRAPH, ast.TRUE, X[2], nil)
		},
	},
	ProdTabEntry{
		String: `DotGraph : graphx "{" StmtList "}"	<< ast.NewGraph(ast.GRAPH, ast.FALSE, nil, X[2]) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      5,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.GRAPH, ast.FALSE, nil, X[2])
		},
	},
	ProdTabEntry{
		String: `DotGraph : graphx Id "{" StmtList "}"	<< ast.NewGraph(ast.GRAPH, ast.FALSE, X[1], X[3]) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      6,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.GRAPH, ast.FALSE, X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `DotGraph : strict graphx "{" StmtList "}"	<< ast.NewGraph(ast.GRAPH, ast.TRUE, nil, X[3]) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      7,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.GRAPH, ast.TRUE, nil, X[3])
		},
	},
	ProdTabEntry{
		String: `DotGraph : strict graphx Id "{" StmtList "}"	<< ast.NewGraph(ast.GRAPH, ast.TRUE, X[2], X[4]) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      8,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.GRAPH, ast.TRUE, X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `DotGraph : digraph "{" "}"	<< ast.NewGraph(ast.DIGRAPH, ast.FALSE, nil, nil) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.DIGRAPH, ast.FALSE, nil, nil)
		},
	},
	ProdTabEntry{
		String: `DotGraph : strict digraph "{" "}"	<< ast.NewGraph(ast.DIGRAPH, ast.TRUE, nil, nil) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      10,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.DIGRAPH, ast.TRUE, nil, nil)
		},
	},
	ProdTabEntry{
		String: `DotGraph : digraph Id "{" "}"	<< ast.NewGraph(ast.DIGRAPH, ast.FALSE, X[1], nil) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      11,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.DIGRAPH, ast.FALSE, X[1], nil)
		},
	},
	ProdTabEntry{
		String: `DotGraph : strict digraph Id "{" "}"	<< ast.NewGraph(ast.DIGRAPH, ast.TRUE, X[2], nil) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      12,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.DIGRAPH, ast.TRUE, X[2], nil)
		},
	},
	ProdTabEntry{
		String: `DotGraph : digraph "{" StmtList "}"	<< ast.NewGraph(ast.DIGRAPH, ast.FALSE, nil, X[2]) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      13,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.DIGRAPH, ast.FALSE, nil, X[2])
		},
	},
	ProdTabEntry{
		String: `DotGraph : digraph Id "{" StmtList "}"	<< ast.NewGraph(ast.DIGRAPH, ast.FALSE, X[1], X[3]) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      14,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.DIGRAPH, ast.FALSE, X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `DotGraph : strict digraph "{" StmtList "}"	<< ast.NewGraph(ast.DIGRAPH, ast.TRUE, nil, X[3]) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      15,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.DIGRAPH, ast.TRUE, nil, X[3])
		},
	},
	ProdTabEntry{
		String: `DotGraph : strict digraph Id "{" StmtList "}"	<< ast.NewGraph(ast.DIGRAPH, ast.TRUE, X[2], X[4]) >>`,
		Id:         "DotGraph",
		NTType:     1,
		Index:      16,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraph(ast.DIGRAPH, ast.TRUE, X[2], X[4])
		},
	},
	ProdTabEntry{
		String: `StmtList : Stmt1	<< ast.NewStmtList(X[0]) >>`,
		Id:         "StmtList",
		NTType:     2,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewStmtList(X[0])
		},
	},
	ProdTabEntry{
		String: `StmtList : StmtList Stmt1	<< ast.AppendStmtList(X[0], X[1]) >>`,
		Id:         "StmtList",
		NTType:     2,
		Index:      18,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendStmtList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Stmt1 : Stmt	<< X[0], nil >>`,
		Id:         "Stmt1",
		NTType:     3,
		Index:      19,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stmt1 : Stmt ";"	<< X[0], nil >>`,
		Id:         "Stmt1",
		NTType:     3,
		Index:      20,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stmt : Id "=" Id	<< ast.NewAttr(X[0], X[2]) >>`,
		Id:         "Stmt",
		NTType:     4,
		Index:      21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAttr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Stmt : NodeStmt	<< X[0], nil >>`,
		Id:         "Stmt",
		NTType:     4,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stmt : EdgeStmt	<< X[0], nil >>`,
		Id:         "Stmt",
		NTType:     4,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stmt : AttrStmt	<< X[0], nil >>`,
		Id:         "Stmt",
		NTType:     4,
		Index:      24,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Stmt : SubGraphStmt	<< X[0], nil >>`,
		Id:         "Stmt",
		NTType:     4,
		Index:      25,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `AttrStmt : graphx AttrList	<< ast.NewGraphAttrs(X[1]) >>`,
		Id:         "AttrStmt",
		NTType:     5,
		Index:      26,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewGraphAttrs(X[1])
		},
	},
	ProdTabEntry{
		String: `AttrStmt : node AttrList	<< ast.NewNodeAttrs(X[1]) >>`,
		Id:         "AttrStmt",
		NTType:     5,
		Index:      27,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNodeAttrs(X[1])
		},
	},
	ProdTabEntry{
		String: `AttrStmt : edge AttrList	<< ast.NewEdgeAttrs(X[1]) >>`,
		Id:         "AttrStmt",
		NTType:     5,
		Index:      28,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEdgeAttrs(X[1])
		},
	},
	ProdTabEntry{
		String: `AttrList : "[" "]"	<< ast.NewAttrList(nil) >>`,
		Id:         "AttrList",
		NTType:     6,
		Index:      29,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAttrList(nil)
		},
	},
	ProdTabEntry{
		String: `AttrList : "[" AList "]"	<< ast.NewAttrList(X[1]) >>`,
		Id:         "AttrList",
		NTType:     6,
		Index:      30,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAttrList(X[1])
		},
	},
	ProdTabEntry{
		String: `AttrList : AttrList "[" "]"	<< ast.AppendAttrList(X[0], nil) >>`,
		Id:         "AttrList",
		NTType:     6,
		Index:      31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendAttrList(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `AttrList : AttrList "[" AList "]"	<< ast.AppendAttrList(X[0], X[2]) >>`,
		Id:         "AttrList",
		NTType:     6,
		Index:      32,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendAttrList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `AList : Attr	<< ast.NewAList(X[0]) >>`,
		Id:         "AList",
		NTType:     7,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAList(X[0])
		},
	},
	ProdTabEntry{
		String: `AList : AList Attr	<< ast.AppendAList(X[0], X[1]) >>`,
		Id:         "AList",
		NTType:     7,
		Index:      34,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendAList(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `AList : AList "," Attr	<< ast.AppendAList(X[0], X[2]) >>`,
		Id:         "AList",
		NTType:     7,
		Index:      35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendAList(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Attr : Id	<< ast.NewAttr(X[0], nil) >>`,
		Id:         "Attr",
		NTType:     8,
		Index:      36,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAttr(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Attr : Id "=" Id	<< ast.NewAttr(X[0], X[2]) >>`,
		Id:         "Attr",
		NTType:     8,
		Index:      37,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewAttr(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `EdgeStmt : NodeId EdgeRHS	<< ast.NewEdgeStmt(X[0], X[1], nil) >>`,
		Id:         "EdgeStmt",
		NTType:     9,
		Index:      38,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEdgeStmt(X[0], X[1], nil)
		},
	},
	ProdTabEntry{
		String: `EdgeStmt : NodeId EdgeRHS AttrList	<< ast.NewEdgeStmt(X[0], X[1], X[2]) >>`,
		Id:         "EdgeStmt",
		NTType:     9,
		Index:      39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEdgeStmt(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `EdgeStmt : SubGraphStmt EdgeRHS	<< ast.NewEdgeStmt(X[0], X[1], nil) >>`,
		Id:         "EdgeStmt",
		NTType:     9,
		Index:      40,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEdgeStmt(X[0], X[1], nil)
		},
	},
	ProdTabEntry{
		String: `EdgeStmt : SubGraphStmt EdgeRHS AttrList	<< ast.NewEdgeStmt(X[0], X[1], X[2]) >>`,
		Id:         "EdgeStmt",
		NTType:     9,
		Index:      41,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEdgeStmt(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `EdgeRHS : EdgeOp NodeId	<< ast.NewEdgeRHS(X[0], X[1]) >>`,
		Id:         "EdgeRHS",
		NTType:     10,
		Index:      42,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEdgeRHS(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `EdgeRHS : EdgeOp SubGraphStmt	<< ast.NewEdgeRHS(X[0], X[1]) >>`,
		Id:         "EdgeRHS",
		NTType:     10,
		Index:      43,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewEdgeRHS(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `EdgeRHS : EdgeRHS EdgeOp NodeId	<< ast.AppendEdgeRHS(X[0], X[1], X[2]) >>`,
		Id:         "EdgeRHS",
		NTType:     10,
		Index:      44,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendEdgeRHS(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `EdgeRHS : EdgeRHS EdgeOp SubGraphStmt	<< ast.AppendEdgeRHS(X[0], X[1], X[2]) >>`,
		Id:         "EdgeRHS",
		NTType:     10,
		Index:      45,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendEdgeRHS(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String: `NodeStmt : NodeId	<< ast.NewNodeStmt(X[0], nil) >>`,
		Id:         "NodeStmt",
		NTType:     11,
		Index:      46,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNodeStmt(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `NodeStmt : NodeId AttrList	<< ast.NewNodeStmt(X[0], X[1]) >>`,
		Id:         "NodeStmt",
		NTType:     11,
		Index:      47,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNodeStmt(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `NodeId : Id	<< ast.NewNodeId(X[0], nil) >>`,
		Id:         "NodeId",
		NTType:     12,
		Index:      48,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNodeId(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `NodeId : Id Port	<< ast.NewNodeId(X[0], X[1]) >>`,
		Id:         "NodeId",
		NTType:     12,
		Index:      49,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewNodeId(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Port : ":" Id	<< ast.NewPort(X[1], nil) >>`,
		Id:         "Port",
		NTType:     13,
		Index:      50,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPort(X[1], nil)
		},
	},
	ProdTabEntry{
		String: `Port : ":" Id ":" Id	<< ast.NewPort(X[1], X[3]) >>`,
		Id:         "Port",
		NTType:     13,
		Index:      51,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPort(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `SubGraphStmt : "{" StmtList "}"	<< ast.NewSubGraph(nil, X[1]) >>`,
		Id:         "SubGraphStmt",
		NTType:     14,
		Index:      52,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSubGraph(nil, X[1])
		},
	},
	ProdTabEntry{
		String: `SubGraphStmt : subgraph "{" StmtList "}"	<< ast.NewSubGraph(nil, X[2]) >>`,
		Id:         "SubGraphStmt",
		NTType:     14,
		Index:      53,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSubGraph(nil, X[2])
		},
	},
	ProdTabEntry{
		String: `SubGraphStmt : subgraph Id "{" StmtList "}"	<< ast.NewSubGraph(X[1], X[3]) >>`,
		Id:         "SubGraphStmt",
		NTType:     14,
		Index:      54,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewSubGraph(X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `EdgeOp : "->"	<< ast.DIRECTED, nil >>`,
		Id:         "EdgeOp",
		NTType:     15,
		Index:      55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.DIRECTED, nil
		},
	},
	ProdTabEntry{
		String: `EdgeOp : "--"	<< ast.UNDIRECTED, nil >>`,
		Id:         "EdgeOp",
		NTType:     15,
		Index:      56,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.UNDIRECTED, nil
		},
	},
	ProdTabEntry{
		String: `Id : id	<< ast.NewId(X[0]) >>`,
		Id:         "Id",
		NTType:     16,
		Index:      57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewId(X[0])
		},
	},
}
