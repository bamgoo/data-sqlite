# data-sqlite

`data-sqlite` 是 `data` 模块的 `sqlite` 驱动。

## 安装

```bash
go get github.com/infrago/data@latest
go get github.com/infrago/data-sqlite@latest
```

## 接入

```go
import (
    _ "github.com/infrago/data"
    _ "github.com/infrago/data-sqlite"
    "github.com/infrago/infra"
)

func main() {
    infra.Run()
}
```

## 配置示例

```toml
[data]
driver = "sqlite"
```

## 公开 API（摘自源码）

- `func Driver() data.Driver`
- `func (d *sqliteDriver) Connect(inst *data.Instance) (data.Connection, error)`
- `func (c *sqliteConnection) Open() error`
- `func (c *sqliteConnection) Close() error`
- `func (c *sqliteConnection) Health() data.Health`
- `func (c *sqliteConnection) DB() *sql.DB`
- `func (c *sqliteConnection) Dialect() data.Dialect`
- `func (sqliteDialect) Name() string { return "sqlite" }`
- `func (sqliteDialect) Quote(s string) string`
- `func (sqliteDialect) Placeholder(_ int) string { return "?" }`
- `func (sqliteDialect) SupportsILike() bool      { return false }`
- `func (sqliteDialect) SupportsReturning() bool  { return false }`

## 排错

- driver 未生效：确认模块段 `driver` 值与驱动名一致
- 连接失败：检查 endpoint/host/port/鉴权配置
