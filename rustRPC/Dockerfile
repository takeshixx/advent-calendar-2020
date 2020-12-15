FROM rust:1.48 as builder

COPY tarpc /root/tarpc

ENV SECRET "H9npZk8j6Uiu31lHorHqCfwQ1j4dBwrF"

WORKDIR /root/tarpc/Xmas-service

RUN cargo build

ENTRYPOINT [ "/root/tarpc/Xmas-service/target/debug/XmasServer", "--port", "8888" ]