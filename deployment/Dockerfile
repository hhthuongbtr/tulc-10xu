FROM centos:7.2.1511

RUN mkdir /app
WORKDIR /app
ADD tulc-10xu ./tulc-10xu
RUN mkdir -p /opt/AlarmFor10xu
ADD application.yml /opt/AlarmFor10xu

RUN rm -f etc/localtime; ln -sf /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
EXPOSE 8000/tcp
# Run
CMD ["/app/tulc-10xu", "-m", "http"]
