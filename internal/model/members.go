package model

type Members struct {
	Member_Info string 	`json:"member_info,omitempty"`
	Teacher_Info string	`json:"teacher_info,omitempty"`
}
var Map_teacher = map[int]Members{
	1: {Teacher_Info: "Name: Mahmoud, Age: 32"},
}
var Map_members = map[int]Members{
	1: {Member_Info: "Name: Aliasghar, Age: 18"},
	2: {Member_Info: "Name: Shahrzad, Age: 17"},
	3: {Member_Info: "Name: AmirHossein, Age: 21"},
}