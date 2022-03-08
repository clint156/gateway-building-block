package model

//LaundryRoom represents the basic information returned as part of requesting and organization
type LaundryRoom struct {
	ID               int
	Name             string
	Status           string
	AvialableWashers int
	AvailableDryers  int
}

//Organization represents the top most level of inforomation provided by teh laundry api
type Organization struct {
	SchoolName   string
	LaundryRooms []*LaundryRoom
}

//RoomDetail represents details about a specific laundry room, including a list of appliances
type RoomDetail struct {
	NumWashers int
	NumDryers  int
	Appliances []*Appliance
}

//Appliance represents the information specific to an identifiiable appliance in a laundry room
type Appliance struct {
	ID               int
	Status           string
	Name             string
	AverageCycleTime int
	TimeRemaining    string
	OutofService     bool
}

//MachineRequestDetail represents the basic details needed in order to submit a request about a machine
type MachineRequestDetail struct {
	MachineID           string
	Message             string
	RecentServiceStatus bool
	ProblemCodes        []string
}

//ServiceRequestResult represents the information returned upon submission of a machine service request
type ServiceRequestResult struct {
	Message       string
	RequestNumber int
	Status        string
}
