// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"mickaelalliel.com/telebot/parser/internal/ent/expense"
	"mickaelalliel.com/telebot/parser/internal/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeExpense = "Expense"
)

// ExpenseMutation represents an operation that mutates the Expense nodes in the graph.
type ExpenseMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	amount        *float64
	addamount     *float64
	category      *string
	method        *string
	ownerName     *string
	timestamp     *time.Time
	createdAt     *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Expense, error)
	predicates    []predicate.Expense
}

var _ ent.Mutation = (*ExpenseMutation)(nil)

// expenseOption allows management of the mutation configuration using functional options.
type expenseOption func(*ExpenseMutation)

// newExpenseMutation creates new mutation for the Expense entity.
func newExpenseMutation(c config, op Op, opts ...expenseOption) *ExpenseMutation {
	m := &ExpenseMutation{
		config:        c,
		op:            op,
		typ:           TypeExpense,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withExpenseID sets the ID field of the mutation.
func withExpenseID(id int64) expenseOption {
	return func(m *ExpenseMutation) {
		var (
			err   error
			once  sync.Once
			value *Expense
		)
		m.oldValue = func(ctx context.Context) (*Expense, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Expense.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withExpense sets the old Expense of the mutation.
func withExpense(node *Expense) expenseOption {
	return func(m *ExpenseMutation) {
		m.oldValue = func(context.Context) (*Expense, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ExpenseMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ExpenseMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ExpenseMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ExpenseMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Expense.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAmount sets the "amount" field.
func (m *ExpenseMutation) SetAmount(f float64) {
	m.amount = &f
	m.addamount = nil
}

// Amount returns the value of the "amount" field in the mutation.
func (m *ExpenseMutation) Amount() (r float64, exists bool) {
	v := m.amount
	if v == nil {
		return
	}
	return *v, true
}

// OldAmount returns the old "amount" field's value of the Expense entity.
// If the Expense object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExpenseMutation) OldAmount(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAmount: %w", err)
	}
	return oldValue.Amount, nil
}

// AddAmount adds f to the "amount" field.
func (m *ExpenseMutation) AddAmount(f float64) {
	if m.addamount != nil {
		*m.addamount += f
	} else {
		m.addamount = &f
	}
}

// AddedAmount returns the value that was added to the "amount" field in this mutation.
func (m *ExpenseMutation) AddedAmount() (r float64, exists bool) {
	v := m.addamount
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmount resets all changes to the "amount" field.
func (m *ExpenseMutation) ResetAmount() {
	m.amount = nil
	m.addamount = nil
}

// SetCategory sets the "category" field.
func (m *ExpenseMutation) SetCategory(s string) {
	m.category = &s
}

// Category returns the value of the "category" field in the mutation.
func (m *ExpenseMutation) Category() (r string, exists bool) {
	v := m.category
	if v == nil {
		return
	}
	return *v, true
}

// OldCategory returns the old "category" field's value of the Expense entity.
// If the Expense object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExpenseMutation) OldCategory(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCategory is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCategory requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCategory: %w", err)
	}
	return oldValue.Category, nil
}

// ResetCategory resets all changes to the "category" field.
func (m *ExpenseMutation) ResetCategory() {
	m.category = nil
}

// SetMethod sets the "method" field.
func (m *ExpenseMutation) SetMethod(s string) {
	m.method = &s
}

// Method returns the value of the "method" field in the mutation.
func (m *ExpenseMutation) Method() (r string, exists bool) {
	v := m.method
	if v == nil {
		return
	}
	return *v, true
}

// OldMethod returns the old "method" field's value of the Expense entity.
// If the Expense object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExpenseMutation) OldMethod(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMethod is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMethod requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMethod: %w", err)
	}
	return oldValue.Method, nil
}

// ResetMethod resets all changes to the "method" field.
func (m *ExpenseMutation) ResetMethod() {
	m.method = nil
}

// SetOwnerName sets the "ownerName" field.
func (m *ExpenseMutation) SetOwnerName(s string) {
	m.ownerName = &s
}

// OwnerName returns the value of the "ownerName" field in the mutation.
func (m *ExpenseMutation) OwnerName() (r string, exists bool) {
	v := m.ownerName
	if v == nil {
		return
	}
	return *v, true
}

// OldOwnerName returns the old "ownerName" field's value of the Expense entity.
// If the Expense object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExpenseMutation) OldOwnerName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOwnerName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOwnerName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwnerName: %w", err)
	}
	return oldValue.OwnerName, nil
}

// ResetOwnerName resets all changes to the "ownerName" field.
func (m *ExpenseMutation) ResetOwnerName() {
	m.ownerName = nil
}

// SetTimestamp sets the "timestamp" field.
func (m *ExpenseMutation) SetTimestamp(t time.Time) {
	m.timestamp = &t
}

// Timestamp returns the value of the "timestamp" field in the mutation.
func (m *ExpenseMutation) Timestamp() (r time.Time, exists bool) {
	v := m.timestamp
	if v == nil {
		return
	}
	return *v, true
}

// OldTimestamp returns the old "timestamp" field's value of the Expense entity.
// If the Expense object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExpenseMutation) OldTimestamp(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTimestamp is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTimestamp requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTimestamp: %w", err)
	}
	return oldValue.Timestamp, nil
}

