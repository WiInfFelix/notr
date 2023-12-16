// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/WiInfFelix/notr/ent/todo"
)

// Todo is the model entity for the Todo schema.
type Todo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TodoQuery when eager-loading is set.
	Edges        TodoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TodoEdges holds the relations/edges for other nodes in the graph.
type TodoEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// Book holds the value of the book edge.
	Book []*Book `json:"book,omitempty"`
	// Chapter holds the value of the chapter edge.
	Chapter []*Chapter `json:"chapter,omitempty"`
	// Page holds the value of the page edge.
	Page []*Page `json:"page,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e TodoEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// BookOrErr returns the Book value or an error if the edge
// was not loaded in eager-loading.
func (e TodoEdges) BookOrErr() ([]*Book, error) {
	if e.loadedTypes[1] {
		return e.Book, nil
	}
	return nil, &NotLoadedError{edge: "book"}
}

// ChapterOrErr returns the Chapter value or an error if the edge
// was not loaded in eager-loading.
func (e TodoEdges) ChapterOrErr() ([]*Chapter, error) {
	if e.loadedTypes[2] {
		return e.Chapter, nil
	}
	return nil, &NotLoadedError{edge: "chapter"}
}

// PageOrErr returns the Page value or an error if the edge
// was not loaded in eager-loading.
func (e TodoEdges) PageOrErr() ([]*Page, error) {
	if e.loadedTypes[3] {
		return e.Page, nil
	}
	return nil, &NotLoadedError{edge: "page"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Todo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case todo.FieldStatus:
			values[i] = new(sql.NullBool)
		case todo.FieldID:
			values[i] = new(sql.NullInt64)
		case todo.FieldTitle, todo.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Todo fields.
func (t *Todo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case todo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case todo.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case todo.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case todo.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = value.Bool
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Todo.
// This includes values selected through modifiers, order, etc.
func (t *Todo) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Todo entity.
func (t *Todo) QueryUser() *UserQuery {
	return NewTodoClient(t.config).QueryUser(t)
}

// QueryBook queries the "book" edge of the Todo entity.
func (t *Todo) QueryBook() *BookQuery {
	return NewTodoClient(t.config).QueryBook(t)
}

// QueryChapter queries the "chapter" edge of the Todo entity.
func (t *Todo) QueryChapter() *ChapterQuery {
	return NewTodoClient(t.config).QueryChapter(t)
}

// QueryPage queries the "page" edge of the Todo entity.
func (t *Todo) QueryPage() *PageQuery {
	return NewTodoClient(t.config).QueryPage(t)
}

// Update returns a builder for updating this Todo.
// Note that you need to call Todo.Unwrap() before calling this method if this Todo
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Todo) Update() *TodoUpdateOne {
	return NewTodoClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Todo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Todo) Unwrap() *Todo {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Todo is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Todo) String() string {
	var builder strings.Builder
	builder.WriteString("Todo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("title=")
	builder.WriteString(t.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Todos is a parsable slice of Todo.
type Todos []*Todo
