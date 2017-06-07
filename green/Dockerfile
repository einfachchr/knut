# von diesem Image starten wir (vanilla nginx)
FROM golang:1.8

# zum Debugging
RUN apt-get update
RUN apt-get install -y curl

# legt das Verzeichnis an und wechselt hinein.
WORKDIR /go/src/app
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]

# der Go-Server lauscht auf Port 8000
EXPOSE 8000
