FROM rust
WORKDIR /usr/app
COPY . .
RUN cargo build --release
CMD ["target/release/shipping-service"]

