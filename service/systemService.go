package service

import (
	"fmt"
	"github.com/martinlebeda/taskmaster/model"
	"github.com/spf13/viper"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const BaseDateTimeFormat = "2006-01-02 15:04"

func SysNotifyDistance(distance model.TimerDistance) {
	format := BaseDateTimeFormat
	t := time.Now()
	roundedToday := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	roundedGoal := time.Date(distance.Goal.Year(), distance.Goal.Month(), distance.Goal.Day(), 0, 0, 0, 0, t.Location())
	if roundedToday == roundedGoal {
		format = "15:04"
	}

	duration, _ := time.ParseDuration(strconv.Itoa(distance.Distance) + "s")

	msg := fmt.Sprintf("'%s - %v - %s'", duration.String(), distance.Goal.Format(format), distance.Note)

	notifycmd := viper.GetString("notifycmd")
	cmd := exec.Command(notifycmd, msg)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		log.Printf("combined out:\n%s\n", string(out))
	}
}

func SysAfterChange() {
	afterChangeCmd := viper.GetString("afterchange")
	fields := strings.Fields(afterChangeCmd)
	cmd := exec.Command(fields[0], fields[1:]...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Command finished with error: %v", err)
		log.Printf("combined out:\n%s\n", string(out))
	}
}
