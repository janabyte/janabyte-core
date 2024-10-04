package model

type Club struct {
	Id            int    `json:"id"`
	Name          string `json:"name" validate:"required"`
	Description   string `json:"description"`
	Address       string `json:"address"`
	WorkTimeStart string `json:"work_time_start"`
	WorkTimeEnd   string `json:"work_time_end"`
	XSize         int    `json:"x_size"`
	YSize         int    `json:"y_size"`
	UserId        int    `json:"user_id" validate:"required"`
}
