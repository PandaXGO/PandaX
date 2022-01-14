package from

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/14 16:53
 **/

type SendMail struct {
	MailId  int64  `json:"mailId"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