// ResetTimestamp resets all changes to the "timestamp" field.
func (m *ExpenseMutation) ResetTimestamp() {
	m.timestamp = nil
}

// SetCreatedAt sets the "createdAt" field.
func (m *ExpenseMutation) SetCreatedAt(t time.Time) {
	m.createdAt = &t
}

// CreatedAt returns the value of the "createdAt" field in the mutation.
func (m *ExpenseMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.createdAt
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "createdAt" field's value of the Expense entity.
// If the Expense object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ExpenseMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "createdAt" field.
func (m *ExpenseMutation) ResetCreatedAt() {
	m.createdAt = nil
}

// Where appends a list predicates to the ExpenseMutation builder.
func (m *ExpenseMutation) Where(ps ...predicate.Expense) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ExpenseMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Expense).
func (m *ExpenseMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ExpenseMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.amount != nil {
		fields = append(fields, expense.FieldAmount)
	}
	if m.category != nil {
		fields = append(fields, expense.FieldCategory)
	}
	if m.method != nil {
		fields = append(fields, expense.FieldMethod)
	}
	if m.ownerName != nil {
		fields = append(fields, expense.FieldOwnerName)
	}
	if m.timestamp != nil {
		fields = append(fields, expense.FieldTimestamp)
	}
	if m.createdAt != nil {
		fields = append(fields, expense.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ExpenseMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case expense.FieldAmount:
		return m.Amount()
	case expense.FieldCategory:
		return m.Category()
	case expense.FieldMethod:
		return m.Method()
	case expense.FieldOwnerName:
		return m.OwnerName()
	case expense.FieldTimestamp:
		return m.Timestamp()
	case expense.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ExpenseMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case expense.FieldAmount:
		return m.OldAmount(ctx)
	case expense.FieldCategory:
		return m.OldCategory(ctx)
	case expense.FieldMethod:
		return m.OldMethod(ctx)
	case expense.FieldOwnerName:
		return m.OldOwnerName(ctx)
	case expense.FieldTimestamp:
		return m.OldTimestamp(ctx)
	case expense.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Expense field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ExpenseMutation) SetField(name string, value ent.Value) error {
	switch name {
	case expense.FieldAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmount(v)
		return nil
	case expense.FieldCategory:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCategory(v)
		return nil
	case expense.FieldMethod:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMethod(v)
		return nil
	case expense.FieldOwnerName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwnerName(v)
		return nil
	case expense.FieldTimestamp:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTimestamp(v)
		return nil
	case expense.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Expense field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ExpenseMutation) AddedFields() []string {
	var fields []string
	if m.addamount != nil {
		fields = append(fields, expense.FieldAmount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ExpenseMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case expense.FieldAmount:
		return m.AddedAmount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ExpenseMutation) AddField(name string, value ent.Value) error {
	switch name {
	case expense.FieldAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmount(v)
		return nil
	}
	return fmt.Errorf("unknown Expense numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ExpenseMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ExpenseMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ExpenseMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Expense nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ExpenseMutation) ResetField(name string) error {
	switch name {
	case expense.FieldAmount:
		m.ResetAmount()
		return nil
	case expense.FieldCategory:
		m.ResetCategory()
		return nil
	case expense.FieldMethod:
		m.ResetMethod()
		return nil
	case expense.FieldOwnerName:
		m.ResetOwnerName()
		return nil
	case expense.FieldTimestamp:
		m.ResetTimestamp()
		return nil
	case expense.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Expense field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ExpenseMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ExpenseMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ExpenseMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ExpenseMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ExpenseMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ExpenseMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ExpenseMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Expense unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ExpenseMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Expense edge %s", name)
}