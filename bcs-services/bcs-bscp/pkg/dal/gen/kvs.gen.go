// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/TencentBlueKing/bk-bcs/bcs-services/bcs-bscp/pkg/dal/table"
)

func newKv(db *gorm.DB, opts ...gen.DOOption) kv {
	_kv := kv{}

	_kv.kvDo.UseDB(db, opts...)
	_kv.kvDo.UseModel(&table.Kv{})

	tableName := _kv.kvDo.TableName()
	_kv.ALL = field.NewAsterisk(tableName)
	_kv.ID = field.NewUint32(tableName, "id")
	_kv.KvState = field.NewString(tableName, "kv_state")
	_kv.Key = field.NewString(tableName, "key")
	_kv.Memo = field.NewString(tableName, "memo")
	_kv.KvType = field.NewString(tableName, "kv_type")
	_kv.Version = field.NewUint32(tableName, "version")
	_kv.BizID = field.NewUint32(tableName, "biz_id")
	_kv.AppID = field.NewUint32(tableName, "app_id")
	_kv.Creator = field.NewString(tableName, "creator")
	_kv.Reviser = field.NewString(tableName, "reviser")
	_kv.CreatedAt = field.NewTime(tableName, "created_at")
	_kv.UpdatedAt = field.NewTime(tableName, "updated_at")
	_kv.Signature = field.NewString(tableName, "signature")
	_kv.ByteSize = field.NewUint64(tableName, "byte_size")
	_kv.Md5 = field.NewString(tableName, "md5")

	_kv.fillFieldMap()

	return _kv
}

type kv struct {
	kvDo kvDo

	ALL       field.Asterisk
	ID        field.Uint32
	KvState   field.String
	Key       field.String
	Memo      field.String
	KvType    field.String
	Version   field.Uint32
	BizID     field.Uint32
	AppID     field.Uint32
	Creator   field.String
	Reviser   field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	Signature field.String
	ByteSize  field.Uint64
	Md5       field.String

	fieldMap map[string]field.Expr
}

func (k kv) Table(newTableName string) *kv {
	k.kvDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k kv) As(alias string) *kv {
	k.kvDo.DO = *(k.kvDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *kv) updateTableName(table string) *kv {
	k.ALL = field.NewAsterisk(table)
	k.ID = field.NewUint32(table, "id")
	k.KvState = field.NewString(table, "kv_state")
	k.Key = field.NewString(table, "key")
	k.Memo = field.NewString(table, "memo")
	k.KvType = field.NewString(table, "kv_type")
	k.Version = field.NewUint32(table, "version")
	k.BizID = field.NewUint32(table, "biz_id")
	k.AppID = field.NewUint32(table, "app_id")
	k.Creator = field.NewString(table, "creator")
	k.Reviser = field.NewString(table, "reviser")
	k.CreatedAt = field.NewTime(table, "created_at")
	k.UpdatedAt = field.NewTime(table, "updated_at")
	k.Signature = field.NewString(table, "signature")
	k.ByteSize = field.NewUint64(table, "byte_size")
	k.Md5 = field.NewString(table, "md5")

	k.fillFieldMap()

	return k
}

func (k *kv) WithContext(ctx context.Context) IKvDo { return k.kvDo.WithContext(ctx) }

func (k kv) TableName() string { return k.kvDo.TableName() }

func (k kv) Alias() string { return k.kvDo.Alias() }

func (k kv) Columns(cols ...field.Expr) gen.Columns { return k.kvDo.Columns(cols...) }

func (k *kv) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *kv) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 15)
	k.fieldMap["id"] = k.ID
	k.fieldMap["kv_state"] = k.KvState
	k.fieldMap["key"] = k.Key
	k.fieldMap["memo"] = k.Memo
	k.fieldMap["kv_type"] = k.KvType
	k.fieldMap["version"] = k.Version
	k.fieldMap["biz_id"] = k.BizID
	k.fieldMap["app_id"] = k.AppID
	k.fieldMap["creator"] = k.Creator
	k.fieldMap["reviser"] = k.Reviser
	k.fieldMap["created_at"] = k.CreatedAt
	k.fieldMap["updated_at"] = k.UpdatedAt
	k.fieldMap["signature"] = k.Signature
	k.fieldMap["byte_size"] = k.ByteSize
	k.fieldMap["md5"] = k.Md5
}

func (k kv) clone(db *gorm.DB) kv {
	k.kvDo.ReplaceConnPool(db.Statement.ConnPool)
	return k
}

func (k kv) replaceDB(db *gorm.DB) kv {
	k.kvDo.ReplaceDB(db)
	return k
}

type kvDo struct{ gen.DO }

