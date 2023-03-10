// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"oruka/ent/predicate"
	"oruka/ent/timerecord"
	"oruka/ent/user"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTimeRecord = "TimeRecord"
	TypeUser       = "User"
)

// TimeRecordMutation represents an operation that mutates the TimeRecord nodes in the graph.
type TimeRecordMutation struct {
	config
	op                Op
	typ               string
	id                *int
	created_at        *time.Time
	updated_at        *time.Time
	clearedFields     map[string]struct{}
	timeKeeper        *int
	clearedtimeKeeper bool
	done              bool
	oldValue          func(context.Context) (*TimeRecord, error)
	predicates        []predicate.TimeRecord
}

var _ ent.Mutation = (*TimeRecordMutation)(nil)

// timerecordOption allows management of the mutation configuration using functional options.
type timerecordOption func(*TimeRecordMutation)

// newTimeRecordMutation creates new mutation for the TimeRecord entity.
func newTimeRecordMutation(c config, op Op, opts ...timerecordOption) *TimeRecordMutation {
	m := &TimeRecordMutation{
		config:        c,
		op:            op,
		typ:           TypeTimeRecord,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTimeRecordID sets the ID field of the mutation.
func withTimeRecordID(id int) timerecordOption {
	return func(m *TimeRecordMutation) {
		var (
			err   error
			once  sync.Once
			value *TimeRecord
		)
		m.oldValue = func(ctx context.Context) (*TimeRecord, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().TimeRecord.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTimeRecord sets the old TimeRecord of the mutation.
func withTimeRecord(node *TimeRecord) timerecordOption {
	return func(m *TimeRecordMutation) {
		m.oldValue = func(context.Context) (*TimeRecord, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TimeRecordMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TimeRecordMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TimeRecordMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TimeRecordMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().TimeRecord.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *TimeRecordMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TimeRecordMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the TimeRecord entity.
// If the TimeRecord object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TimeRecordMutation) OldCreatedAt(ctx context.Context) (v *time.Time, err error) {
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

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TimeRecordMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TimeRecordMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TimeRecordMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the TimeRecord entity.
// If the TimeRecord object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TimeRecordMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TimeRecordMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetTimeKeeperID sets the "timeKeeper" edge to the User entity by id.
func (m *TimeRecordMutation) SetTimeKeeperID(id int) {
	m.timeKeeper = &id
}

// ClearTimeKeeper clears the "timeKeeper" edge to the User entity.
func (m *TimeRecordMutation) ClearTimeKeeper() {
	m.clearedtimeKeeper = true
}

// TimeKeeperCleared reports if the "timeKeeper" edge to the User entity was cleared.
func (m *TimeRecordMutation) TimeKeeperCleared() bool {
	return m.clearedtimeKeeper
}

// TimeKeeperID returns the "timeKeeper" edge ID in the mutation.
func (m *TimeRecordMutation) TimeKeeperID() (id int, exists bool) {
	if m.timeKeeper != nil {
		return *m.timeKeeper, true
	}
	return
}

// TimeKeeperIDs returns the "timeKeeper" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// TimeKeeperID instead. It exists only for internal usage by the builders.
func (m *TimeRecordMutation) TimeKeeperIDs() (ids []int) {
	if id := m.timeKeeper; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetTimeKeeper resets all changes to the "timeKeeper" edge.
func (m *TimeRecordMutation) ResetTimeKeeper() {
	m.timeKeeper = nil
	m.clearedtimeKeeper = false
}

// Where appends a list predicates to the TimeRecordMutation builder.
func (m *TimeRecordMutation) Where(ps ...predicate.TimeRecord) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TimeRecordMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TimeRecordMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.TimeRecord, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TimeRecordMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TimeRecordMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (TimeRecord).
func (m *TimeRecordMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TimeRecordMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.created_at != nil {
		fields = append(fields, timerecord.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, timerecord.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TimeRecordMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case timerecord.FieldCreatedAt:
		return m.CreatedAt()
	case timerecord.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TimeRecordMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case timerecord.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case timerecord.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown TimeRecord field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TimeRecordMutation) SetField(name string, value ent.Value) error {
	switch name {
	case timerecord.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case timerecord.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown TimeRecord field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TimeRecordMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TimeRecordMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TimeRecordMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown TimeRecord numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TimeRecordMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TimeRecordMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TimeRecordMutation) ClearField(name string) error {
	return fmt.Errorf("unknown TimeRecord nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TimeRecordMutation) ResetField(name string) error {
	switch name {
	case timerecord.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case timerecord.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown TimeRecord field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TimeRecordMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.timeKeeper != nil {
		edges = append(edges, timerecord.EdgeTimeKeeper)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TimeRecordMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case timerecord.EdgeTimeKeeper:
		if id := m.timeKeeper; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TimeRecordMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TimeRecordMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TimeRecordMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtimeKeeper {
		edges = append(edges, timerecord.EdgeTimeKeeper)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TimeRecordMutation) EdgeCleared(name string) bool {
	switch name {
	case timerecord.EdgeTimeKeeper:
		return m.clearedtimeKeeper
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TimeRecordMutation) ClearEdge(name string) error {
	switch name {
	case timerecord.EdgeTimeKeeper:
		m.ClearTimeKeeper()
		return nil
	}
	return fmt.Errorf("unknown TimeRecord unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TimeRecordMutation) ResetEdge(name string) error {
	switch name {
	case timerecord.EdgeTimeKeeper:
		m.ResetTimeKeeper()
		return nil
	}
	return fmt.Errorf("unknown TimeRecord edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	created_at         *time.Time
	updated_at         *time.Time
	first_name         *string
	last_name          *string
	manavis_id         *int
	addmanavis_id      *int
	grade              *int
	addgrade           *int
	clearedFields      map[string]struct{}
	timeRecords        map[int]struct{}
	removedtimeRecords map[int]struct{}
	clearedtimeRecords bool
	done               bool
	oldValue           func(context.Context) (*User, error)
	predicates         []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v *time.Time, err error) {
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

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetFirstName sets the "first_name" field.
func (m *UserMutation) SetFirstName(s string) {
	m.first_name = &s
}

// FirstName returns the value of the "first_name" field in the mutation.
func (m *UserMutation) FirstName() (r string, exists bool) {
	v := m.first_name
	if v == nil {
		return
	}
	return *v, true
}

// OldFirstName returns the old "first_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldFirstName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFirstName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFirstName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFirstName: %w", err)
	}
	return oldValue.FirstName, nil
}

// ResetFirstName resets all changes to the "first_name" field.
func (m *UserMutation) ResetFirstName() {
	m.first_name = nil
}

// SetLastName sets the "last_name" field.
func (m *UserMutation) SetLastName(s string) {
	m.last_name = &s
}

// LastName returns the value of the "last_name" field in the mutation.
func (m *UserMutation) LastName() (r string, exists bool) {
	v := m.last_name
	if v == nil {
		return
	}
	return *v, true
}

// OldLastName returns the old "last_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldLastName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLastName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLastName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastName: %w", err)
	}
	return oldValue.LastName, nil
}

// ResetLastName resets all changes to the "last_name" field.
func (m *UserMutation) ResetLastName() {
	m.last_name = nil
}

// SetManavisID sets the "manavis_id" field.
func (m *UserMutation) SetManavisID(i int) {
	m.manavis_id = &i
	m.addmanavis_id = nil
}

// ManavisID returns the value of the "manavis_id" field in the mutation.
func (m *UserMutation) ManavisID() (r int, exists bool) {
	v := m.manavis_id
	if v == nil {
		return
	}
	return *v, true
}

// OldManavisID returns the old "manavis_id" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldManavisID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldManavisID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldManavisID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldManavisID: %w", err)
	}
	return oldValue.ManavisID, nil
}

// AddManavisID adds i to the "manavis_id" field.
func (m *UserMutation) AddManavisID(i int) {
	if m.addmanavis_id != nil {
		*m.addmanavis_id += i
	} else {
		m.addmanavis_id = &i
	}
}

// AddedManavisID returns the value that was added to the "manavis_id" field in this mutation.
func (m *UserMutation) AddedManavisID() (r int, exists bool) {
	v := m.addmanavis_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetManavisID resets all changes to the "manavis_id" field.
func (m *UserMutation) ResetManavisID() {
	m.manavis_id = nil
	m.addmanavis_id = nil
}

// SetGrade sets the "grade" field.
func (m *UserMutation) SetGrade(i int) {
	m.grade = &i
	m.addgrade = nil
}

// Grade returns the value of the "grade" field in the mutation.
func (m *UserMutation) Grade() (r int, exists bool) {
	v := m.grade
	if v == nil {
		return
	}
	return *v, true
}

// OldGrade returns the old "grade" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldGrade(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGrade is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGrade requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGrade: %w", err)
	}
	return oldValue.Grade, nil
}

// AddGrade adds i to the "grade" field.
func (m *UserMutation) AddGrade(i int) {
	if m.addgrade != nil {
		*m.addgrade += i
	} else {
		m.addgrade = &i
	}
}

// AddedGrade returns the value that was added to the "grade" field in this mutation.
func (m *UserMutation) AddedGrade() (r int, exists bool) {
	v := m.addgrade
	if v == nil {
		return
	}
	return *v, true
}

// ResetGrade resets all changes to the "grade" field.
func (m *UserMutation) ResetGrade() {
	m.grade = nil
	m.addgrade = nil
}

// AddTimeRecordIDs adds the "timeRecords" edge to the TimeRecord entity by ids.
func (m *UserMutation) AddTimeRecordIDs(ids ...int) {
	if m.timeRecords == nil {
		m.timeRecords = make(map[int]struct{})
	}
	for i := range ids {
		m.timeRecords[ids[i]] = struct{}{}
	}
}

// ClearTimeRecords clears the "timeRecords" edge to the TimeRecord entity.
func (m *UserMutation) ClearTimeRecords() {
	m.clearedtimeRecords = true
}

// TimeRecordsCleared reports if the "timeRecords" edge to the TimeRecord entity was cleared.
func (m *UserMutation) TimeRecordsCleared() bool {
	return m.clearedtimeRecords
}

// RemoveTimeRecordIDs removes the "timeRecords" edge to the TimeRecord entity by IDs.
func (m *UserMutation) RemoveTimeRecordIDs(ids ...int) {
	if m.removedtimeRecords == nil {
		m.removedtimeRecords = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.timeRecords, ids[i])
		m.removedtimeRecords[ids[i]] = struct{}{}
	}
}

// RemovedTimeRecords returns the removed IDs of the "timeRecords" edge to the TimeRecord entity.
func (m *UserMutation) RemovedTimeRecordsIDs() (ids []int) {
	for id := range m.removedtimeRecords {
		ids = append(ids, id)
	}
	return
}

// TimeRecordsIDs returns the "timeRecords" edge IDs in the mutation.
func (m *UserMutation) TimeRecordsIDs() (ids []int) {
	for id := range m.timeRecords {
		ids = append(ids, id)
	}
	return
}

// ResetTimeRecords resets all changes to the "timeRecords" edge.
func (m *UserMutation) ResetTimeRecords() {
	m.timeRecords = nil
	m.clearedtimeRecords = false
	m.removedtimeRecords = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 6)
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	if m.first_name != nil {
		fields = append(fields, user.FieldFirstName)
	}
	if m.last_name != nil {
		fields = append(fields, user.FieldLastName)
	}
	if m.manavis_id != nil {
		fields = append(fields, user.FieldManavisID)
	}
	if m.grade != nil {
		fields = append(fields, user.FieldGrade)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	case user.FieldFirstName:
		return m.FirstName()
	case user.FieldLastName:
		return m.LastName()
	case user.FieldManavisID:
		return m.ManavisID()
	case user.FieldGrade:
		return m.Grade()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case user.FieldFirstName:
		return m.OldFirstName(ctx)
	case user.FieldLastName:
		return m.OldLastName(ctx)
	case user.FieldManavisID:
		return m.OldManavisID(ctx)
	case user.FieldGrade:
		return m.OldGrade(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case user.FieldFirstName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFirstName(v)
		return nil
	case user.FieldLastName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastName(v)
		return nil
	case user.FieldManavisID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetManavisID(v)
		return nil
	case user.FieldGrade:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGrade(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addmanavis_id != nil {
		fields = append(fields, user.FieldManavisID)
	}
	if m.addgrade != nil {
		fields = append(fields, user.FieldGrade)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldManavisID:
		return m.AddedManavisID()
	case user.FieldGrade:
		return m.AddedGrade()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldManavisID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddManavisID(v)
		return nil
	case user.FieldGrade:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGrade(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case user.FieldFirstName:
		m.ResetFirstName()
		return nil
	case user.FieldLastName:
		m.ResetLastName()
		return nil
	case user.FieldManavisID:
		m.ResetManavisID()
		return nil
	case user.FieldGrade:
		m.ResetGrade()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.timeRecords != nil {
		edges = append(edges, user.EdgeTimeRecords)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTimeRecords:
		ids := make([]ent.Value, 0, len(m.timeRecords))
		for id := range m.timeRecords {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtimeRecords != nil {
		edges = append(edges, user.EdgeTimeRecords)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTimeRecords:
		ids := make([]ent.Value, 0, len(m.removedtimeRecords))
		for id := range m.removedtimeRecords {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtimeRecords {
		edges = append(edges, user.EdgeTimeRecords)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeTimeRecords:
		return m.clearedtimeRecords
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeTimeRecords:
		m.ResetTimeRecords()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
