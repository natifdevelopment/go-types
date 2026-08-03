package types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Cursor represents a pagination cursor for cursor-based pagination.
// It encodes the sort field value and ID of the last item from the
// previous page, allowing the next page to be fetched efficiently
// without counting all preceding rows.
//
// The cursor is base64-encoded JSON so it's opaque to the client
// but can be decoded server-side.
type Cursor struct {
	// SortValue is the value of the sort field (e.g. created_at timestamp)
	// from the last item of the previous page.
	SortValue string `json:"sv"`
	// ID is the primary key of the last item, used as a tiebreaker
	// when multiple items have the same sort value.
	ID string `json:"id"`
	// SortField is the field name used for sorting (e.g. "created_at").
	SortField string `json:"sf"`
	// SortDir is the sort direction: "asc" or "desc".
	SortDir string `json:"sd"`
}

// Encode serializes the cursor to a base64 string safe for use in URLs.
func (c Cursor) Encode() (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("cursor marshal: %w", err)
	}
	return base64.URLEncoding.EncodeToString(data), nil
}

// DecodeCursor decodes a base64-encoded cursor string.
func DecodeCursor(s string) (*Cursor, error) {
	if s == "" {
		return nil, nil
	}
	data, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("cursor base64 decode: %w", err)
	}
	var c Cursor
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("cursor unmarshal: %w", err)
	}
	return &c, nil
}

// CursorMeta is the pagination metadata returned to the client when
// using cursor-based pagination. It replaces the traditional
// page/totalPage meta with a nextCursor that the client passes back
// to fetch the next page.
type CursorMeta struct {
	PerPage    int    `json:"perPage"`
	HasMore    bool   `json:"hasMore"`
	NextCursor string `json:"nextCursor,omitempty"`
	SortField  string `json:"sortField"`
	SortDir    string `json:"sortDir"`
}

// BuildCursorFromRow creates a cursor from the last row in a result set.
// The row must have an "ID" field (uuid as string) and a sort field
// (e.g. "created_at" as time.Time).
//
// Usage:
//
//	cursor, _ := types.BuildCursorFromRow(lastRow, "created_at", "desc")
//	nextCursorStr, _ := cursor.Encode()
func BuildCursorFromRow(id string, sortValue string, sortField string, sortDir string) (string, error) {
	c := Cursor{
		SortValue: sortValue,
		ID:        id,
		SortField: sortField,
		SortDir:   sortDir,
	}
	return c.Encode()
}

// ParseSortDir normalizes sort direction to "asc" or "desc".
func ParseSortDir(dir string) string {
	dir = strings.ToLower(strings.TrimSpace(dir))
	if dir != "asc" && dir != "desc" {
		dir = "desc"
	}
	return dir
}

// FormatTimeForCursor converts a time.Time to a cursor-compatible string.
func FormatTimeForCursor(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}

// ParseTimeFromCursor converts a cursor string back to time.Time.
func ParseTimeFromCursor(s string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, s)
}
