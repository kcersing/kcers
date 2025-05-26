package cron

import (
	"time"
)

func InitCron() {

	go func() {
		InitMinute1()
		InitMinute5()
		//InitMinute60()
	}()

}
func InitMinute1() {
	for {
		ticker := time.NewTicker(30 * time.Second)
		select {
		case <-ticker.C:
			SetMemberProductStatus()
			SetScheduleStatus()
			SetScheduleMemberStatus()
			SetScheduleCoachStatus()
			SetBootcampStatus()
			SetCommunityStatus()
			SetContestStatus()
			SetContestGroupsStatus()
			SetDeductMemberNumberSurplus()

			SetBootcampParticipantStatus()
			SetCommunityParticipantStatus()
			SetContestParticipantStatus()
			ticker.Stop()
		}
	}
}
func InitMinute5() {
	for {
		ticker := time.NewTicker(5 * time.Minute)
		select {
		case <-ticker.C:
			SetProductStatus()
			ticker.Stop()
		}
	}

}

//func InitMinute60() {
//	for {
//		ticker := time.NewTicker(60 * time.Minute)
//		select {
//		case <-ticker.C:
//
//			ticker.Stop()
//		}
//	}
//
//}
