FROM golang:1.20
RUN wget https://github.com/blablatov/web-base64jwt-codec.git
WORKDIR /webparser
COPY . .
RUN go build -o /webparser 
EXPOSE 8060
CMD ["./webparser"]
