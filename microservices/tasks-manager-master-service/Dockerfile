FROM liuguanyan/grpc-golang

WORKDIR /app
ADD conf/tasks-manager-master-config.json /app/config.json
ADD tasks-manager-master-service/main /app/main


EXPOSE 8080
EXPOSE 50051

CMD ["/app/main"]

