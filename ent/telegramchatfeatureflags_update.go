// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nekomeowww/insights-bot/ent/internal"
	"github.com/nekomeowww/insights-bot/ent/predicate"
	"github.com/nekomeowww/insights-bot/ent/telegramchatfeatureflags"
)

// TelegramChatFeatureFlagsUpdate is the builder for updating TelegramChatFeatureFlags entities.
type TelegramChatFeatureFlagsUpdate struct {
	config
	hooks    []Hook
	mutation *TelegramChatFeatureFlagsMutation
}

// Where appends a list predicates to the TelegramChatFeatureFlagsUpdate builder.
func (tcffu *TelegramChatFeatureFlagsUpdate) Where(ps ...predicate.TelegramChatFeatureFlags) *TelegramChatFeatureFlagsUpdate {
	tcffu.mutation.Where(ps...)
	return tcffu
}

// SetChatID sets the "chat_id" field.
func (tcffu *TelegramChatFeatureFlagsUpdate) SetChatID(i int64) *TelegramChatFeatureFlagsUpdate {
	tcffu.mutation.ResetChatID()
	tcffu.mutation.SetChatID(i)
	return tcffu
}

// AddChatID adds i to the "chat_id" field.
func (tcffu *TelegramChatFeatureFlagsUpdate) AddChatID(i int64) *TelegramChatFeatureFlagsUpdate {
	tcffu.mutation.AddChatID(i)
	return tcffu
}

// SetChatType sets the "chat_type" field.
func (tcffu *TelegramChatFeatureFlagsUpdate) SetChatType(s string) *TelegramChatFeatureFlagsUpdate {
	tcffu.mutation.SetChatType(s)
	return tcffu
}

// SetChatTitle sets the "chat_title" field.
func (tcffu *TelegramChatFeatureFlagsUpdate) SetChatTitle(s string) *TelegramChatFeatureFlagsUpdate {
	tcffu.mutation.SetChatTitle(s)
	return tcffu
}

// SetNillableChatTitle sets the "chat_title" field if the given value is not nil.
func (tcffu *TelegramChatFeatureFlagsUpdate) SetNillableChatTitle(s *string) *TelegramChatFeatureFlagsUpdate {
	if s != nil {
		tcffu.SetChatTitle(*s)
	}
	return tcffu
}

// SetFeatureChatHistoriesRecap sets the "feature_chat_histories_recap" field.
func (tcffu *TelegramChatFeatureFlagsUpdate) SetFeatureChatHistoriesRecap(b bool) *TelegramChatFeatureFlagsUpdate {
	tcffu.mutation.SetFeatureChatHistoriesRecap(b)
	return tcffu
}

// SetCreatedAt sets the "created_at" field.
func (tcffu *TelegramChatFeatureFlagsUpdate) SetCreatedAt(i int64) *TelegramChatFeatureFlagsUpdate {
	tcffu.mutation.ResetCreatedAt()
	tcffu.mutation.SetCreatedAt(i)
	return tcffu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tcffu *TelegramChatFeatureFlagsUpdate) SetNillableCreatedAt(i *int64) *TelegramChatFeatureFlagsUpdate {
	if i != nil {
		tcffu.SetCreatedAt(*i)
	}
	return tcffu
}

// AddCreatedAt adds i to the "created_at" field.
func (tcffu *TelegramChatFeatureFlagsUpdate) AddCreatedAt(i int64) *TelegramChatFeatureFlagsUpdate {
	tcffu.mutation.AddCreatedAt(i)
	return tcffu
}

// SetUpdatedAt sets the "updated_at" field.
func (tcffu *TelegramChatFeatureFlagsUpdate) SetUpdatedAt(i int64) *TelegramChatFeatureFlagsUpdate {
	tcffu.mutation.ResetUpdatedAt()
	tcffu.mutation.SetUpdatedAt(i)
	return tcffu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tcffu *TelegramChatFeatureFlagsUpdate) SetNillableUpdatedAt(i *int64) *TelegramChatFeatureFlagsUpdate {
	if i != nil {
		tcffu.SetUpdatedAt(*i)
	}
	return tcffu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tcffu *TelegramChatFeatureFlagsUpdate) AddUpdatedAt(i int64) *TelegramChatFeatureFlagsUpdate {
	tcffu.mutation.AddUpdatedAt(i)
	return tcffu
}

