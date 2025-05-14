package cron

import (
	"time"
)

func InitCron() {

	go func() {
		InitMinute1()
		InitMinute5()
	}()

}
func InitMinute1() {
	for {
		ticker := time.NewTicker(10 * time.Second)
		select {
		case <-ticker.C:
			SetScheduleStatus()
			SetScheduleMemberStatus()
			SetScheduleCoachStatus()
			ticker.Stop()
		}
	}
}
func InitMinute5() {
	for {
		ticker := time.NewTicker(1 * time.Minute)
		select {
		case <-ticker.C:
			SetMemberProductStatus()
			ticker.Stop()
		}
	}

}
