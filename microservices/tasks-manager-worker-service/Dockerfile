FROM liuguanyan/grpc-golang

WORKDIR /app
ADD conf/tasks-manager-worker-config.json /app/config.json
ADD tasks-manager-worker-service/main /app/main

CMD ["/app/main"]

