package entity

type Car struct {
	ID                   uint    `json:"id"`
	Name                 string  `json:"name"`
	Grade                string  `json:"grade"`
	Manufacture          string  `json:"manufacture"`
	ModelYear            string  `json:"model_year"`
	ModelNumber          string  `json:"model_number"`
	BodyType             string  `json:"body_type"`
	DoorNum              int     `json:"door_num"`
	Color                string  `json:"color"`
	DriveSystem          string  `json:"drive_system"`
	Transmission         string  `json:"transmission"`
	Displacement         int     `json:"displacement"`
	GasOil               string  `json:"gas_oil"`
	SeatingCapacity      int     `json:"seating_capacity"`
	ExPoint              int     `json:"ex_point"`
	InPoint              int     `json:"in_point"`
	Mileage              int     `json:"mileage"`
	OneOwner             int     `json:"one_owner"`
	RepairHistory        string  `json:"repair_history"`
	NonSmoking           int     `json:"non_smoking"`
	AutoInspection       string  `json:"auto_inspection"`
	CertifiedUsed        int     `json:"certified_used"`
	RecyclingConsignment string  `json:"recycling_consignment"`
	LeagalMaintenance    string  `json:"leagal_maintenance"`
	Warranty             string  `json:"warranty"`
	PurchasePrice        float64 `json:"purchase_price"`
	Description          string  `json:"description"`
	PwSteering           int     `json:"pw_steering"`
	PwWindow             int     `json:"pw_window"`
	Aircon               int     `json:"aircon"`
	WAircon              int     `json:"w_aircon"`
	Keyless              int     `json:"keyless"`
	SmartKey             int     `json:"smart_key"`
	CarNabigation        int     `json:"car_nabigation"`
	Tv                   string  `json:"tv"`
	Audio                string  `json:"audio"`
	Visual               string  `json:"visual"`
	Bluetooth            int     `json:"bluetooth"`
	Usb                  int     `json:"usb"`
	PwSupply             int     `json:"pw_supply"`
	BackCamera           int     `json:"back_camera"`
	AroundCamera         int     `json:"around_camera"`
	Etc                  int     `json:"etc"`
	ThirdSeat            int     `json:"third_seat"`
	EleSeat              int     `json:"ele_seat"`
	SeatHeater           int     `json:"seat_heater"`
	SeatAircon           int     `json:"seat_aircon"`
	LeatherSeat          int     `json:"leather_seat"`
	SlideDoor            int     `json:"slide_door"`
	EleGate              int     `json:"ele_gate"`
	WalkThrough          int     `json:"walk_through"`
	CruiseControl        int     `json:"cruise_control"`
	LaneAssist           int     `json:"lane_assist"`
	ObstacleSensor       int     `json:"obstacle_sensor"`
	AutoParking          int     `json:"auto_parking"`
	ParkingAssist        int     `json:"parking_assist"`
	Abs                  int     `json:"abs"`
	StabilityControl     int     `json:"stability_control"`
	MitigationBrake      int     `json:"mitigation_brake"`
	HeadLight            string  `json:"head_light"`
	SunRoof              int     `json:"sun_roof"`
	Aero                 int     `json:"aero"`
	AlWheel              int     `json:"al_wheel"`
	LowDown              int     `json:"low_down"`
	LiftUp               int     `json:"lift_up"`
	SuperCharger         int     `json:"super_charger"`
	AirSuspension        int     `json:"air_suspension"`
	StartTime            string  `json:"start_time"`
	ImageNum             int     `json:"image_num"`
	StartPrice           float64 `json:"start_price"`
}
