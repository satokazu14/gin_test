package service

import (
	"gin_test/db"
	"gin_test/entity"

	"github.com/gin-gonic/gin"
)

type CarService struct{}

type Car entity.Car

//車両登録
func (s CarService) CreateModel(c *gin.Context) (Car, error) {
	db := db.GetDB()
	var car Car

	if err := c.BindJSON(&car); err != nil {
		return car, err
	}

	db.Exec("INSERT INTO cars (name, grade, manufacture, model_year, model_number, body_type, door_num, color, drive_system, transmission, displacement, gas_oil, seating_capacity, ex_point, in_point, mileage, one_owner, repair_history, non_smoking, auto_inspection, certified_used, recycling_consignment, leagal_maintenance, warranty, purchase_price, description, pw_steering, pw_window, aircon, w_aircon, keyless, smart_key, car_nabigation, tv, audio, visual, bluetooth, usb, pw_supply, back_camera, around_camera, etc, third_seat, ele_seat, seat_heater, seat_aircon, leather_seat, slide_door, ele_gate, walk_through, cruise_control, lane_assist, obstacle_sensor, auto_parking, parking_assist, abs, stability_control, mitigation_brake, head_light, sun_roof, aero, al_wheel, low_down, lift_up, super_charger, air_suspension,image_num) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", car.Name, car.Grade, car.Manufacture, car.ModelYear, car.ModelNumber, car.BodyType, car.DoorNum, car.Color, car.DriveSystem, car.Transmission, car.Displacement, car.GasOil, car.SeatingCapacity, car.ExPoint, car.InPoint, car.Mileage, car.OneOwner, car.RepairHistory, car.NonSmoking, car.AutoInspection, car.CertifiedUsed, car.RecyclingConsignment, car.LeagalMaintenance, car.Warranty, car.PurchasePrice, car.Description, car.PwSteering, car.PwWindow, car.Aircon, car.WAircon, car.Keyless, car.SmartKey, car.CarNabigation, car.Tv, car.Audio, car.Visual, car.Bluetooth, car.Usb, car.PwSupply, car.BackCamera, car.AroundCamera, car.Etc, car.ThirdSeat, car.EleSeat, car.SeatHeater, car.SeatAircon, car.LeatherSeat, car.SlideDoor, car.EleGate, car.WalkThrough, car.CruiseControl, car.LaneAssist, car.ObstacleSensor, car.AutoParking, car.ParkingAssist, car.Abs, car.StabilityControl, car.MitigationBrake, car.HeadLight, car.SunRoof, car.Aero, car.AlWheel, car.LowDown, car.LiftUp, car.SuperCharger, car.AirSuspension, car.ImageNum)

	return car, nil
}

//車両全件取得
func (s CarService) GetAll() ([]Car, error) {
	db := db.GetDB()
	var car []Car

	if err := db.Find(&car).Error; err != nil {
		return nil, err
	}

	return car, nil
}

//車両指定件数取得
func (s CarService) GetSum(limit string) ([]Car, error) {
	db := db.GetDB()
	var car []Car

	if err := db.Limit(limit).Find(&car).Error; err != nil {
		return car, err
	}

	return car, nil
}

//IDで指定した車両を取得
func (s CarService) GetByID(carId string) (Car, error) {
	db := db.GetDB()
	var car Car

	if err := db.Where("id = ?", carId).First(&car).Error; err != nil {
		return car, err
	}

	return car, nil
}

//IDで指定した車両を更新
func (s CarService) UpdateByID(carId string, c *gin.Context) (Car, error) {
	db := db.GetDB()
	var car Car

	if err := db.Where("id = ?", carId).First(&car).Error; err != nil {
		return car, err
	}

	if err := c.BindJSON(&car); err != nil {
		return car, err
	}

	db.Save(&car)

	return car, nil
}

//IDで指定した車両を削除
func (s CarService) DeleteByID(carId string) error {
	db := db.GetDB()
	var car Car

	if err := db.Where("id = ?", carId).Delete(&car).Error; err != nil {
		return err
	}

	return nil
}