// Mutation returns the TelegramChatFeatureFlagsMutation object of the builder.
func (tcffu *TelegramChatFeatureFlagsUpdate) Mutation() *TelegramChatFeatureFlagsMutation {
	return tcffu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcffu *TelegramChatFeatureFlagsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TelegramChatFeatureFlagsMutation](ctx, tcffu.sqlSave, tcffu.mutation, tcffu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcffu *TelegramChatFeatureFlagsUpdate) SaveX(ctx context.Context) int {
	affected, err := tcffu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcffu *TelegramChatFeatureFlagsUpdate) Exec(ctx context.Context) error {
	_, err := tcffu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcffu *TelegramChatFeatureFlagsUpdate) ExecX(ctx context.Context) {
	if err := tcffu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tcffu *TelegramChatFeatureFlagsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(telegramchatfeatureflags.Table, telegramchatfeatureflags.Columns, sqlgraph.NewFieldSpec(telegramchatfeatureflags.FieldID, field.TypeUUID))
	if ps := tcffu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcffu.mutation.ChatID(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := tcffu.mutation.AddedChatID(); ok {
		_spec.AddField(telegramchatfeatureflags.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := tcffu.mutation.ChatType(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldChatType, field.TypeString, value)
	}
	if value, ok := tcffu.mutation.ChatTitle(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldChatTitle, field.TypeString, value)
	}
	if value, ok := tcffu.mutation.FeatureChatHistoriesRecap(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldFeatureChatHistoriesRecap, field.TypeBool, value)
	}
	if value, ok := tcffu.mutation.CreatedAt(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := tcffu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(telegramchatfeatureflags.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := tcffu.mutation.UpdatedAt(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tcffu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(telegramchatfeatureflags.FieldUpdatedAt, field.TypeInt64, value)
	}
	_spec.Node.Schema = tcffu.schemaConfig.TelegramChatFeatureFlags
	ctx = internal.NewSchemaConfigContext(ctx, tcffu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, tcffu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{telegramchatfeatureflags.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tcffu.mutation.done = true
	return n, nil
}

// TelegramChatFeatureFlagsUpdateOne is the builder for updating a single TelegramChatFeatureFlags entity.
type TelegramChatFeatureFlagsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TelegramChatFeatureFlagsMutation
}

// SetChatID sets the "chat_id" field.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) SetChatID(i int64) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.mutation.ResetChatID()
	tcffuo.mutation.SetChatID(i)
	return tcffuo
}

// AddChatID adds i to the "chat_id" field.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) AddChatID(i int64) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.mutation.AddChatID(i)
	return tcffuo
}

// SetChatType sets the "chat_type" field.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) SetChatType(s string) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.mutation.SetChatType(s)
	return tcffuo
}

// SetChatTitle sets the "chat_title" field.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) SetChatTitle(s string) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.mutation.SetChatTitle(s)
	return tcffuo
}

// SetNillableChatTitle sets the "chat_title" field if the given value is not nil.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) SetNillableChatTitle(s *string) *TelegramChatFeatureFlagsUpdateOne {
	if s != nil {
		tcffuo.SetChatTitle(*s)
	}
	return tcffuo
}

// SetFeatureChatHistoriesRecap sets the "feature_chat_histories_recap" field.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) SetFeatureChatHistoriesRecap(b bool) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.mutation.SetFeatureChatHistoriesRecap(b)
	return tcffuo
}

