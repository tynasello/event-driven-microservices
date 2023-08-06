FROM rust AS builder
WORKDIR /usr/src/app
COPY bin/entrypoint.sh /bin/entrypoint.sh
RUN chmod +x /bin/entrypoint.sh
COPY . .
RUN cargo build --release
CMD ["/bin/entrypoint.sh"]
