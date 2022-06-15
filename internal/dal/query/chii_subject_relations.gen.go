// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/internal/dal/dao"
)

func newSubjectRelation(db *gorm.DB) subjectRelation {
	_subjectRelation := subjectRelation{}

	_subjectRelation.subjectRelationDo.UseDB(db)
	_subjectRelation.subjectRelationDo.UseModel(&dao.SubjectRelation{})

	tableName := _subjectRelation.subjectRelationDo.TableName()
	_subjectRelation.ALL = field.NewField(tableName, "*")
	_subjectRelation.SubjectID = field.NewField(tableName, "rlt_subject_id")
	_subjectRelation.SubjectTypeID = field.NewUint8(tableName, "rlt_subject_type_id")
	_subjectRelation.RelationType = field.NewUint16(tableName, "rlt_relation_type")
	_subjectRelation.RelatedSubjectID = field.NewField(tableName, "rlt_related_subject_id")
	_subjectRelation.RelatedSubjectTypeID = field.NewUint8(tableName, "rlt_related_subject_type_id")
	_subjectRelation.ViceVersa = field.NewBool(tableName, "rlt_vice_versa")
	_subjectRelation.Order = field.NewUint8(tableName, "rlt_order")
	_subjectRelation.Subject = subjectRelationHasOneSubject{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Subject", "dao.Subject"),
		Fields: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Subject.Fields", "dao.SubjectField"),
		},
	}

	_subjectRelation.fillFieldMap()

	return _subjectRelation
}

type subjectRelation struct {
	subjectRelationDo subjectRelationDo

	ALL                  field.Field
	SubjectID            field.Field
	SubjectTypeID        field.Uint8
	RelationType         field.Uint16
	RelatedSubjectID     field.Field
	RelatedSubjectTypeID field.Uint8
	ViceVersa            field.Bool
	Order                field.Uint8
	Subject              subjectRelationHasOneSubject

	fieldMap map[string]field.Expr
}

func (s subjectRelation) Table(newTableName string) *subjectRelation {
	s.subjectRelationDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s subjectRelation) As(alias string) *subjectRelation {
	s.subjectRelationDo.DO = *(s.subjectRelationDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *subjectRelation) updateTableName(table string) *subjectRelation {
	s.ALL = field.NewField(table, "*")
	s.SubjectID = field.NewField(table, "rlt_subject_id")
	s.SubjectTypeID = field.NewUint8(table, "rlt_subject_type_id")
	s.RelationType = field.NewUint16(table, "rlt_relation_type")
	s.RelatedSubjectID = field.NewField(table, "rlt_related_subject_id")
	s.RelatedSubjectTypeID = field.NewUint8(table, "rlt_related_subject_type_id")
	s.ViceVersa = field.NewBool(table, "rlt_vice_versa")
	s.Order = field.NewUint8(table, "rlt_order")

	s.fillFieldMap()

	return s
}

func (s *subjectRelation) WithContext(ctx context.Context) *subjectRelationDo {
	return s.subjectRelationDo.WithContext(ctx)
}

func (s subjectRelation) TableName() string { return s.subjectRelationDo.TableName() }

func (s subjectRelation) Alias() string { return s.subjectRelationDo.Alias() }

func (s *subjectRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *subjectRelation) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["rlt_subject_id"] = s.SubjectID
	s.fieldMap["rlt_subject_type_id"] = s.SubjectTypeID
	s.fieldMap["rlt_relation_type"] = s.RelationType
	s.fieldMap["rlt_related_subject_id"] = s.RelatedSubjectID
	s.fieldMap["rlt_related_subject_type_id"] = s.RelatedSubjectTypeID
	s.fieldMap["rlt_vice_versa"] = s.ViceVersa
	s.fieldMap["rlt_order"] = s.Order

}

func (s subjectRelation) clone(db *gorm.DB) subjectRelation {
	s.subjectRelationDo.ReplaceDB(db)
	return s
}

type subjectRelationHasOneSubject struct {
	db *gorm.DB

	field.RelationField

	Fields struct {
		field.RelationField
	}
}

func (a subjectRelationHasOneSubject) Where(conds ...field.Expr) *subjectRelationHasOneSubject {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a subjectRelationHasOneSubject) WithContext(ctx context.Context) *subjectRelationHasOneSubject {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a subjectRelationHasOneSubject) Model(m *dao.SubjectRelation) *subjectRelationHasOneSubjectTx {
	return &subjectRelationHasOneSubjectTx{a.db.Model(m).Association(a.Name())}
}

