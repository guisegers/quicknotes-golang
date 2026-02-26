FROM golang:1.26.0-bookworm
LABEL authors="Guilherme Segers"


ENTRYPOINT ["top", "-b"]