package daysteps

import (
	"errors"
	"fmt"
	spcs "github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// var myLog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

// parsePackage
//
// Парсинг строки вида "678,0h50m"
func parsePackage(data string) (int, time.Duration, error) {
	strs := strings.Split(data, ",")
	if len(strs) != 2 {
		return 0, 0, errors.New("(parsePackage) bad income data")
	}

	steps, walkDur := strs[0], strs[1]

	if steps == "" || walkDur == "" {
		return 0, 0, errors.New("(parsePackage) no income required data")
	}

	// steps
	stepsCnt, err := strconv.Atoi(steps)
	if err != nil || stepsCnt <= 0 {
		return 0, 0, fmt.Errorf("(parsePackage) bad income steps data: %v", err)
	}

	// duration of walk
	duration, err := time.ParseDuration(walkDur)
	if err != nil || duration <= 0 {
		return 0, 0, fmt.Errorf("(parsePackage) bad income duration data: %v", err)
	}

	return stepsCnt, duration, nil
}

// DayActionInfo
//
//	data — строка с данными, содержащими количество шагов и продолжительность прогулки в формате 3h50m,
//	weight — вес пользователя в килограммах,
//	height — рост пользователя в метрах.
func DayActionInfo(data string, weight, height float64) string {
	if data == "" || weight <= 0 || height <= 0 {
		log.Println("(DayActionInfo) Bad input data")
		return ""
	}

	steps, walkDur, err := parsePackage(data)
	if err != nil {
		log.Println("(DayActionInfo) bad parsePackage calculate:", err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distKm := float64(steps) * stepLength / float64(mInKm)

	calories, err := spcs.WalkingSpentCalories(steps, weight, height, walkDur)
	if err != nil {
		log.Println("(DayActionInfo) bad WalkingSpentCalories calculate:", err)
		return ""
	}

	output := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distKm, calories)

	return output
}
