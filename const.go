package amlich

type DiaChi int

const (
	// 12 địa chi
	Ty   DiaChi = iota // 0
	Suu                // 1
	Dan                // 2
	Mao                // 3
	Thin               // 4
	Ti                 // 5
	Ngo                // 6
	Mui                // 7
	Than               // 8
	Dau                // 9
	Tuat               // 10
	Hoi                // 11
)

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

	// [ngày] -> giờ
	ThoTuHour = map[DiaChi]DiaChi{
		Ty:   Suu,
		Suu:  Ngo,
		Dan:  Ngo,
		Mao:  Ti,
		Thin: Ti,
		Ti:   Ngo,
		Ngo:  Mui,
		Mui:  Ngo,
		Than: Mao,
		Dau:  Ti,
		Tuat: Mui,
		Hoi:  Ngo,
	}

	// [tháng] -> giờ
	SatChuHour = map[int]DiaChi{
		1:  Dan,
		2:  Ti,
		3:  Than,
		4:  Thin,
		5:  Dau,
		6:  Mao,
		7:  Dan,
		8:  Ti,
		9:  Than,
		10: Thin,
		11: Dau,
		12: Mao,
	}
)

func (c DiaChi) String() string {
	switch c {
	case Ty:
		return "Tý"
	case Suu:
		return "Sửu"
	case Dan:
		return "Dần"
	case Mao:
		return "Mão"
	case Thin:
		return "Thìn"
	case Ti:
		return "Tị"
	case Ngo:
		return "Ngọ"
	case Mui:
		return "Mùi"
	case Than:
		return "Thân"
	case Dau:
		return "Dậu"
	case Tuat:
		return "Tuất"
	case Hoi:
		return "Hợi"
	default:
		return "Không hợp lệ"
	}
}
