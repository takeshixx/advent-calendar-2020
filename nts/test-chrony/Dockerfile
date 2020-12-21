FROM fedora:latest

RUN dnf install -y --nodocs chrony && \
 dnf -y clean all

# Specify it during `docker run` as parameter: "-p <host_port>:<container_port>"
EXPOSE 123

CMD ["/usr/sbin/chronyd", "-d", "-d"]