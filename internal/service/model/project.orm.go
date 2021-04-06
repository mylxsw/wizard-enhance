package model

// !!! DO NOT EDIT THIS FILE

import (
	"context"
	"encoding/json"
	"github.com/iancoleman/strcase"
	"github.com/mylxsw/coll"
	"github.com/mylxsw/eloquent/query"
	"gopkg.in/guregu/null.v3"
	"time"
)

func init() {

	// AddProjectGlobalScope assign a global scope to a model for soft delete
	AddGlobalScopeForProject("soft_delete", func(builder query.Condition) {
		builder.WhereNull("deleted_at")
	})

}

// Project is a Project object
type Project struct {
	original     *projectOriginal
	projectModel *ProjectModel

	Id               int64
	Name             string
	Description      string
	Visibility       int8
	UserId           int
	SortLevel        int
	CatalogId        int
	CatalogFoldStyle int8
	CatalogSortStyle int8
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        null.Time
}

// As convert object to other type
// dst must be a pointer to struct
func (inst *Project) As(dst interface{}) error {
	return coll.CopyProperties(inst, dst)
}

// SetModel set model for Project
func (inst *Project) SetModel(projectModel *ProjectModel) {
	inst.projectModel = projectModel
}

// projectOriginal is an object which stores original Project from database
type projectOriginal struct {
	Id               int64
	Name             string
	Description      string
	Visibility       int8
	UserId           int
	SortLevel        int
	CatalogId        int
	CatalogFoldStyle int8
	CatalogSortStyle int8
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        null.Time
}

// Staled identify whether the object has been modified
func (inst *Project) Staled() bool {
	if inst.original == nil {
		inst.original = &projectOriginal{}
	}

	if inst.Id != inst.original.Id {
		return true
	}
	if inst.Name != inst.original.Name {
		return true
	}
	if inst.Description != inst.original.Description {
		return true
	}
	if inst.Visibility != inst.original.Visibility {
		return true
	}
	if inst.UserId != inst.original.UserId {
		return true
	}
	if inst.SortLevel != inst.original.SortLevel {
		return true
	}
	if inst.CatalogId != inst.original.CatalogId {
		return true
	}
	if inst.CatalogFoldStyle != inst.original.CatalogFoldStyle {
		return true
	}
	if inst.CatalogSortStyle != inst.original.CatalogSortStyle {
		return true
	}
	if inst.CreatedAt != inst.original.CreatedAt {
		return true
	}
	if inst.UpdatedAt != inst.original.UpdatedAt {
		return true
	}
	if inst.DeletedAt != inst.original.DeletedAt {
		return true
	}

	return false
}

// StaledKV return all fields has been modified
func (inst *Project) StaledKV() query.KV {
	kv := make(query.KV, 0)

	if inst.original == nil {
		inst.original = &projectOriginal{}
	}

	if inst.Id != inst.original.Id {
		kv["id"] = inst.Id
	}
	if inst.Name != inst.original.Name {
		kv["name"] = inst.Name
	}
	if inst.Description != inst.original.Description {
		kv["description"] = inst.Description
	}
	if inst.Visibility != inst.original.Visibility {
		kv["visibility"] = inst.Visibility
	}
	if inst.UserId != inst.original.UserId {
		kv["user_id"] = inst.UserId
	}
	if inst.SortLevel != inst.original.SortLevel {
		kv["sort_level"] = inst.SortLevel
	}
	if inst.CatalogId != inst.original.CatalogId {
		kv["catalog_id"] = inst.CatalogId
	}
	if inst.CatalogFoldStyle != inst.original.CatalogFoldStyle {
		kv["catalog_fold_style"] = inst.CatalogFoldStyle
	}
	if inst.CatalogSortStyle != inst.original.CatalogSortStyle {
		kv["catalog_sort_style"] = inst.CatalogSortStyle
	}
	if inst.CreatedAt != inst.original.CreatedAt {
		kv["created_at"] = inst.CreatedAt
	}
	if inst.UpdatedAt != inst.original.UpdatedAt {
		kv["updated_at"] = inst.UpdatedAt
	}
	if inst.DeletedAt != inst.original.DeletedAt {
		kv["deleted_at"] = inst.DeletedAt
	}

	return kv
}

// Save create a new model or update it
func (inst *Project) Save() error {
	if inst.projectModel == nil {
		return query.ErrModelNotSet
	}

	id, _, err := inst.projectModel.SaveOrUpdate(*inst)
	if err != nil {
		return err
	}

	inst.Id = id
	return nil
}

// Delete remove a project
func (inst *Project) Delete() error {
	if inst.projectModel == nil {
		return query.ErrModelNotSet
	}

	_, err := inst.projectModel.DeleteById(inst.Id)
	if err != nil {
		return err
	}

	return nil
}

