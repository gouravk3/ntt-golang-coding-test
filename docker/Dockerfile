FROM golang:1.19-alpine as build 
ENV GO111MODULE=on
ENV CGO_ENABLED=0
COPY . /service
WORKDIR /service
go build -o exoplanet cmd/exoplanet/main.go
go build -o fuelestimation cmd/fuelestimation/main.go

FROM alpine:3.17
WORKDIR /service
COPY --from=build /service/exoplanet .
COPY --from=build /service/fuelestimation .

CMD [ "./exoplanet", "./fuelestimation" ]

