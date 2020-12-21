FROM alpine
RUN apk update && apk add go make libpcap-dev
WORKDIR /srv
COPY server.go Makefile go.mod /srv/
RUN make
ENV XMAS_SECRET "ZmQ8GCDbOnh2XcdwOxEESyxW92G8Zgix"
CMD /srv/bin/server eth0@if101