## Head API

---

Use to server real user's request

---

- Framwork: Gin Gonic

### Router for Invoke Rest service

- `/rest`: Request to REST API service endpoint

  - Query

    - `from`: Should bigger than 0, `required`
    - `to`: Should bigger than 0, `required`

  - Example

    ```bash
    127.0.0.1:8080/rest?from=1&to=100000
    ```

- `/grpc/unary`: Invoke grpc-service Unary Communication

  - Query

    - `from`: Should bigger than 0, `required`
    - `to`: Should bigger than 0, `required`

  - Example

    ```bash
    127.0.0.1:8080/grpc/unary?from=1&to=100000
    ```

- `/grpc/stream/client`: Invoke grpc-service Client-Stream Communication

  - Query

    - `from`: Should bigger than 0, `required`
    - `to`: Should bigger than 0, `required`

  - Example

    ```bash
    127.0.0.1:8080/grpc/stream/client?from=1&to=100000
    ```

- `/grpc/stream/server`: Invoke grpc-service Server-Stream Communication

  - Query

    - `from`: Should bigger than 0, `required`
    - `to`: Should bigger than 0, `required`

  - Example

    ```bash
    127.0.0.1:8080/grpc/stream/server?from=1&to=100000
    ```

- `/grpc/stream/bi`: Invoke grpc-service Bi-Directional-Stream Communication

  - Query

    - `from`: Should bigger than 0, `required`
    - `to`: Should bigger than 0, `required`

  - Example

    ```bash
    127.0.0.1:8080/grpc/stream/bi?from=1&to=100000
    ```
