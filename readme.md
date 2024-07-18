# Fisheries API

# Getting Started
<ins>**1. Downloading the repository**</ins>

Start by cloning the repository with `git clone --recursive https://github.com/harisspace/fisheries-api`

<ins>**2. Build and Running**<ins>

Setup your postgresql if you have locally or you can use containerization approach by running `docker compose -f docker-compose.yml up -d`

Run `main.go` file in fisheries-api folder using this command `go run ./main.go` at relative directory

Default http port running is 3000, you can change the port number in `.env` file

# Code Structure
 ```
 config
 middleware
 modules                      (Domain partitioning approach)
    - farm
        - handlers            (Http handler)
        - models              (Structure model)
        - respositories       (Repository to consume database)
        - usecases            (Business logic)
    - statistic
        - handlers            (Http handler)
        - models              (Structure model)
        - respositories       (Repository to consume database)
        - usecases            (Business logic)
  pkg
    - database                (Database connection and configuration)
    - http_error              (HTTP error wrapper)
    - utils                   (Utility)
```

# List of Enpoints

**This endpoint use authorization type Basic Auth with *username = delos password = 2024***

    - Farm
        - /v1/farm --> [GET] Get List of Farm
            - query
                - page [Number, OPTIONAL]
                - quantity [Number, Optional]
                - order [String, Optional]
            - expected response
                - [200] Return list of Farm
                - [404] If there is no record found
        - /v1/farm --> [POST]
            - body (JSON)
                - name [String, REQUIRED]
            - expected response
                - [200] Return the new created farm
                - [409] If name already exist
        - /v1/farm/:id --> [GET]
            - param
                - id --> Farm id
            - expected response
                - [200] Return farm data
                - [404] Farm data not found
        - /v1/farm --> [PUT]
            - body (JSON)
                - farm_id [String, OPTIONAL]
                - name [String, REQUIRED]
            - expected response
                - [200] Updated or created
        - /v1/farm/:id [DELETE]
            - param
                - id --> Farm id
            - expected response
                - [200] Soft delete success
                - [404] Farm not found
    
    - Pond
        - /v1/pond --> [GET] Get List of Pond
            - query
                - page [Number, OPTIONAL]
                - quantity [Number, Optional]
                - order [String, Optional]
            - expected response
                - [200] Return list of pond
                - [404] If there is no record found
        - /v1/pond --> [POST]
            - body (JSON)
                - name [String, REQUIRED]
            - expected response
                - [200] Return the new created pond
                - [409] If name already exist or farmId doesn't exist
        - /v1/pond/:id --> [GET]
            - param
                - id --> Pond id
            - expected response
                - [200] Return pond data
                - [404] Pond data not found
        - /v1/pond --> [PUT]
            - body (JSON)
                - pond_id [String, REQUIRED]
                - name [String, REQUIRED]
                - farm_id [String, REQUIRED]
            - expected response
                - [200] Updated or created
        - /v1/pond/:id [DELETE]
            - param
                - id --> Pond id
            - expected response
                - [200] Soft delete success
                - [404] Pond not found
    
    - Statistic
        - /v1/statistic/:id --> [GET]
            - param
                - id --> User-agent
            - expected response
                - [200] List of user-agent statistic record
                - [404] User agent not found

# Postman collection

[<img src="https://run.pstmn.io/button.svg" alt="Run In Postman" style="width: 128px; height: 32px;">](https://app.getpostman.com/run-collection/11874264-cbc1f53a-f62f-45d5-a27b-163acd3055bf?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D11874264-cbc1f53a-f62f-45d5-a27b-163acd3055bf%26entityType%3Dcollection%26workspaceId%3D4fda4a3d-a6a8-4345-bd53-8bec6724d399)

# Unit test 

Running unit test `go test ./test/modules/farm/usecases -v`
