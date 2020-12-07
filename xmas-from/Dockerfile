FROM alpine
RUN apk update && apk add go make
WORKDIR /srv
COPY server.go Makefile /srv/
RUN make
ENV XMAS_SECRET "JuJiWZtTTQvVjDxF6qBB6F919hKappLb"
CMD /srv/bin/server