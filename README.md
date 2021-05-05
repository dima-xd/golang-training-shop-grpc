# golang-training-shop-grpc

This is grpc-gateway project for working with products via database shop.

**To run this application locally** you need to clone this project and also you need to have active docker server.



|             Path            | Method | Description                           | Body example                                                                                                                                                                                                                     |
|:---------------------------:|--------|---------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| /api/v1/products                   | GET    | get all products                      |```{"product":[{"id":"1","name":"Kingston DDR4-1600 8 Gb","productCategoryId":"3","quantity":"15","unitPrice":"$67.00"},{"id":"2","name":"Samsung Galaxy Buds","productCategoryId":"1","quantity":"8","unitPrice":"$120.00"},{"id":"3","name":"Samsung Electronics EVO Select 256GB MicroSDXC","productCategoryId":"1","quantity":"26","unitPrice":"$30.00"},{"id":"4","name":"Windows 10 Pro Upgrade","productCategoryId":"5","quantity":"40","unitPrice":"$100.00"},{"id":"5","name":"Kraus KAG-2MB Dishwasher Air Gap","productCategoryId":"2","quantity":"31","unitPrice":"$25.00"}]}```|
| /api/v1/products                   | POST   | create new product                    |                                                                                                                                                                                                                                  |
| /api/v1/products/{id}              | GET    | get product by the id                 | ```{"product":{"id":"1","name":"Kingston DDR4-1600 8 Gb","productCategoryId":"3","quantity":"15","unitPrice":"$67.00"}}```                                                                                                                                  |
| /api/v1/products/{id}/{unit_price} | PUT    | update product's unit price by the id |                                                                                                                                                                                                                                  |
| /api/v1/products/{id}              | DELETE | delete product by the id              |                                                                                                                                                                                                                                  |

# Protobuf

1. To generate go code from brotobuf type this command below  in terminal while you're in project root:
`protoc -I proto product.proto --grpc-gateway_out proto --go_out=plugins=grpc:.`

# docker-compose

1. To up docker-compose type this  command below in terminal while you're in project root:
`docker-compose -f docker-compose.yaml up`
