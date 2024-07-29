FROM golang

WORKDIR /Travel-agency

COPY . .

RUN go build -o travelagency .

CMD ["/Travel-agency/travelagency"]

