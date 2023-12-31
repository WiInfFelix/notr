// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/WiInfFelix/notr/ent/book"
	"github.com/WiInfFelix/notr/ent/chapter"
	"github.com/WiInfFelix/notr/ent/predicate"
)

// ChapterUpdate is the builder for updating Chapter entities.
type ChapterUpdate struct {
	config
	hooks    []Hook
	mutation *ChapterMutation
}

// Where appends a list predicates to the ChapterUpdate builder.
func (cu *ChapterUpdate) Where(ps ...predicate.Chapter) *ChapterUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetTitle sets the "title" field.
func (cu *ChapterUpdate) SetTitle(s string) *ChapterUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cu *ChapterUpdate) SetNillableTitle(s *string) *ChapterUpdate {
	if s != nil {
		cu.SetTitle(*s)
	}
	return cu
}

// SetDescription sets the "description" field.
func (cu *ChapterUpdate) SetDescription(s string) *ChapterUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *ChapterUpdate) SetNillableDescription(s *string) *ChapterUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// AddBookIDs adds the "book" edge to the Book entity by IDs.
func (cu *ChapterUpdate) AddBookIDs(ids ...int) *ChapterUpdate {
	cu.mutation.AddBookIDs(ids...)
	return cu
}

// AddBook adds the "book" edges to the Book entity.
func (cu *ChapterUpdate) AddBook(b ...*Book) *ChapterUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.AddBookIDs(ids...)
}

// Mutation returns the ChapterMutation object of the builder.
func (cu *ChapterUpdate) Mutation() *ChapterMutation {
	return cu.mutation
}

// ClearBook clears all "book" edges to the Book entity.
func (cu *ChapterUpdate) ClearBook() *ChapterUpdate {
	cu.mutation.ClearBook()
	return cu
}

// RemoveBookIDs removes the "book" edge to Book entities by IDs.
func (cu *ChapterUpdate) RemoveBookIDs(ids ...int) *ChapterUpdate {
	cu.mutation.RemoveBookIDs(ids...)
	return cu
}

// RemoveBook removes "book" edges to Book entities.
func (cu *ChapterUpdate) RemoveBook(b ...*Book) *ChapterUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.RemoveBookIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChapterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChapterUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChapterUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChapterUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ChapterUpdate) check() error {
	if v, ok := cu.mutation.Title(); ok {
		if err := chapter.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Chapter.title": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Description(); ok {
		if err := chapter.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Chapter.description": %w`, err)}
		}
	}
	return nil
}

func (cu *ChapterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(chapter.Table, chapter.Columns, sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.SetField(chapter.FieldTitle, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(chapter.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.BookCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chapter.BookTable,
			Columns: []string{chapter.BookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedBookIDs(); len(nodes) > 0 && !cu.mutation.BookCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chapter.BookTable,
			Columns: []string{chapter.BookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BookIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chapter.BookTable,
			Columns: []string{chapter.BookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chapter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChapterUpdateOne is the builder for updating a single Chapter entity.
type ChapterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChapterMutation
}

// SetTitle sets the "title" field.
func (cuo *ChapterUpdateOne) SetTitle(s string) *ChapterUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cuo *ChapterUpdateOne) SetNillableTitle(s *string) *ChapterUpdateOne {
	if s != nil {
		cuo.SetTitle(*s)
	}
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ChapterUpdateOne) SetDescription(s string) *ChapterUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *ChapterUpdateOne) SetNillableDescription(s *string) *ChapterUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// AddBookIDs adds the "book" edge to the Book entity by IDs.
func (cuo *ChapterUpdateOne) AddBookIDs(ids ...int) *ChapterUpdateOne {
	cuo.mutation.AddBookIDs(ids...)
	return cuo
}

// AddBook adds the "book" edges to the Book entity.
func (cuo *ChapterUpdateOne) AddBook(b ...*Book) *ChapterUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.AddBookIDs(ids...)
}

// Mutation returns the ChapterMutation object of the builder.
func (cuo *ChapterUpdateOne) Mutation() *ChapterMutation {
	return cuo.mutation
}

// ClearBook clears all "book" edges to the Book entity.
func (cuo *ChapterUpdateOne) ClearBook() *ChapterUpdateOne {
	cuo.mutation.ClearBook()
	return cuo
}

// RemoveBookIDs removes the "book" edge to Book entities by IDs.
func (cuo *ChapterUpdateOne) RemoveBookIDs(ids ...int) *ChapterUpdateOne {
	cuo.mutation.RemoveBookIDs(ids...)
	return cuo
}

// RemoveBook removes "book" edges to Book entities.
func (cuo *ChapterUpdateOne) RemoveBook(b ...*Book) *ChapterUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.RemoveBookIDs(ids...)
}

// Where appends a list predicates to the ChapterUpdate builder.
func (cuo *ChapterUpdateOne) Where(ps ...predicate.Chapter) *ChapterUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChapterUpdateOne) Select(field string, fields ...string) *ChapterUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chapter entity.
func (cuo *ChapterUpdateOne) Save(ctx context.Context) (*Chapter, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChapterUpdateOne) SaveX(ctx context.Context) *Chapter {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChapterUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChapterUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ChapterUpdateOne) check() error {
	if v, ok := cuo.mutation.Title(); ok {
		if err := chapter.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Chapter.title": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Description(); ok {
		if err := chapter.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Chapter.description": %w`, err)}
		}
	}
	return nil
}

func (cuo *ChapterUpdateOne) sqlSave(ctx context.Context) (_node *Chapter, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(chapter.Table, chapter.Columns, sqlgraph.NewFieldSpec(chapter.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Chapter.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chapter.FieldID)
		for _, f := range fields {
			if !chapter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chapter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.SetField(chapter.FieldTitle, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(chapter.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.BookCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chapter.BookTable,
			Columns: []string{chapter.BookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedBookIDs(); len(nodes) > 0 && !cuo.mutation.BookCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chapter.BookTable,
			Columns: []string{chapter.BookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BookIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chapter.BookTable,
			Columns: []string{chapter.BookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Chapter{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chapter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
