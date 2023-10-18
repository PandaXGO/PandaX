package global_model

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/rand"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
	"pandax/pkg/cache"
	"strconv"
	"strings"
	"time"
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

func OrgAuthSet(tx *gorm.DB, roleId int64, owner string) {
	if roleId == 0 {
		return
	}
	role, err := services.SysRoleModelDao.FindOrganizationsByRoleId(roleId)
	biz.ErrIsNil(err, "查询角色数据权限失败")
	if role.DataScope != entity.SELFDATASCOPE {
		biz.IsTrue(len(role.Org) > 0, "该角色下未分配组织权限")
		tx.Where("org_id in (?)", strings.Split(role.Org, ","))
	} else {
		tx.Where("owner = ?", owner)
	}
}
func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	id := make([]byte, 7) // 由于base64编码会增加字符数，这里使用7个字节生成10位ID
	_, err := rand.Read(id)
	if err != nil {
		panic(err) // 错误处理，根据实际情况进行处理
	}
	return base64.URLEncoding.EncodeToString(id)[:10]
}
