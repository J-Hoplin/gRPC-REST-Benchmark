## Head API

---

Use to server real user's request

---

- Framwork: Gin Gonic

### Router for Invoke Rest service

- `/rest`

  - Query

    - `from`: Should bigger than 0, `required`
    - `to`: Should bigger than 0, `required`

  - Example

    ```bash
    127.0.0.1:8080/rest?from=1&to=100000
    ```
