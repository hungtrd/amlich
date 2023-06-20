# amlich

[![Build Status](https://github.com/hungtrd/amlich/workflows/Go/badge.svg?branch=master)](https://github.com/hungtrd/amlich/actions?query=branch%3Amaster)
[![Go Report Card](https://goreportcard.com/badge/github.com/hungtrd/amlich)](https://goreportcard.com/report/github.com/hungtrd/amlich)
[![Sourcegraph](https://sourcegraph.com/github.com/hungtrd/amlich/-/badge.svg)](https://sourcegraph.com/github.com/hungtrd/amlich?badge)
[![Release](https://img.shields.io/github/release/hungtrd/amlich.svg?style=flat-square)](https://github.com/hungtrd/amlich/releases)

Sử dụng để chuyển đổi từ ngày Dương Lịch sang ngày Âm Lịch và ngược lại

## Cài đặt package

```bash
go get github.com/hungtrd/amlich
```

## Sử dụng

### Khởi tạo

```go
lunar := amlich.New(time.Now().In(amlich.VietnamLocation()))

fmt.Printf("Âm lịch: ngày %02d, tháng %02d, năm %d\n", l.Day, l.Month, l.Year) // Âm lịch: ngày 03, tháng 05, năm 2023

// Can chi
fmt.Printf("Âm lịch: ngày %s, tháng %s, năm %s\n", l.DayAlias(), l.MonthAlias(), l.YearAlias()) // Âm lịch: ngày Kỷ Dậu, tháng Mậu Ngọ, năm Quý Mão

// Chuyển đổi
solar := lunar.ToSolar()
fmt.Println(solar.Weekday()) // Thứ Hai
fmt.Println(solar.String()) // Thứ Hai, ngày 20, tháng 06, năm 2023
```

### Sử dụng hàm convert trực tiếp

Đổi từ ngày Dương Lịch sang ngày Âm Lịch \
`leap` có giá trị `1` nếu tháng đó là tháng nhuận

```go
timeLoc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
today := time.Now().In(timeLoc)
tz := 7 // GMT +7
lunD, lunM, lunY, leap := amlich.Solar2Lunar(today.Day(), int(today.Month()), today.Year(), tz)
```

Đổi từ ngày Âm Lịch sang ngày Dương Lịch

```go
lunD, lunM, lunY, leap := 1, 1, 2023, 0
tz := 7 // GMT +7
d, m, y := amlich.Lunar2Solar(lunD, lunM, lunY, leap, tz)
```

## Todo

- [x] Đổi từ ngày dương lịch sang ngày âm lịch và ngược lại
- [x] Đổi ngày tháng năm ra can chi
- [ ] Tính giờ hoàng đạo trong ngày
- [ ] Ngày lễ đặc biệt trong lịch âm
- [ ] Tính tiết khí

## Tài liệu tham khảo

- Hồ Ngọc Đức
  [https://www.informatik.uni-leipzig.de/~duc/amlich/amlich.html](https://www.informatik.uni-leipzig.de/~duc/amlich/amlich.html)
- Khác nhau giữa âm lịch Việt Nam và âm lịch Trung Quốc: \
  [https://www.informatik.uni-leipzig.de/~duc/amlich/calrules_en.html#comparison](https://www.informatik.uni-leipzig.de/~duc/amlich/calrules_en.html#comparison) \
  [https://www.informatik.uni-leipzig.de/~duc/amlich/LichTa-DoanHung.html](https://www.informatik.uni-leipzig.de/~duc/amlich/LichTa-DoanHung.html)
