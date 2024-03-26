package query_builder

import (
	"strconv"
	"strings"
)

type Builder interface {
	Select(columns ...string) Builder
	From(table string) Builder
	Where(condition string, args ...interface{}) Builder
	Join(joinType, table, condition string, args ...interface{}) Builder
	OrderBy(columns ...string) Builder
	GroupBy(columns ...string) Builder
	Limit(limit int) Builder
	Build() (string, []interface{})
}

type SqlBuilder struct {
	selectColumns []string
	fromTable     string
	whereClause   string
	whereArgs     []interface{}
	joinClause    string
	joinArgs      []interface{}
	orderByClause string
	groupByClause string
	limitCount    int
}

func NewSqlBuilder() Builder {
	return &SqlBuilder{}
}

func (b *SqlBuilder) Join(joinType, table, condition string, args ...interface{}) Builder {
	b.joinClause = joinType + " JOIN " + table + " ON " + condition
	b.joinArgs = append(b.joinArgs, args...)
	return b
}

func (b *SqlBuilder) GroupBy(columns ...string) Builder {
	b.groupByClause = strings.Join(columns, ", ")
	return b
}

func (b *SqlBuilder) OrderBy(columns ...string) Builder {
	b.orderByClause = strings.Join(columns, ", ")
	return b
}

func (b *SqlBuilder) Where(condition string, args ...interface{}) Builder {
	b.whereClause = condition
	b.whereArgs = append(b.whereArgs, args...)
	return b
}

func (b *SqlBuilder) From(table string) Builder {
	b.fromTable = table
	return b
}

func (b *SqlBuilder) Limit(limit int) Builder {
	b.limitCount = limit
	return b
}

func (b *SqlBuilder) Select(columns ...string) Builder {
	if len(columns) > 0 {
		b.selectColumns = append(b.selectColumns, columns...)
	}
	return b
}

func (b *SqlBuilder) Build() (string, []interface{}) {
	query := ""
	args := []interface{}{}

	// Construct the SELECT clause
	query += "SELECT "
	if len(b.selectColumns) > 0 {
		query += strings.Join(b.selectColumns, ", ")
	} else {
		query += "*"
	}
	query = strings.TrimSuffix(query, " ,")
	// Construct the FROM clause
	query += " FROM " + b.fromTable

	// Construct the WHERE clause
	if b.whereClause != "" {
		query += " WHERE " + b.whereClause
		args = append(args, b.whereArgs...)
	}

	// Construct the JOIN clause
	if b.joinClause != "" {
		query += " " + b.joinClause
		args = append(args, b.joinArgs...)
	}

	// Construct the ORDER BY clause
	if b.orderByClause != "" {
		query += " ORDER BY " + b.orderByClause
	}

	// Construct the GROUP BY clause
	if b.groupByClause != "" {
		query += " GROUP BY " + b.groupByClause
	}

	// Construct the LIMIT clause
	if b.limitCount > 0 {
		query += " LIMIT " + strconv.Itoa(b.limitCount)
	}

	return query, args
}