type subjectRelationHasOneSubjectTx struct{ tx *gorm.Association }

func (a subjectRelationHasOneSubjectTx) Find() (result *dao.Subject, err error) {
	return result, a.tx.Find(&result)
}

func (a subjectRelationHasOneSubjectTx) Append(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a subjectRelationHasOneSubjectTx) Replace(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a subjectRelationHasOneSubjectTx) Delete(values ...*dao.Subject) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a subjectRelationHasOneSubjectTx) Clear() error {
	return a.tx.Clear()
}

func (a subjectRelationHasOneSubjectTx) Count() int64 {
	return a.tx.Count()
}

type subjectRelationDo struct{ gen.DO }

func (s subjectRelationDo) Debug() *subjectRelationDo {
	return s.withDO(s.DO.Debug())
}

func (s subjectRelationDo) WithContext(ctx context.Context) *subjectRelationDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s subjectRelationDo) ReadDB(ctx context.Context) *subjectRelationDo {
	return s.WithContext(ctx).Clauses(dbresolver.Read)
}

func (s subjectRelationDo) WriteDB(ctx context.Context) *subjectRelationDo {
	return s.WithContext(ctx).Clauses(dbresolver.Write)
}

func (s subjectRelationDo) Clauses(conds ...clause.Expression) *subjectRelationDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s subjectRelationDo) Returning(value interface{}, columns ...string) *subjectRelationDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s subjectRelationDo) Not(conds ...gen.Condition) *subjectRelationDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s subjectRelationDo) Or(conds ...gen.Condition) *subjectRelationDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s subjectRelationDo) Select(conds ...field.Expr) *subjectRelationDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s subjectRelationDo) Where(conds ...gen.Condition) *subjectRelationDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s subjectRelationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *subjectRelationDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s subjectRelationDo) Order(conds ...field.Expr) *subjectRelationDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s subjectRelationDo) Distinct(cols ...field.Expr) *subjectRelationDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s subjectRelationDo) Omit(cols ...field.Expr) *subjectRelationDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s subjectRelationDo) Join(table schema.Tabler, on ...field.Expr) *subjectRelationDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s subjectRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *subjectRelationDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s subjectRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) *subjectRelationDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s subjectRelationDo) Group(cols ...field.Expr) *subjectRelationDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s subjectRelationDo) Having(conds ...gen.Condition) *subjectRelationDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s subjectRelationDo) Limit(limit int) *subjectRelationDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s subjectRelationDo) Offset(offset int) *subjectRelationDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s subjectRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *subjectRelationDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s subjectRelationDo) Unscoped() *subjectRelationDo {
	return s.withDO(s.DO.Unscoped())
}

func (s subjectRelationDo) Create(values ...*dao.SubjectRelation) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s subjectRelationDo) CreateInBatches(values []*dao.SubjectRelation, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s subjectRelationDo) Save(values ...*dao.SubjectRelation) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s subjectRelationDo) First() (*dao.SubjectRelation, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectRelation), nil
	}
}

func (s subjectRelationDo) Take() (*dao.SubjectRelation, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectRelation), nil
	}
}

func (s subjectRelationDo) Last() (*dao.SubjectRelation, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectRelation), nil
	}
}

func (s subjectRelationDo) Find() ([]*dao.SubjectRelation, error) {
	result, err := s.DO.Find()
	return result.([]*dao.SubjectRelation), err
}

func (s subjectRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.SubjectRelation, err error) {
	buf := make([]*dao.SubjectRelation, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s subjectRelationDo) FindInBatches(result *[]*dao.SubjectRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s subjectRelationDo) Attrs(attrs ...field.AssignExpr) *subjectRelationDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s subjectRelationDo) Assign(attrs ...field.AssignExpr) *subjectRelationDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s subjectRelationDo) Joins(fields ...field.RelationField) *subjectRelationDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s subjectRelationDo) Preload(fields ...field.RelationField) *subjectRelationDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s subjectRelationDo) FirstOrInit() (*dao.SubjectRelation, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectRelation), nil
	}
}

func (s subjectRelationDo) FirstOrCreate() (*dao.SubjectRelation, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.SubjectRelation), nil
	}
}

func (s subjectRelationDo) FindByPage(offset int, limit int) (result []*dao.SubjectRelation, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s subjectRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s *subjectRelationDo) withDO(do gen.Dao) *subjectRelationDo {
	s.DO = *do.(*gen.DO)
	return s
}