// SetCreatedAt sets the "created_at" field.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) SetCreatedAt(i int64) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.mutation.ResetCreatedAt()
	tcffuo.mutation.SetCreatedAt(i)
	return tcffuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) SetNillableCreatedAt(i *int64) *TelegramChatFeatureFlagsUpdateOne {
	if i != nil {
		tcffuo.SetCreatedAt(*i)
	}
	return tcffuo
}

// AddCreatedAt adds i to the "created_at" field.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) AddCreatedAt(i int64) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.mutation.AddCreatedAt(i)
	return tcffuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) SetUpdatedAt(i int64) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.mutation.ResetUpdatedAt()
	tcffuo.mutation.SetUpdatedAt(i)
	return tcffuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) SetNillableUpdatedAt(i *int64) *TelegramChatFeatureFlagsUpdateOne {
	if i != nil {
		tcffuo.SetUpdatedAt(*i)
	}
	return tcffuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) AddUpdatedAt(i int64) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.mutation.AddUpdatedAt(i)
	return tcffuo
}

// Mutation returns the TelegramChatFeatureFlagsMutation object of the builder.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) Mutation() *TelegramChatFeatureFlagsMutation {
	return tcffuo.mutation
}

// Where appends a list predicates to the TelegramChatFeatureFlagsUpdate builder.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) Where(ps ...predicate.TelegramChatFeatureFlags) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.mutation.Where(ps...)
	return tcffuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) Select(field string, fields ...string) *TelegramChatFeatureFlagsUpdateOne {
	tcffuo.fields = append([]string{field}, fields...)
	return tcffuo
}

// Save executes the query and returns the updated TelegramChatFeatureFlags entity.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) Save(ctx context.Context) (*TelegramChatFeatureFlags, error) {
	return withHooks[*TelegramChatFeatureFlags, TelegramChatFeatureFlagsMutation](ctx, tcffuo.sqlSave, tcffuo.mutation, tcffuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) SaveX(ctx context.Context) *TelegramChatFeatureFlags {
	node, err := tcffuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) Exec(ctx context.Context) error {
	_, err := tcffuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcffuo *TelegramChatFeatureFlagsUpdateOne) ExecX(ctx context.Context) {
	if err := tcffuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tcffuo *TelegramChatFeatureFlagsUpdateOne) sqlSave(ctx context.Context) (_node *TelegramChatFeatureFlags, err error) {
	_spec := sqlgraph.NewUpdateSpec(telegramchatfeatureflags.Table, telegramchatfeatureflags.Columns, sqlgraph.NewFieldSpec(telegramchatfeatureflags.FieldID, field.TypeUUID))
	id, ok := tcffuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TelegramChatFeatureFlags.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcffuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, telegramchatfeatureflags.FieldID)
		for _, f := range fields {
			if !telegramchatfeatureflags.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != telegramchatfeatureflags.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcffuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcffuo.mutation.ChatID(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := tcffuo.mutation.AddedChatID(); ok {
		_spec.AddField(telegramchatfeatureflags.FieldChatID, field.TypeInt64, value)
	}
	if value, ok := tcffuo.mutation.ChatType(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldChatType, field.TypeString, value)
	}
	if value, ok := tcffuo.mutation.ChatTitle(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldChatTitle, field.TypeString, value)
	}
	if value, ok := tcffuo.mutation.FeatureChatHistoriesRecap(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldFeatureChatHistoriesRecap, field.TypeBool, value)
	}
	if value, ok := tcffuo.mutation.CreatedAt(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := tcffuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(telegramchatfeatureflags.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := tcffuo.mutation.UpdatedAt(); ok {
		_spec.SetField(telegramchatfeatureflags.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tcffuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(telegramchatfeatureflags.FieldUpdatedAt, field.TypeInt64, value)
	}
	_spec.Node.Schema = tcffuo.schemaConfig.TelegramChatFeatureFlags
	ctx = internal.NewSchemaConfigContext(ctx, tcffuo.schemaConfig)
	_node = &TelegramChatFeatureFlags{config: tcffuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcffuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{telegramchatfeatureflags.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tcffuo.mutation.done = true
	return _node, nil
}
