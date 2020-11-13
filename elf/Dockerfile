FROM httpd:2.4

RUN apt-get update -y
RUN apt-get install -y python3 python3-pip
RUN pip3 install lief

ENV SECRET replacemeplease

COPY ./elf.py /
RUN python3 /elf.py $SECRET
COPY ./elf.png /usr/local/apache2/htdocs/
COPY ./BeyondWonderland.ttf /usr/local/apache2/htdocs/
