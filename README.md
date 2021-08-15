# buzzao-test

## Start
- > $ go run main.go

## Endpoint

- POST `/config/{n}`
  - `n` is the amount of threads configured

- POST `/process`
  - ```json
        {
            "nums": [1,2,3,4,5]
        }
    ```
  - Returns:
  - ```json
        {
            "result": 15
        }
    ```
