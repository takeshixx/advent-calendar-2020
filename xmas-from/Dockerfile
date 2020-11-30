FROM alpine
RUN apk update && apk add go make
WORKDIR /srv
COPY server.go Makefile /srv/
RUN make
ENV XMAS_SECRET todo
CMD /srv/bin/server