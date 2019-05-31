package broker

import (
	"github.com/uber/aresdb/metastore/common"
	"github.com/uber/aresdb/metastore"
	"github.com/uber/aresdb/memstore"
	"sync"
)

// BrokerSchemaMutator implements 2 interfaces
//  - metastore.TableSchemaMutator needed by SchemaFetchJob
//  - memstore.TableSchemaReader needed by compiler
type BrokerSchemaMutator struct {
	sync.RWMutex

	tables map[string]*memstore.TableSchema
}

func NewBrokerSchemaMutator() *BrokerSchemaMutator {
	return &BrokerSchemaMutator{}
}

// ====  metastore.TableSchemaMutator ====
// TODO: implement. TableSchemaMutator should update b.tables
func (b *BrokerSchemaMutator) ListTables() (tables []string, err error) {
	return
}

func (b *BrokerSchemaMutator) GetTable(name string) (table *common.Table, err error) {
	tableSchema, ok := b.tables[name]
	if !ok {
		err = metastore.ErrTableDoesNotExist
	}
	table = &tableSchema.Schema
	return
}

func (b *BrokerSchemaMutator) CreateTable(table *common.Table) (err error) {
	return
}
func (b *BrokerSchemaMutator) DeleteTable(name string) (err error) {
	return
}
func (b *BrokerSchemaMutator) UpdateTableConfig(table string, config common.TableConfig) (err error) {
	return
}
func (b *BrokerSchemaMutator) UpdateTable(table common.Table) (err error) {
	return
}
func (b *BrokerSchemaMutator) AddColumn(table string, column common.Column, appendToArchivingSortOrder bool) (err error) {
	return
}
func (b *BrokerSchemaMutator) UpdateColumn(table string, column string, config common.ColumnConfig) (err error) {
	return
}

func (b *BrokerSchemaMutator) DeleteColumn(table string, column string) (err error) {
	return
}

// ====  memstore.TableSchemaReader ====
// TODO: implement. these are used by compiler
func (b *BrokerSchemaMutator) GetSchema(table string) (tableSchema *memstore.TableSchema, err error) {
	return
}

func (b *BrokerSchemaMutator) GetSchemas() (schemas map[string]*memstore.TableSchema) {
	return
}
