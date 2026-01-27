package amlich

var (
	Can             = [10]string{"Giáp", "Ất", "Bính", "Đinh", "Mậu", "Kỷ", "Canh", "Tân", "Nhâm", "Quý"}
	Chi             = [12]string{"Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ", "Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi"}
	DaysOfWeek      = [7]string{"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"}
	VietnamTimezone = "Asia/Ho_Chi_Minh"

	HourNames = [12]string{
		"Tý (23-01h)",
		"Sửu (01-03h)",
		"Dần (03-05h)",
		"Mão (05-07h)",
		"Thìn (07-09h)",
		"Tỵ (09-11h)",
		"Ngọ (11-13h)",
		"Mùi (13-15h)",
		"Thân (15-17h)",
		"Dậu (17-19h)",
		"Tuất (19-21h)",
		"Hợi (21-23h)",
	}

	AuspiciousHours = map[int][]int{
		0:  {0, 1, 4, 6, 8, 9},
		1:  {2, 3, 5, 8, 10, 11},
		2:  {0, 1, 4, 5, 7, 10},
		3:  {0, 2, 3, 6, 7, 9},
		4:  {2, 4, 5, 8, 9, 11},
		5:  {1, 4, 6, 7, 10, 11},
		6:  {0, 1, 4, 6, 8, 9},
		7:  {2, 3, 5, 8, 10, 11},
		8:  {0, 1, 4, 5, 7, 10},
		9:  {0, 2, 3, 6, 7, 9},
		10: {2, 4, 5, 8, 9, 11},
		11: {1, 4, 6, 7, 10, 11},
	}
)
