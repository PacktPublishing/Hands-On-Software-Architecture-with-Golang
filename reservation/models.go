package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Status int

const (
	Initial     Status = 0
	BookingMade Status = 1
	EmailSent   Status = 2
)

type HotelReservation struct {
	//HotelReservationDetails `gorm:"embedded"`
	CustomerId        uint      `json:"customer_id" `
	PaymentIdentifier uint      `json:"payment_identifier" `
	SKU               uint      `json:"entity_id" `
	RoomId            uint      `json:"room_id" `
	CheckIn           time.Time `json:"check_in" gorm:"type:datetime"`
	CheckOut          time.Time `json:"check_out" gorm:"type:datetime"`
	Id                string
	Status            Status
}

type AvailabilityThreshold struct {
	gorm.Model
	SKU          uint
	RoomId       uint
	Availability int
}

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/hotels?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	//db.AutoMigrate(&HotelReservation{})
	db.AutoMigrate(&AvailabilityThreshold{})

	// dummy thresholds
	db.Create(&AvailabilityThreshold{SKU: 1, RoomId: 2, Availability: 3})

}

// generates ID for the reservation from SKU, roomid and checkin date
func makeId(res *HotelReservationDTO) string {
	// NOTE : for uniquess, non-overlapping reservations, there should be another explicit check
	return fmt.Sprintf("%v#%v#%v", res.SKU, res.RoomId, res.CheckIn)
}

func persistReservation(res *HotelReservationDTO) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	//TODO : Check that there is no overlapping reservation

	if err := tx.Create(&HotelReservation{
		CustomerId:        res.CustomerId,
		PaymentIdentifier: res.PaymentIdentifier,
		SKU:               res.SKU,
		RoomId:            res.RoomId,
		CheckIn:           time.Time(res.CheckIn),
		CheckOut:          time.Time(res.CheckOut),
		Id:                makeId(res),
		Status:            Initial}).Error; err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("created hotel reservation..")

	// update the entry for availability threshold
	var threshold AvailabilityThreshold
	tx.Where("entity_id = ? AND room_id = ?", res.SKU, res.RoomId).First(&threshold)

	fmt.Printf("\nthreshold = %+v\n", threshold)
	tx.Model(&threshold).Where("id = ?", threshold.ID).Update("availability", threshold.Availability-1)

	// NOTE : availability is just a threshold for update here.
	// Even if availability is 0, reservation is forwarded to the Seller
	// And availability >0 in thresholds DB is not a guarantee of reservation certainity.
	if threshold.Availability <= 1 {
		// we have reached threshold
		sendInvaliationMessageToPriceStore(threshold.SKU, threshold.RoomId)

	}

	return tx.Commit().Error
}

func sendInvaliationMessageToPriceStore(eid, rid uint) {
	fmt.Println("sending message to invalide catalog for entity id ", eid, " room id ", rid)
}

func updateReservationStatus(reservationDTO *HotelReservationDTO, status Status) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}

	tx.Model(&HotelReservation{}).Where("id = ?", makeId(reservationDTO)).Update("status", status)

	return tx.Commit().Error
}
