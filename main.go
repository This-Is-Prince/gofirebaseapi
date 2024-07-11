package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/This-Is-Prince/gofirebaseapi/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var db *firestore.Client

func init() {
	utils.LoadEnv()

	firestoreCredentialsJSON := utils.GetEnv("FIREBASE_SERVICE_ACC_CONFIG")

	_db, err := firestore.NewClient(context.Background(), utils.GetEnv("FIREBASE_SERVICE_ACC_CONFIG_PROJECT_ID"), option.WithCredentialsJSON([]byte(firestoreCredentialsJSON)))

	if err != nil {
		log.Fatalf("Failed to create Firestore client : %v\n", err)
	}
	log.Println("Firestore client created successfully.")
	db = _db
}

type CalendarEvent struct {
	Content   string    `json:"content"`
	EndTime   time.Time `json:"endTime"`
	StartTime time.Time `json:"startTime"`
	Status    string    `json:"status"`
	Id        string    `json:"id,omitempty"`
	Location  string    `json:"location"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	UserId    string    `json:"userId,omitempty"`
	IsOnline  bool      `json:"isOnline"`
	HouseId   string    `json:"houseId"`
}

type CalendarEventBody struct {
	Event *struct {
		Content   string `json:"content"`
		EndTime   string `json:"endTime"`
		StartTime string `json:"startTime"`
		Status    string `json:"status"`
		Id        string `json:"id,omitempty"`
		Location  string `json:"location"`
		Title     string `json:"title"`
		Url       string `json:"url"`
		UserId    string `json:"userId,omitempty"`
		IsOnline  bool   `json:"isOnline"`
		HouseId   string `json:"houseId"`
	} `json:"event,omitempty"`
}

func main() {
	fmt.Println("Starting server....")

	router := gin.Default()

	router.POST("/createCalendarEvent", func(c *gin.Context) {
		var body CalendarEventBody

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error invalid request body %s", err.Error())})
			return
		}

		_event := body.Event

		if _event == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Event can't be nil."})
			return
		}

		col := db.Collection("testing_gofirebaseapi")

		doc := col.NewDoc()

		userId := "random"
		eventId := doc.ID

		event := CalendarEvent{
			Content: _event.Content,
			EndTime: utils.ConvertStringDateIntoGolangDateTime(_event.EndTime),
			// EndTime:   _event.EndTime,
			StartTime: utils.ConvertStringDateIntoGolangDateTime(_event.StartTime),
			// StartTime: _event.StartTime,
			Status:   _event.Status,
			Id:       eventId,
			Location: _event.Location,
			Title:    _event.Title,
			Url:      _event.Url,
			UserId:   userId,
			IsOnline: _event.IsOnline,
			HouseId:  _event.HouseId,
		}

		_, err := doc.Set(context.Background(), utils.StructToMap(event))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error saving event to firestore %s", err.Error())})
			return
		}

		// save this into firestore

		c.JSON(200, gin.H{
			"event": event,
		})
	})

	router.Run(fmt.Sprintf(":%s", utils.GetEnv("PORT")))
}