type IKvDo interface {
	gen.SubQuery
	Debug() IKvDo
	WithContext(ctx context.Context) IKvDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IKvDo
	WriteDB() IKvDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IKvDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IKvDo
	Not(conds ...gen.Condition) IKvDo
	Or(conds ...gen.Condition) IKvDo
	Select(conds ...field.Expr) IKvDo
	Where(conds ...gen.Condition) IKvDo
	Order(conds ...field.Expr) IKvDo
	Distinct(cols ...field.Expr) IKvDo
	Omit(cols ...field.Expr) IKvDo
	Join(table schema.Tabler, on ...field.Expr) IKvDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IKvDo
	RightJoin(table schema.Tabler, on ...field.Expr) IKvDo
	Group(cols ...field.Expr) IKvDo
	Having(conds ...gen.Condition) IKvDo
	Limit(limit int) IKvDo
	Offset(offset int) IKvDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IKvDo
	Unscoped() IKvDo
	Create(values ...*table.Kv) error
	CreateInBatches(values []*table.Kv, batchSize int) error
	Save(values ...*table.Kv) error
	First() (*table.Kv, error)
	Take() (*table.Kv, error)
	Last() (*table.Kv, error)
	Find() ([]*table.Kv, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Kv, err error)
	FindInBatches(result *[]*table.Kv, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.Kv) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IKvDo
	Assign(attrs ...field.AssignExpr) IKvDo
	Joins(fields ...field.RelationField) IKvDo
	Preload(fields ...field.RelationField) IKvDo
	FirstOrInit() (*table.Kv, error)
	FirstOrCreate() (*table.Kv, error)
	FindByPage(offset int, limit int) (result []*table.Kv, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IKvDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (k kvDo) Debug() IKvDo {
	return k.withDO(k.DO.Debug())
}

func (k kvDo) WithContext(ctx context.Context) IKvDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k kvDo) ReadDB() IKvDo {
	return k.Clauses(dbresolver.Read)
}

func (k kvDo) WriteDB() IKvDo {
	return k.Clauses(dbresolver.Write)
}

func (k kvDo) Session(config *gorm.Session) IKvDo {
	return k.withDO(k.DO.Session(config))
}

func (k kvDo) Clauses(conds ...clause.Expression) IKvDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k kvDo) Returning(value interface{}, columns ...string) IKvDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k kvDo) Not(conds ...gen.Condition) IKvDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k kvDo) Or(conds ...gen.Condition) IKvDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k kvDo) Select(conds ...field.Expr) IKvDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k kvDo) Where(conds ...gen.Condition) IKvDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k kvDo) Order(conds ...field.Expr) IKvDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k kvDo) Distinct(cols ...field.Expr) IKvDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k kvDo) Omit(cols ...field.Expr) IKvDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k kvDo) Join(table schema.Tabler, on ...field.Expr) IKvDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k kvDo) LeftJoin(table schema.Tabler, on ...field.Expr) IKvDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k kvDo) RightJoin(table schema.Tabler, on ...field.Expr) IKvDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k kvDo) Group(cols ...field.Expr) IKvDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k kvDo) Having(conds ...gen.Condition) IKvDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k kvDo) Limit(limit int) IKvDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k kvDo) Offset(offset int) IKvDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k kvDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IKvDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k kvDo) Unscoped() IKvDo {
	return k.withDO(k.DO.Unscoped())
}

func (k kvDo) Create(values ...*table.Kv) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k kvDo) CreateInBatches(values []*table.Kv, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k kvDo) Save(values ...*table.Kv) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k kvDo) First() (*table.Kv, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.Kv), nil
	}
}

func (k kvDo) Take() (*table.Kv, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.Kv), nil
	}
}

func (k kvDo) Last() (*table.Kv, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.Kv), nil
	}
}

func (k kvDo) Find() ([]*table.Kv, error) {
	result, err := k.DO.Find()
	return result.([]*table.Kv), err
}

func (k kvDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Kv, err error) {
	buf := make([]*table.Kv, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k kvDo) FindInBatches(result *[]*table.Kv, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k kvDo) Attrs(attrs ...field.AssignExpr) IKvDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k kvDo) Assign(attrs ...field.AssignExpr) IKvDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k kvDo) Joins(fields ...field.RelationField) IKvDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k kvDo) Preload(fields ...field.RelationField) IKvDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k kvDo) FirstOrInit() (*table.Kv, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.Kv), nil
	}
}

func (k kvDo) FirstOrCreate() (*table.Kv, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.Kv), nil
	}
}

func (k kvDo) FindByPage(offset int, limit int) (result []*table.Kv, count int64, err error) {
	result, err = k.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = k.Offset(-1).Limit(-1).Count()
	return
}

func (k kvDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k kvDo) Scan(result interface{}) (err error) {
	return k.DO.Scan(result)
}

func (k kvDo) Delete(models ...*table.Kv) (result gen.ResultInfo, err error) {
	return k.DO.Delete(models)
}

func (k *kvDo) withDO(do gen.Dao) *kvDo {
	k.DO = *do.(*gen.DO)
	return k
}
