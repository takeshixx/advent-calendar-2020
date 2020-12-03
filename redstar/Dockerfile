FROM nginx
RUN apt update && apt install -y python3 python3-pip
RUN pip3 install pycrypto
WORKDIR /srv
RUN mkdir logs
COPY www /srv/www
COPY nginx.conf /etc/nginx/nginx.conf
COPY wm_decrypt.py ./
COPY kindergarten.jpg /srv/www
RUN ./wm_decrypt.py e /srv/www/kindergarten.jpg "One of Santa's secrets: xHGid0BvzJfc6ezGh293HxXbFJrbbJNG"
EXPOSE 80
CMD nginx