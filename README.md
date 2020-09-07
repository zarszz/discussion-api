# Discussion API

## golang version 
- 1.5

## library requirements
- gin
- gorm
- godotenv

## Installation
* define your .env file in <b>root project</b>
1. define CONN_STRING in .env file, for example :    
    ```shell
        CONN_STRING="user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
    ```
2. define LISTEN_PORT in .env file, for example : 
    ```shell
        LISTEN_PORT=":8080"
    ```
3. running docker-compose
    ```shell
        $~ docker-compose up --build -d
    ```
