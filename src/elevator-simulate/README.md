# Cài đặt Golang

    https://golang.org/dl/

# Chạy

    go run *.go

# Thiết kế
Ví dụ một toà nhà 10 tầng có 1 thang máy.
Khi hành khác gửi 1 yêu cầu lên hoặc xuống, trong yêu cầu đó có:
	- Current - vị trị mà hành khách đang đứng.
	- Destination - vị trí hành khách muốn tới.
	- Direction - hướng yêu cầu.
Trong thang máy sẽ có 2 biến quan trọng là 2 ma trận:
	- 1: PickUp lưu tất cả các yêu cầu sử dụng thang máy.
	- 2: DropOff lưu tất cả các yêu cầu hành khách muốn đến.
Thang máy sẽ di chuyển qua từng tầng và đồng thời kiểm tra xem trong 2 ma trận PickUp và DropOff có chứa yêu cầu của hành khách không. Nếu có thì chạy 2 hàm đó, làm rỗng hàng đấy và tiếp tục di chuyển cho đến tầng cao nhất cùng hướng đang có trong 2 ma trận đó.

Khi nhiều người xài quá nên nâng cấp lên 3 thang máy :v
Để cho 3 thang máy này có thể hoạt động cùng lúc với nhau thì em tạo ra các [Goroutine và Channel](https://www.golang-book.com/books/intro/10).
Em tạo ra 2 goroutine. Do các goroutine chạy mà không cần phải chờ đợi lẫn nhau:
	- 1: Dùng để lắng nghe các sự kiện Logger và chạy các hàm PickUp, DropOff, Move. 
	- 2: Dùng để gửi các Request từ hành khách đến Goroutine 1.

Đồng thời em cũng tạo ra một Chanel để truyển các Request của hành khách vào Goroutine. Để giao tiếp với Goroutine cần phải thông qua Chanel. 

Khi nhận được một sự kiện sẽ chạy một hàm để tìm thang máy còn trống gần đó nhất để chuyển đến vị trí hiện tại của hành khác. Em có tìm thấy thuật toán [Nearest Car](https://www.quora.com/Is-there-any-public-elevator-scheduling-algorithm-standard) để áp dụng cho thang máy.