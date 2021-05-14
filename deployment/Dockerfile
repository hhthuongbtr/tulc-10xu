FROM centos:7.2.1511

RUN mkdir /app
WORKDIR /app
ADD talarm ./talarm
RUN mkdir -p /opt/AlarmFor10xu
ADD application.yml /opt/AlarmFor10xu/application.yml

RUN rm -f etc/localtime; ln -sf /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
EXPOSE 8000/tcp
# Run
CMD ["/app/talarm", "-m", "http"]
