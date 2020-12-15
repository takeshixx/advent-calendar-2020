FROM nginx:alpine
RUN apk update && apk add nmap-ncat
WORKDIR /srv
COPY www /srv/www
COPY nginx.conf /etc/nginx/nginx.conf
RUN mkdir logs
EXPOSE 8000
CMD nginx && ncat --ssl-cert /etc/letsencrypt/archive/xmas.rip/fullchain1.pem --ssl-key /etc/letsencrypt/archive/xmas.rip/privkey1.pem --sctp -kl --sh-exec "ncat localhost 80" -p 8000 