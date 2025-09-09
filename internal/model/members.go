package model

type Members struct {
	Member_Info  string `json:"member_info,omitempty"`
	Teacher_Info string `json:"teacher_info,omitempty"`
	Contact_Info string `json:"contact_info,omitempty"`
	Website_Info string `json:"website_info,omitempty"`
}

var Map_teacher = map[int]Members{
	1: {Teacher_Info: "Name: Mahmoud, Age: 32"},
}
var Map_members = map[int]Members{
	1: {Member_Info: "Name: Aliasghar, Age: 18"},
	2: {Member_Info: "Name: Shahrzad, Age: 17"},
	3: {Member_Info: "Name: AmirHossein, Age: 21"},
}
var Map_contact = map[int]Members{
	1: {Contact_Info: "Email: something@example.com, Phone: 09123456789"},
}

var Map_website = map[int]Members{
	1: {Website_Info: "https://ArvanSchool.ir"},
}