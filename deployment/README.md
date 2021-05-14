# Cấu trúc file
    - application.yml: Template file config
    - Dockerfile: Docker file
    - talarm: Binary của ứng dụng, tương thích với Centos 7
# Sử dụng:
    - Tải file config application.yml lưu về /opt/AlarmFor10xu/application.yml 
```bash
        mkdir -p /opt/AlarmFor10xu
        mv application.yml /opt/AlarmFor10xu/application.yml
```
    - Điền thông tin bot chat Telegram
```bash
telegram:
        chat_id: 'Chat id ở đây'
        token_env: 'Điền token của bot chát ở đây'
```
    - Chạy daemon hoặc chạy với Docker 
```bash
        chmod +x talarm && nohup talarm -m http
        docker run -p 80:8000 -itd -e GIN_MODE=release talarm
```