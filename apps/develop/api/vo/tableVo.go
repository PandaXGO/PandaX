package vo

import "pandax/apps/develop/entity"

/**
 * @Description
 * @Author 熊猫
 * @Date 2022/8/4 15:52
 **/

type TableInfoVo struct {
	List []entity.DevGenTableColumn `json:"list"`
	Info entity.DevGenTable         `json:"info"`
}
