package main

import (
	"fmt"
	"log"
	"net/http"

	"2.12/internal/calendar"
	"2.12/internal/server"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error of reading config: %v", err)
	}
	return nil
}

func main() {
	cal := calendar.NewCalendar()
	srv := &server.Server{Calendar: cal}

	if err := initConfig(); err != nil {
		log.Fatal("config initialization error: ", err)
	}
	port := viper.GetString("server.port")

	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", srv.CreateEventHandler)
	mux.HandleFunc("/update_event", srv.UpdateEventHandler)
	mux.HandleFunc("/delete_event", srv.DeleteEventHandler)
	mux.HandleFunc("/events_for_day", srv.EventsForDayHandler)
	mux.HandleFunc("/events_for_week", srv.EventsForWeekHandler)
	mux.HandleFunc("/events_for_month", srv.EventsForMonthHandler)

	handler := server.LoggingMidleware(mux)

	log.Println("Server is running on the port:", port)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal("Server startup error: ", err)
	}
}
