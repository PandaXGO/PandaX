package jobs

import (
	"log"
)

type CronDeviceHandle struct {
}

func (t CronDeviceHandle) Exec(arg any, content any) error {
	log.Println("执行设备任务", arg, content)

	return nil
}

type CronProductHandle struct {
}

func (t CronProductHandle) Exec(arg any, content any) error {
	log.Println("执行产品任务", arg, content)

	return nil
}
