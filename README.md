### Services

order-service

- java
- rest api

payment-service

- python

user-service

- typescript
- contains endpoint for login, signup

shipping-service

- go

inventory-service

- java
- rest api crud functionality for inventory items

---

### Example Workflow

- Create order via endpoint on order service. Customer service calls user service to authenticate user. Emits event if order requested.
- On order requested, payment service attempts to reserve stock.
- On stock reserved, payment service attempts to make transaction.
- On empty stock, order service cancels order.
- On transaction failed, order service cancels order, emits order cancelled event, inventory service unreserves inventory.
- On transaction completed, order service accepts order request and emits event.
- On order request accepted, shipping service looks for inventory. On inventory found, shipping service ships, order service completes order.
- On inventory not found, order service cancels order, emits order cancelled event, payment service performs refund.

### Isolated Behaviours:

order-service:

payment-service:

- On order request attempt to perform transaction.
- On order cancelled see if a refund is in order.

user-service:

shipping-service:

inventory-service:

---

### Entities

order-service:

payment-service:

user-service:

shipping-service:

inventory-service:
