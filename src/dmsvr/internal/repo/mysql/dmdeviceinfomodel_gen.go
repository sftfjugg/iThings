// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	dmDeviceInfoFieldNames          = builder.RawFieldNames(&DmDeviceInfo{})
	dmDeviceInfoRows                = strings.Join(dmDeviceInfoFieldNames, ",")
	dmDeviceInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(dmDeviceInfoFieldNames, "`id`", "`createdTime`", "`deletedTime`", "`updatedTime`"), ",")
	dmDeviceInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(dmDeviceInfoFieldNames, "`id`", "`createdTime`", "`deletedTime`", "`updatedTime`"), "=?,") + "=?"
)

type (
	dmDeviceInfoModel interface {
		Insert(ctx context.Context, data *DmDeviceInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DmDeviceInfo, error)
		FindOneByProductIDDeviceName(ctx context.Context, productID string, deviceName string) (*DmDeviceInfo, error)
		Update(ctx context.Context, data *DmDeviceInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDmDeviceInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	DmDeviceInfo struct {
		Id          int64        `db:"id"`
		ProductID   string       `db:"productID"`  // 产品id
		DeviceName  string       `db:"deviceName"` // 设备名称
		Secret      string       `db:"secret"`     // 设备秘钥
		Cert        string       `db:"cert"`       // 设备证书
		Imei        string       `db:"imei"`       // IMEI号信息
		Mac         string       `db:"mac"`        // MAC号信息
		Version     string       `db:"version"`    // 固件版本
		HardInfo    string       `db:"hardInfo"`   // 模组硬件型号
		SoftInfo    string       `db:"softInfo"`   // 模组软件版本
		Position    string       `db:"position"`   // 设备的位置(默认百度坐标系BD09)
		Address     string       `db:"address"`    // 所在地址
		Tags        string       `db:"tags"`       // 设备标签
		IsOnline    int64        `db:"isOnline"`   // 是否在线,1是2否
		FirstLogin  sql.NullTime `db:"firstLogin"` // 激活时间
		LastLogin   sql.NullTime `db:"lastLogin"`  // 最后上线时间
		LogLevel    int64        `db:"logLevel"`   // 日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试
		CreatedTime time.Time    `db:"createdTime"`
		UpdatedTime time.Time    `db:"updatedTime"`
		DeletedTime sql.NullTime `db:"deletedTime"`
	}
)

func newDmDeviceInfoModel(conn sqlx.SqlConn) *defaultDmDeviceInfoModel {
	return &defaultDmDeviceInfoModel{
		conn:  conn,
		table: "`dm_device_info`",
	}
}

func (m *defaultDmDeviceInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultDmDeviceInfoModel) FindOne(ctx context.Context, id int64) (*DmDeviceInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dmDeviceInfoRows, m.table)
	var resp DmDeviceInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDmDeviceInfoModel) FindOneByProductIDDeviceName(ctx context.Context, productID string, deviceName string) (*DmDeviceInfo, error) {
	var resp DmDeviceInfo
	query := fmt.Sprintf("select %s from %s where `productID` = ? and `deviceName` = ? limit 1", dmDeviceInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, productID, deviceName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDmDeviceInfoModel) Insert(ctx context.Context, data *DmDeviceInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, dmDeviceInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductID, data.DeviceName, data.Secret, data.Cert, data.Imei, data.Mac, data.Version, data.HardInfo, data.SoftInfo, data.Position, data.Address, data.Tags, data.IsOnline, data.FirstLogin, data.LastLogin, data.LogLevel)
	return ret, err
}

func (m *defaultDmDeviceInfoModel) Update(ctx context.Context, newData *DmDeviceInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, dmDeviceInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.ProductID, newData.DeviceName, newData.Secret, newData.Cert, newData.Imei, newData.Mac, newData.Version, newData.HardInfo, newData.SoftInfo, newData.Position, newData.Address, newData.Tags, newData.IsOnline, newData.FirstLogin, newData.LastLogin, newData.LogLevel, newData.Id)
	return err
}

func (m *defaultDmDeviceInfoModel) tableName() string {
	return m.table
}
