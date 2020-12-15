FROM alpine
RUN apk update && apk add go make nmap
WORKDIR /srv
COPY server.go Makefile /srv/
RUN make
ENV XMAS_SECRET "sUyHObOMPnf4eu56IMaWEJqI1ng5VkoS"
CMD /srv/bin/server