Clean Architecture (Kiến Trúc Sạch)

- Ref: https://medium.com/@DrunknCode/clean-architecture-simplified-and-in-depth-guide-026333c54454
- Ref: https://dev.to/julianlasso/clean-architecture-domain-layer-3bdd
- Ref: https://medium.com/@rayato159/how-to-implement-clean-architecture-in-golang-en-f50d66378ebf

Kiến trúc tổng quan của hệ thống bao gồm các thành phần được mô tả như hình dưới đây

![arch](/docs/img/arch/arch.png)

Trong đó, một microservice sẽ được chia thành các layer cơ bản như sau:

- controllers: đóng vai trò như lớp để kiểm tra giá trị vào
- domain: thực hiện logic chính của ứng dụng
- repo: các chức năng truy vấn vào cơ sở dữ liệu
- data (models): các models của ứng dụng sẽ khai báo ở đây

Tất cả các hàm public ở `controllers` và `domain` phải được khai báo trong các file proto tương ứng ở `api/proto`. Xem ví dụ ở `api/proto/tickex/greeter`

Hãy tuân thủ các nguyên tắc chính của clean architecture

- Các phụ thuộc chỉ được phép đi từ ngoài vào trong (từ tầng cao hơn xuống tầng thấp hơn). Các phụ thuộc không được phụ thuộc trực tiếp mà phải thông qua interface/abstract class
- Không có tầng nào bên trong được phép phụ thuộc vào tầng bên ngoài.