// String convert instance to json string
func (inst *Project) String() string {
	rs, _ := json.Marshal(inst)
	return string(rs)
}

type projectScope struct {
	name  string
	apply func(builder query.Condition)
}

var projectGlobalScopes = make([]projectScope, 0)
var projectLocalScopes = make([]projectScope, 0)

// AddGlobalScopeForProject assign a global scope to a model
func AddGlobalScopeForProject(name string, apply func(builder query.Condition)) {
	projectGlobalScopes = append(projectGlobalScopes, projectScope{name: name, apply: apply})
}

// AddLocalScopeForProject assign a local scope to a model
func AddLocalScopeForProject(name string, apply func(builder query.Condition)) {
	projectLocalScopes = append(projectLocalScopes, projectScope{name: name, apply: apply})
}

func (m *ProjectModel) applyScope() query.Condition {
	scopeCond := query.ConditionBuilder()
	for _, g := range projectGlobalScopes {
		if m.globalScopeEnabled(g.name) {
			g.apply(scopeCond)
		}
	}

	for _, s := range projectLocalScopes {
		if m.localScopeEnabled(s.name) {
			s.apply(scopeCond)
		}
	}

	return scopeCond
}

func (m *ProjectModel) localScopeEnabled(name string) bool {
	for _, n := range m.includeLocalScopes {
		if name == n {
			return true
		}
	}

	return false
}

func (m *ProjectModel) globalScopeEnabled(name string) bool {
	for _, n := range m.excludeGlobalScopes {
		if name == n {
			return false
		}
	}

	return true
}

type projectWrap struct {
	Id               null.Int
	Name             null.String
	Description      null.String
	Visibility       null.Int
	UserId           null.Int
	SortLevel        null.Int
	CatalogId        null.Int
	CatalogFoldStyle null.Int
	CatalogSortStyle null.Int
	CreatedAt        null.Time
	UpdatedAt        null.Time
	DeletedAt        null.Time
}

func (w projectWrap) ToProject() Project {
	return Project{
		original: &projectOriginal{
			Id:               w.Id.Int64,
			Name:             w.Name.String,
			Description:      w.Description.String,
			Visibility:       int8(w.Visibility.Int64),
			UserId:           int(w.UserId.Int64),
			SortLevel:        int(w.SortLevel.Int64),
			CatalogId:        int(w.CatalogId.Int64),
			CatalogFoldStyle: int8(w.CatalogFoldStyle.Int64),
			CatalogSortStyle: int8(w.CatalogSortStyle.Int64),
			CreatedAt:        w.CreatedAt.Time,
			UpdatedAt:        w.UpdatedAt.Time,
			DeletedAt:        w.DeletedAt,
		},

		Id:               w.Id.Int64,
		Name:             w.Name.String,
		Description:      w.Description.String,
		Visibility:       int8(w.Visibility.Int64),
		UserId:           int(w.UserId.Int64),
		SortLevel:        int(w.SortLevel.Int64),
		CatalogId:        int(w.CatalogId.Int64),
		CatalogFoldStyle: int8(w.CatalogFoldStyle.Int64),
		CatalogSortStyle: int8(w.CatalogSortStyle.Int64),
		CreatedAt:        w.CreatedAt.Time,
		UpdatedAt:        w.UpdatedAt.Time,
		DeletedAt:        w.DeletedAt,
	}
}

// ProjectModel is a model which encapsulates the operations of the object
type ProjectModel struct {
	db        *query.DatabaseWrap
	tableName string

	excludeGlobalScopes []string
	includeLocalScopes  []string

	query query.SQLBuilder
}

var projectTableName = "wz_projects"

func SetProjectTable(tableName string) {
	projectTableName = tableName
}

// NewProjectModel create a ProjectModel
func NewProjectModel(db query.Database) *ProjectModel {
	return &ProjectModel{
		db:                  query.NewDatabaseWrap(db),
		tableName:           projectTableName,
		excludeGlobalScopes: make([]string, 0),
		includeLocalScopes:  make([]string, 0),
		query:               query.Builder(),
	}
}

// GetDB return database instance
func (m *ProjectModel) GetDB() query.Database {
	return m.db.GetDB()
}

// WithTrashed force soft deleted models to appear in a result set
func (m *ProjectModel) WithTrashed() *ProjectModel {
	return m.WithoutGlobalScopes("soft_delete")
}

func (m *ProjectModel) clone() *ProjectModel {
	return &ProjectModel{
		db:                  m.db,
		tableName:           m.tableName,
		excludeGlobalScopes: append([]string{}, m.excludeGlobalScopes...),
		includeLocalScopes:  append([]string{}, m.includeLocalScopes...),
		query:               m.query,
	}
}

