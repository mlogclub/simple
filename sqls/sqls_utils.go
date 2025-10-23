package sqls

import (
	"database/sql"
	"strings"

	"github.com/mlogclub/simple/common/strs"
)

func SqlNullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  len(value) > 0,
	}
}

func KeywordWrap(keyword string) string {
	if strs.IsBlank(keyword) {
		return keyword
	}
	// If already quoted, return as-is
	if (strings.HasPrefix(keyword, "`") && strings.HasSuffix(keyword, "`")) ||
		(strings.HasPrefix(keyword, "\"") && strings.HasSuffix(keyword, "\"")) {
		return keyword
	}

	// Detect current DB dialect via GORM and choose quote style
	dialect := ""
	if db := DB(); db != nil && db.Dialector != nil {
		dialect = db.Dialector.Name()
	}

	quote := "`" // default to MySQL-style backticks to keep prior behavior
	if strings.Contains(strings.ToLower(dialect), "postgre") {
		quote = "\""
	}

	// If identifier contains dot, quote each part separately (e.g., schema.table or table.column)
	if strings.Contains(keyword, ".") {
		parts := strings.Split(keyword, ".")
		for i, p := range parts {
			if p == "*" || p == "" {
				// don't quote wildcard or empty part
				continue
			}
			// avoid double quoting if a part is already quoted
			if (strings.HasPrefix(p, quote) && strings.HasSuffix(p, quote)) ||
				(strings.HasPrefix(p, "`") && strings.HasSuffix(p, "`")) ||
				(strings.HasPrefix(p, "\"") && strings.HasSuffix(p, "\"")) {
				continue
			}
			parts[i] = quote + p + quote
		}
		return strings.Join(parts, ".")
	}

	return quote + keyword + quote
}
