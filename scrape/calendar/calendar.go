package calendar

import (
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/uesteibar/phenomena_calendar_scraper/scrape/phenomena"
)

func CreateICS(months []phenomena.Month) string {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	for _, month := range months {
		for rawDate, day := range month {
			for _, scheduling := range day {
				rawDatetime := fmt.Sprintf("%s %s", rawDate, scheduling.Time)
				startsAt, _ := time.Parse("2006-01-02 15:04h", rawDatetime)
				endsAt := startsAt.Add(time.Minute * time.Duration(scheduling.Duration))

				event := cal.AddEvent(fmt.Sprintf("%s@phenomena_calendar", rawDatetime))
				event.SetStartAt(startsAt)
				event.SetEndAt(endsAt)
				event.SetCreatedTime(time.Now())
				event.SetDtStampTime(time.Now())
				event.SetModifiedAt(time.Now())
				event.SetSummary(scheduling.Title)
				event.SetLocation(scheduling.Url)
			}
		}
	}

	return cal.Serialize()
}
