
FROM iron/go
WORKDIR /app

ADD grpc_client /app/

CMD [ "./grpc_client" ]

