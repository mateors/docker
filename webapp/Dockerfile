FROM golang:alpine
ENV GO111MODULE=on
ENV BG_COLOR=skyblue
WORKDIR /opt/webapp
COPY . .
# Install the package
RUN go build
RUN go install -v ./...
EXPOSE 8180
CMD ["dockerapp"]