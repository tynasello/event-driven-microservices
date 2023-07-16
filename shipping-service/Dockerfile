FROM rust:latest AS builder
WORKDIR /usr/src/app
COPY . .
RUN cargo build --release
CMD ["cargo", "run", "--release"]

