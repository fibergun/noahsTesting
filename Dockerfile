FROM node:latest AS frontendbuilder
WORKDIR /build
COPY backend/src/frontend/package.json package.json
COPY backend/src/frontend/package-lock.json package-lock.json
RUN npm ci
COPY backend/src/frontend .
RUN npm run build

FROM golang:1.26-alpine AS backendbuilder
WORKDIR /build
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY ./backend ./backend
COPY --from=frontendbuilder /build/dist backend/src/frontend/dist
RUN cd backend && go build -o noahsspel .

FROM alpine:latest
COPY --from=backendbuilder /build/backend/noahsspel /usr/bin/noahsspel
EXPOSE 8080
WORKDIR /opt
VOLUME /opt
CMD ["/usr/bin/noahsspel"]