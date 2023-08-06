### Event-Driven Microservices

Building an ordering system to learn more about microservices, event-driven architectures, Kubernetes, Kafka, and some languages (Go and Rust).

All services are managed by Kubernetes and events are emitted and consumed through Kafka.

- To deploy services locally with Kubernetes run: `./bin/deploy-k8s-local.sh`.

- To teardown Kubernetes deployments run: `./bin/teardown-k8s-local.sh`.

- To interact with and oversee the system, use the management-service CLI by running `./bin/management-cli-local.sh CLI_COMMAND`.

- Each service has it's own `./bin/up-local.sh`, and `./bin/down-local.sh` scripts for local development using docker-compose.

### The System

<img width="1311" alt="Screenshot 2023-08-06 at 4 01 22 PM" src="https://github.com/tynasello/event-driven-microservices/assets/63558019/699ff1c0-0a1b-46d3-8441-233e2e963ba6">
<br><br>

**Management Service**:

- CLI allowing a user to perform actions (login, create order, etc.), and oversee the flow of events in the system (**Rust** and the Clap crate).

**Order Service**:

- REST API exposing endpoints to manage user orders (**Java**, **PostgreSQL**, and **Spring Boot**).
- Service consumes various events, updates order entities, and emits events to reflect status' of orders.

**Payment Service**:

- **Python** service that attempts to complete a transaction once inventory had been reserved for an order.

**User Service**:

- REST API allowing users to signup and login (**Go**, **MySQL**). Authentication implemented using JWT tokens stored in http-only cookies. These tokens are used for authentication in the order service.

**Shipping Service**:

- **Rust** service that emits a order shipped event when an order has been accepted (inventory reserved and payment successful)

**Inventory Service**:

- REST API exposing endpoints to create and update inventory items.
- Service validates that inventory is free to reserve when a order is created, and frees any reserved inventory when an order is cancelled (**Go**, **PostgreSQL**).