// WithoutGlobalScopes remove a global scope for given query
func (m *ProjectModel) WithoutGlobalScopes(names ...string) *ProjectModel {
	mc := m.clone()
	mc.excludeGlobalScopes = append(mc.excludeGlobalScopes, names...)

	return mc
}

// WithLocalScopes add a local scope for given query
func (m *ProjectModel) WithLocalScopes(names ...string) *ProjectModel {
	mc := m.clone()
	mc.includeLocalScopes = append(mc.includeLocalScopes, names...)

	return mc
}

// Condition add query builder to model
func (m *ProjectModel) Condition(builder query.SQLBuilder) *ProjectModel {
	mm := m.clone()
	mm.query = mm.query.Merge(builder)

	return mm
}

// Find retrieve a model by its primary key
func (m *ProjectModel) Find(id int64) (Project, error) {
	return m.First(m.query.Where("id", "=", id))
}

// Exists return whether the records exists for a given query
func (m *ProjectModel) Exists(builders ...query.SQLBuilder) (bool, error) {
	count, err := m.Count(builders...)
	return count > 0, err
}

// Count return model count for a given query
func (m *ProjectModel) Count(builders ...query.SQLBuilder) (int64, error) {
	sqlStr, params := m.query.
		Merge(builders...).
		Table(m.tableName).
		AppendCondition(m.applyScope()).
		ResolveCount()

	rows, err := m.db.QueryContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	rows.Next()
	var res int64
	if err := rows.Scan(&res); err != nil {
		return 0, err
	}

	return res, nil
}

func (m *ProjectModel) Paginate(page int64, perPage int64, builders ...query.SQLBuilder) ([]Project, query.PaginateMeta, error) {
	if page <= 0 {
		page = 1
	}

	if perPage <= 0 {
		perPage = 15
	}

	meta := query.PaginateMeta{
		PerPage: perPage,
		Page:    page,
	}

	count, err := m.Count(builders...)
	if err != nil {
		return nil, meta, err
	}

	meta.Total = count
	meta.LastPage = count / perPage
	if count%perPage != 0 {
		meta.LastPage += 1
	}

	res, err := m.Get(append([]query.SQLBuilder{query.Builder().Limit(perPage).Offset((page - 1) * perPage)}, builders...)...)
	if err != nil {
		return res, meta, err
	}

	return res, meta, nil
}

