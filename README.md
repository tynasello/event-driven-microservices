Building an ordering system to learn more about microservices, event-driven architectures, Kubernetes, Kafka, and some languages (Go and Rust).

All services are managed by Kubernetes (can use docker-compose for local development) and events are emitted and consumed via Kafka.

---

### The Services

management-service:

- CLI allowing a user to perform actions, and oversee and ~~view the flow of events in the system~~.
- Service created with Rust and the Clap crate.

order-service:

- REST API allows users to create orders.
- Service emits ORDER_REQUESTED event when a user attempts to create an order.
- On INVENTORY_NOT_RESERVED event, service updates order status column to CANCELLED and emits ORDER_CANCELLED event.
- On INVENTORY_RESERVED event, service updates inventory reserved column to true for order.
- On TRANSACTION_FAILED event, service updates order status column to CANCELLED , emits ORDER_CANCELLED event.
- On TRANSACTION_COMPLETED, service updates order status column to APPROVED and emits ORDER_ACCEPTED event.
- On ORDER_SHIPPED event, service updates order status column to COMPLETED and emits ORDER_COMPLETED event.
- Service created with Java, PostgreSQL, and Spring Boot/Spring Data JPA/Maven.

payment-service:

- On INVENTORY_RESERVED event, service attempts to perform a transaction (hard-coded probability).

  - If attempted transaction is successful, service emits TRANSACTION_COMPLETED event.
  - If attempted transaction is not successful, service emits TRANSACTION_FAILED event.

- Service created with Python and MySQL.

user-service:

- REST API allows users to signup and login. Authentication implemented using JWTs.
- Upon login, an access JWT token is stored in http-only cookies. This token is used in the order service.
- Service created with Go, PostgreSQL, JWT, and Gin/Gorm.

shipping-service:

- On ORDER_ACCEPTED event, service emits ORDER_SHIPPED event.
- Service created with Rust, and PostgreSQL.

inventory-service:

- REST API contains endpoint to create and update inventory items.
- On ORDER_REQUESTED event, service validates that inventory is free to reserve.

  - If inventory is found, service emits INVENTORY_RESERVED event.
  - If inventory is not found, service emits INVENTORY_NOT_RESERVED event.

- On ORDER_CANCELLED event, service frees any reserved inventory for the corresponding order.
- Service created with Go, PostgreSQL, and Gin/Gorm.

### The Events

- ORDER_REQUESTED
- ORDER_CANCELLED
- ORDER_SHIPPED
- ORDER_COMPLETED
- ORDER_ACCEPTED
- INVENTORY_NOT_RESERVED
- INVENTORY_RESERVED
- TRANSACTION_COMPLETED
- TRANSACTION_FAILED
