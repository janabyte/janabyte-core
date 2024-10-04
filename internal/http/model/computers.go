package model

type Computers struct {
	Id             int    `json:"id"`
	ComputerNumber int    `json:"computer_number"`
	IsNearToNext   bool   `json:"is_near_to_next"`
	IsNearToPrev   bool   `json:"is_near_to_prev"`
	Gpu            string `json:"gpu"`
	Cpu            string `json:"cpu"`
	Ram            string `json:"ram"`
	XPos           int    `json:"x_pos"`
	YPos           int    `json:"y_pos"`
	ClubId         int    `json:"club_id"`
	InstanceId     int    `json:"instance_id"`
}
