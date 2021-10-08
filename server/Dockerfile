###################################################################################################
#                                                                                                 #
#                                   Miron-developer                                               #
#                                   Photographer                                                  #
#                                                                                                 #
###################################################################################################

FROM golang:1.16

COPY . .
WORKDIR /internal
RUN go mod download; go build -o ./cmd/main/main cmd/main/main.go

LABEL description="This is Photographer project" \
    authors="Miron-developer" \
    contacts="https://github.com/miron-developer" \
    site="https://photographer.herokuapp.com"

CMD ["cmd/main/main"]

EXPOSE 4430