FROM nginx
WORKDIR /srv
COPY www /srv/www
COPY nginx.conf /etc/nginx/nginx.conf
RUN mkdir /srv/logs
EXPOSE 80
EXPOSE 443
CMD nginx