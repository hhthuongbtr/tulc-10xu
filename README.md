# Giới thiệu
    Đây là mã nguồn được xây đựng nhằm tạo api callback
    Chức năng: nhận cảnh báo từ Tencent monitor và gửi cảnh báo qua Telegarm
    Uri api: http://<host>/api/v1/Callback
    Tích hợp: vui lòng tham khảo tài liệu phía Tencent
# Chuẩn bị:
    Tạo bot chat Telegram, ứng dụng cần các thông tin:
        + Chat ID
        + Chat bot Token
# Chức năng:
    - Hỗ trợ cấu hình chat bot qua file
    - Hỗ trợ parse một số mẫu cảnh báo như:
        + Shutdown, lost connect 
        + Disk utilization
        + CPU load
        + CPU Util
        + Memory
        + Các cảnh báo còn lại ứng dụng sẽ gửi nội dung cảnh báo gốc đến Telegram thông qua bot chat
# Sử dụng:
    - Xem thêm tại mục deployment