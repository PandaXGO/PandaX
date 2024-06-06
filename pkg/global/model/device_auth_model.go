package model

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
	"pandax/pkg/cache"
	"strconv"
	"strings"
)

type DeviceAuth struct {
	Owner          string `json:"owner"`
	OrgId          int64  `json:"orgId"`
	DeviceId       string `json:"deviceId"`
	DeviceType     string `json:"deviceType"`
	DeviceProtocol string `json:"deviceProtocol"`
	ProductId      string `json:"productId"`
	RuleChainId    string `json:"ruleChainId"`
	Name           string `json:"name"`
	CreatedAt      int64  `json:"created_at"`
	ExpiredAt      int64  `json:"expired_at"`
}

func (entity *DeviceAuth) GetDeviceToken(key string) error {
	if err := cache.GetDeviceEtoken(key, entity); err != nil {
		return err
	}
	return nil
}

func (token *DeviceAuth) MD5ID() string {
	buf := bytes.NewBufferString(token.DeviceId)
	buf.WriteString(token.DeviceType)
	buf.WriteString(strconv.FormatInt(token.CreatedAt, 10))
	access := base64.URLEncoding.EncodeToString([]byte(uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes()).String()))
	access = strings.TrimRight(access, "=")
	return access
}

func (token *DeviceAuth) GetMarshal() string {
	marshal, _ := json.Marshal(*token)
	return string(marshal)
}

func (token *DeviceAuth) GetUnMarshal(data []byte) error {
	return json.Unmarshal(data, token)
}

// 序列化
func (m *DeviceAuth) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// 反序列化
func (m *DeviceAuth) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func OrgAuthSet(tx *gorm.DB, roleId int64, owner string) error {
	if roleId == 0 {
		return errors.New("角色Id不能为空")
	}
	role, err := services.SysRoleModelDao.FindOrganizationsByRoleId(roleId)
	if err != nil {
		return err
	}
	if role.DataScope != entity.SELFDATASCOPE {
		if len(role.Org) <= 0 {
			return errors.New("该角色下未分配组织权限")
		}
		tx.Where("org_id in (?)", strings.Split(role.Org, ","))
	} else {
		tx.Where("owner = ?", owner)
	}

	return nil
}
