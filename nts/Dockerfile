FROM debian:buster

WORKDIR /build

RUN apt-get update ;\
    apt-get install -y curl

RUN curl -o ntpsec.tar.gz https://gitlab.com/NTPsec/ntpsec/-/archive/NTPsec_1_2_0/ntpsec-NTPsec_1_2_0.tar.gz

# prepare build environment
RUN tar -xzf ntpsec.tar.gz ;\
    rm ntpsec.tar.gz ;\
    cd ntpsec-NTPsec_* ;\
    ./buildprep

# install ntpsec
RUN cd ntpsec-NTPsec_* ;\
    ./waf configure && \
    ./waf build && \
    ./waf install

# install gpsd
RUN apt-get install -y gpsd

# create config files
RUN mkdir /etc/ntp.d && \
    mkdir /var/log/ntpstats && \
    mkdir /var/lib/ntp

COPY config /etc/ntp.d/config

# clean up
RUN rm -rf /var/lib/apt/lists/* && \
    rm -rf /build/ntpsec-NTPsec_*

# fix tools like ntpq 
RUN echo "/usr/local/lib/ntp" > /etc/ld.so.conf.d/ntp.conf
RUN ldconfig 

EXPOSE 4460/udp
EXPOSE 4460/tcp

ENTRYPOINT /usr/local/sbin/ntpd -c /etc/ntp.d/config -n