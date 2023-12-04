package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type TimeResponse struct {
	TorontoTime string `json:"toronto_time"`
}

func main() {
	http.HandleFunc("/current_time", timeHandler)
	http.ListenAndServe(":9090", nil)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	torontoTime := getCurrentTorontoTime()
	saveTimeToDatabase(torontoTime)

	response := TimeResponse{TorontoTime: torontoTime.Format(time.RFC3339)}
	json.NewEncoder(w).Encode(response)
}

func getCurrentTorontoTime() time.Time {
	loc, _ := time.LoadLocation("America/Toronto")
	return time.Now().In(loc)
}

func saveTimeToDatabase(time time.Time) {
	db, err := sql.Open("mysql", "harshdeep:Kaurharshdeep@123@/GOAPI")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", time)
	if err != nil {
		panic(err)
	}
}
func getTimeLogsFromDatabase() []TimeLog {
	db, err := sql.Open("mysql", "harshdeep:Kaurharshdeep@123@/GOAPI")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT timestamp FROM time_log ORDER BY timestamp DESC LIMIT 10")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var timeLogs []TimeLog
	for rows.Next() {
		var timeLog TimeLog
		err := rows.Scan(&timeLog.Timestamp)
		if err != nil {
			panic(err)
		}
		timeLogs = append(timeLogs, timeLog)
	}

	return timeLogs
}
