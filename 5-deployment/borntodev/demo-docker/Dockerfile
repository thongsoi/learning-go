FROM 1.22.1-alpine3.19
WORKDIR /home/jack/learning-go/5-deployment/borntodev/demo-docker
COPY . .
RUN go build -o index
EXPOSE 8000
CMD [ "./index" ]
