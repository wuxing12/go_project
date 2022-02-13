package send

import (
	"context"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

//定义接收者的邮箱
const (
	recipient = "462118329@qq.com"
)

type SendService struct {
}

func (this *SendService) Send(ctx context.Context, in *SendReq) (*SendRsp, error) {
	fmt.Printf("时间戳:%d，性能:%s，维度:%s，值:%s，告警类型: %s ",
		in.Timestamp, in.Metric, in.Dimensions, in.Value, in.AlertType)
	//根据告警类型发送相对应的邮件
	if in.AlertType == "SEVERE" {
		SevereSendMail(recipient, in.Timestamp, in.Metric, in.Dimensions, in.Value, in.AlertType)
	}
	if in.AlertType == "FATAL" {
		FatalSendMail(recipient, in.Timestamp, in.Metric, in.Dimensions, in.Value, in.AlertType)
	}
	if in.AlertType == "WARN" {
		return &SendRsp{Code: 0, Msg: "cpu使用过高警告"}, nil
	}
	return &SendRsp{Code: 1, Msg: "告警邮件发送成功！！！"}, nil
}

func SevereSendMail(Receiver string, Time int64, metric string, dim map[string]string, value float64, Alter string) (feedback string) {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "CPU监控提醒 <462118329@qq.com>"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{Receiver}

	// 抄送
	em.Cc = []string{Receiver}

	// 密送
	em.Bcc = []string{Receiver}
	// 设置主题
	em.Subject = "您的CPU内存即将爆满了," + fmt.Sprintf("告警类型为 “%s” ", Alter)

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("您的CPU阈值快满了，请尽快查看！！各项指标分别为：\n" +
		fmt.Sprintf("Timestamp: %d \n Metric: %s \n Value: %v \n Dimensions: %+v \n Altertype: %s \n",
			Time, metric, value, dim, Alter))

	//em.Text = []byte("您的CPU阈值爆满了，请请马上上线查看！！")
	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "462118329@qq.com", "jupwhpzgdoqnbhba", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Severe send successfully ... ")
	return "Send Completely"
}

func FatalSendMail(Receiver string, Time int64, metric string, dim map[string]string, value float64, Alter string) (feedback string) {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "CPU监控紧急通知！！！ <462118329@qq.com>"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{Receiver}

	// 抄送
	em.Cc = []string{Receiver}

	// 密送
	em.Bcc = []string{Receiver}

	// 设置主题
	em.Subject = "急！！！您的CPU使用率过高," + fmt.Sprintf("告警类型为 “%s” ", Alter)

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = []byte("您的CPU阈值爆满了，请马上上线修复！！！各项指标分别为：\n" +
		fmt.Sprintf("Timestamp: %d \n Metric: %s \n Value: %v \n Dimensions: %s \n Altertype: %s \n",
			Time, metric, value, dim, Alter))

	//em.Text = []byte("您的CPU阈值爆满了，请请马上上线查看！！")
	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "462118329@qq.com", "jupwhpzgdoqnbhba", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Fatal send successfully ... ")
	return "Send Completely"
}
