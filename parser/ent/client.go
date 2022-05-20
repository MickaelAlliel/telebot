// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"mickaelalliel.com/telebot/parser/ent/migrate"

	"mickaelalliel.com/telebot/parser/ent/expense"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Expense is the client for interacting with the Expense builders.
	Expense *ExpenseClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Expense = NewExpenseClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Expense: NewExpenseClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Expense: NewExpenseClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Expense.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Expense.Use(hooks...)
}

// ExpenseClient is a client for the Expense schema.
type ExpenseClient struct {
	config
}

// NewExpenseClient returns a client for the Expense from the given config.
func NewExpenseClient(c config) *ExpenseClient {
	return &ExpenseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `expense.Hooks(f(g(h())))`.
func (c *ExpenseClient) Use(hooks ...Hook) {
	c.hooks.Expense = append(c.hooks.Expense, hooks...)
}

// Create returns a create builder for Expense.
func (c *ExpenseClient) Create() *ExpenseCreate {
	mutation := newExpenseMutation(c.config, OpCreate)
	return &ExpenseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Expense entities.
func (c *ExpenseClient) CreateBulk(builders ...*ExpenseCreate) *ExpenseCreateBulk {
	return &ExpenseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Expense.
func (c *ExpenseClient) Update() *ExpenseUpdate {
	mutation := newExpenseMutation(c.config, OpUpdate)
	return &ExpenseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ExpenseClient) UpdateOne(e *Expense) *ExpenseUpdateOne {
	mutation := newExpenseMutation(c.config, OpUpdateOne, withExpense(e))
	return &ExpenseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ExpenseClient) UpdateOneID(id int64) *ExpenseUpdateOne {
	mutation := newExpenseMutation(c.config, OpUpdateOne, withExpenseID(id))
	return &ExpenseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Expense.
func (c *ExpenseClient) Delete() *ExpenseDelete {
	mutation := newExpenseMutation(c.config, OpDelete)
	return &ExpenseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ExpenseClient) DeleteOne(e *Expense) *ExpenseDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ExpenseClient) DeleteOneID(id int64) *ExpenseDeleteOne {
	builder := c.Delete().Where(expense.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ExpenseDeleteOne{builder}
}

// Query returns a query builder for Expense.
func (c *ExpenseClient) Query() *ExpenseQuery {
	return &ExpenseQuery{
		config: c.config,
	}
}

// Get returns a Expense entity by its id.
func (c *ExpenseClient) Get(ctx context.Context, id int64) (*Expense, error) {
	return c.Query().Where(expense.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ExpenseClient) GetX(ctx context.Context, id int64) *Expense {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ExpenseClient) Hooks() []Hook {
	return c.hooks.Expense
}