// Get retrieve all results for given query
func (m *ProjectModel) Get(builders ...query.SQLBuilder) ([]Project, error) {
	b := m.query.Merge(builders...).Table(m.tableName).AppendCondition(m.applyScope())
	if len(b.GetFields()) == 0 {
		b = b.Select(
			"id",
			"name",
			"description",
			"visibility",
			"user_id",
			"sort_level",
			"catalog_id",
			"catalog_fold_style",
			"catalog_sort_style",
			"created_at",
			"updated_at",
			"deleted_at",
		)
	}

	fields := b.GetFields()
	selectFields := make([]query.Expr, 0)

	for _, f := range fields {
		switch strcase.ToSnake(f.Value) {

		case "id":
			selectFields = append(selectFields, f)
		case "name":
			selectFields = append(selectFields, f)
		case "description":
			selectFields = append(selectFields, f)
		case "visibility":
			selectFields = append(selectFields, f)
		case "user_id":
			selectFields = append(selectFields, f)
		case "sort_level":
			selectFields = append(selectFields, f)
		case "catalog_id":
			selectFields = append(selectFields, f)
		case "catalog_fold_style":
			selectFields = append(selectFields, f)
		case "catalog_sort_style":
			selectFields = append(selectFields, f)
		case "created_at":
			selectFields = append(selectFields, f)
		case "updated_at":
			selectFields = append(selectFields, f)
		case "deleted_at":
			selectFields = append(selectFields, f)
		}
	}

	var createScanVar = func(fields []query.Expr) (*projectWrap, []interface{}) {
		var projectVar projectWrap
		scanFields := make([]interface{}, 0)

		for _, f := range fields {
			switch strcase.ToSnake(f.Value) {

			case "id":
				scanFields = append(scanFields, &projectVar.Id)
			case "name":
				scanFields = append(scanFields, &projectVar.Name)
			case "description":
				scanFields = append(scanFields, &projectVar.Description)
			case "visibility":
				scanFields = append(scanFields, &projectVar.Visibility)
			case "user_id":
				scanFields = append(scanFields, &projectVar.UserId)
			case "sort_level":
				scanFields = append(scanFields, &projectVar.SortLevel)
			case "catalog_id":
				scanFields = append(scanFields, &projectVar.CatalogId)
			case "catalog_fold_style":
				scanFields = append(scanFields, &projectVar.CatalogFoldStyle)
			case "catalog_sort_style":
				scanFields = append(scanFields, &projectVar.CatalogSortStyle)
			case "created_at":
				scanFields = append(scanFields, &projectVar.CreatedAt)
			case "updated_at":
				scanFields = append(scanFields, &projectVar.UpdatedAt)
			case "deleted_at":
				scanFields = append(scanFields, &projectVar.DeletedAt)
			}
		}

		return &projectVar, scanFields
	}

	sqlStr, params := b.Fields(selectFields...).ResolveQuery()

	rows, err := m.db.QueryContext(context.Background(), sqlStr, params...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	projects := make([]Project, 0)
	for rows.Next() {
		projectVar, scanFields := createScanVar(fields)
		if err := rows.Scan(scanFields...); err != nil {
			return nil, err
		}

		projectReal := projectVar.ToProject()
		projectReal.SetModel(m)
		projects = append(projects, projectReal)
	}

	return projects, nil
}

// First return first result for given query
func (m *ProjectModel) First(builders ...query.SQLBuilder) (Project, error) {
	res, err := m.Get(append(builders, query.Builder().Limit(1))...)
	if err != nil {
		return Project{}, err
	}

	if len(res) == 0 {
		return Project{}, query.ErrNoResult
	}

	return res[0], nil
}

// Create save a new project to database
func (m *ProjectModel) Create(kv query.KV) (int64, error) {

	if _, ok := kv["created_at"]; !ok {
		kv["created_at"] = time.Now()
	}

	if _, ok := kv["updated_at"]; !ok {
		kv["updated_at"] = time.Now()
	}

	sqlStr, params := m.query.Table(m.tableName).ResolveInsert(kv)

	res, err := m.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// SaveAll save all projects to database
func (m *ProjectModel) SaveAll(projects []Project) ([]int64, error) {
	ids := make([]int64, 0)
	for _, project := range projects {
		id, err := m.Save(project)
		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

// Save save a project to database
func (m *ProjectModel) Save(project Project) (int64, error) {
	return m.Create(project.StaledKV())
}

// SaveOrUpdate save a new project or update it when it has a id > 0
func (m *ProjectModel) SaveOrUpdate(project Project) (id int64, updated bool, err error) {
	if project.Id > 0 {
		_, _err := m.UpdateById(project.Id, project)
		return project.Id, true, _err
	}

	_id, _err := m.Save(project)
	return _id, false, _err
}

// UpdateFields update kv for a given query
func (m *ProjectModel) UpdateFields(kv query.KV, builders ...query.SQLBuilder) (int64, error) {
	if len(kv) == 0 {
		return 0, nil
	}

	kv["updated_at"] = time.Now()

	sqlStr, params := m.query.Merge(builders...).AppendCondition(m.applyScope()).
		Table(m.tableName).
		ResolveUpdate(kv)

	res, err := m.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Update update a model for given query
func (m *ProjectModel) Update(project Project, builders ...query.SQLBuilder) (int64, error) {
	return m.UpdateFields(project.StaledKV(), builders...)
}

// UpdateById update a model by id
func (m *ProjectModel) UpdateById(id int64, project Project) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Update(project)
}

// ForceDelete permanently remove a soft deleted model from the database
func (m *ProjectModel) ForceDelete(builders ...query.SQLBuilder) (int64, error) {
	m2 := m.WithTrashed()

	sqlStr, params := m2.query.Merge(builders...).AppendCondition(m2.applyScope()).Table(m2.tableName).ResolveDelete()

	res, err := m2.db.ExecContext(context.Background(), sqlStr, params...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// ForceDeleteById permanently remove a soft deleted model from the database by id
func (m *ProjectModel) ForceDeleteById(id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).ForceDelete()
}

// Restore restore a soft deleted model into an active state
func (m *ProjectModel) Restore(builders ...query.SQLBuilder) (int64, error) {
	m2 := m.WithTrashed()
	return m2.UpdateFields(query.KV{
		"deleted_at": nil,
	}, builders...)
}

// RestoreById restore a soft deleted model into an active state by id
func (m *ProjectModel) RestoreById(id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Restore()
}

// Delete remove a model
func (m *ProjectModel) Delete(builders ...query.SQLBuilder) (int64, error) {

	return m.UpdateFields(query.KV{
		"deleted_at": time.Now(),
	}, builders...)

}

// DeleteById remove a model by id
func (m *ProjectModel) DeleteById(id int64) (int64, error) {
	return m.Condition(query.Builder().Where("id", "=", id)).Delete()
}