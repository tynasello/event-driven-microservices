### Services:

All services are managed through Docker and emit and receive events from a Kafka broker.

order-service:

- Endpoint to create an order, use user-service to login, save user username to order.
- Emits order requested event (with order id, product name & quantity).

- On inventory not found, cancel order, emit event.
- On inventory reserved, update inventory reserved column value for order.

- On transaction failed, cancel order, emit event (with order id, product name & quantity).
- On transaction completed, accept order request and emit event.

- On order shipped, complete order.

payment-service:

- On inventory found attempt to perform transaction.
- Randomly accept or decline payment, emit transaction completed/failed event (with order id).

user-service:

- Endpoint to signup, login.

shipping-service:

- On order accepted, emit order shipped event (with order id).

inventory-service:

- Endpoint to create inventory items.
- On order requested, look for inventory.

  - If found emit inventory found (with order id), and reserve inventory. If not found, emit inventory not found event (with order id).

- On order cancelled event, free any reserved inventory.

### Domain Entities

order-service:

- Order entity:
  - order username
  - status
  - product name
  - product quantity
  - inventory reserved: bool

user-service:

- User:
  - username
  - password

inventory-service:

- Inventory:
  - product name
  - quantity
  - quantity reserved
