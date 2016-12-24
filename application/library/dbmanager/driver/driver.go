package driver

import (
	"github.com/admpub/nging/application/library/dbmanager/result"
	"github.com/webx-top/echo"
)

var (
	drivers       = map[string]Driver{}
	DefaultDriver = &BaseDriver{}
)

type Driver interface {
	Init(echo.Context, *DbAuth)
	Results() []result.Resulter
	AddResults(...result.Resulter) Driver
	SetResults(...result.Resulter) Driver
	SaveResults() Driver
	SavedResults() interface{}
	IsSupported(string) bool
	Login() error
	Logout() error
	ProcessList() error
	Privileges() error
	Info() error
	CreateDb() error
	ModifyDb() error
	ListDb() error
	CreateTable() error
	ModifyTable() error
	ListTable() error
	ViewTable() error
	ListData() error
	CreateData() error
	Indexes() error
	Foreign() error
	Trigger() error
	RunCommand() error
	Import() error
	Export() error
}

func NewBaseDriver() *BaseDriver {
	return &BaseDriver{}
}

type BaseDriver struct {
	echo.Context
	*DbAuth
	results []result.Resulter
}

func (m *BaseDriver) Results() []result.Resulter {
	return m.results
}

func (m *BaseDriver) AddResults(rs ...result.Resulter) Driver {
	if m.results == nil {
		m.results = []result.Resulter{}
	}
	m.results = append(m.results, rs...)
	return m
}

func (m *BaseDriver) SetResults(rs ...result.Resulter) Driver {
	m.results = rs
	return m
}

func (m *BaseDriver) SaveResults() Driver {
	if m.results == nil {
		return m
	}
	if v, y := m.Flash(`dbMgrResults`).([]result.Resulter); y {
		m.results = append(v, m.results...)
	}
	m.Session().AddFlash(m.results, `dbMgrResults`)
	return m
}

func (m *BaseDriver) SavedResults() interface{} {
	return m.Flash(`dbMgrResults`)
}

func (m *BaseDriver) Init(ctx echo.Context, auth *DbAuth) {
	m.Context = ctx
	m.DbAuth = auth
}
func (m *BaseDriver) IsSupported(operation string) bool {
	return true
}
func (m *BaseDriver) Login() error {
	return nil
}
func (m *BaseDriver) Logout() error {
	return nil
}
func (m *BaseDriver) ProcessList() error {
	return nil
}
func (m *BaseDriver) Privileges() error {
	return nil
}
func (m *BaseDriver) Info() error {
	return nil
}
func (m *BaseDriver) CreateDb() error {
	return nil
}
func (m *BaseDriver) ModifyDb() error {
	return nil
}
func (m *BaseDriver) ListDb() error {
	return nil
}
func (m *BaseDriver) CreateTable() error {
	return nil
}
func (m *BaseDriver) ModifyTable() error {
	return nil
}
func (m *BaseDriver) ListTable() error {
	return nil
}
func (m *BaseDriver) ViewTable() error {
	return nil
}
func (m *BaseDriver) ListData() error {
	return nil
}
func (m *BaseDriver) CreateData() error {
	return nil
}
func (m *BaseDriver) Indexes() error {
	return nil
}
func (m *BaseDriver) Foreign() error {
	return nil
}
func (m *BaseDriver) Trigger() error {
	return nil
}
func (m *BaseDriver) RunCommand() error {
	return nil
}
func (m *BaseDriver) Import() error {
	return nil
}
func (m *BaseDriver) Export() error {
	return nil
}

func Register(name string, driver Driver) {
	drivers[name] = driver
}

func Get(name string) (Driver, bool) {
	d, y := drivers[name]
	return d, y
}

func GetForce(name string) Driver {
	d, y := drivers[name]
	if !y {
		d = DefaultDriver
	}
	return d
}

func Has(name string) bool {
	_, y := drivers[name]
	return y
}

func GetAll() map[string]Driver {
	return drivers
}

func Unregister(name string) {
	_, y := drivers[name]
	if y {
		delete(drivers, name)
	}
}