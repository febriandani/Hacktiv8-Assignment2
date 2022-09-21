# Create New Customer

Method: `POST`
URL: `localhost:8080/api/v1/customer`
Content-Type: application/json

input :
`
{
    "customer name" : "Wandi Subarkah"
}
`

Response :
`
{
    "meta": {
        "message": "Success to create customer",
        "code": 200,
        "status": "success"
    },
    "data": {
        "customer id": 7,
        "customer name": "Wandi Subarkah"
    }
}
`
# Create Order
Method: `POST`
URL: `localhost:8080/api/v1/order`
Content-Type: application/json

input :
`
{
    "customer_id" : 5,
    "items_code" : "ip512pr4mx",
    "description" : "Iphone xs Max 512 GB",
    "quantity" : 1
}
`

Response if customer_id Not Found:
`
{
    "meta": {
        "message": "Failed to create order, customer id not found",
        "code": 400,
        "status": "error"
    },
    "data": null
}
`

Response success:
`
{
    "meta": {
        "message": "Success to create order",
        "code": 200,
        "status": "success"
    },
    "data": {
        "orderedAt": "0001-01-01T00:00:00Z",
        "order id": 70,
        "customer id": 7,
        "items": {
            "item code": "ip512pr7mx",
            "description": "Iphone xs Max 512 GB",
            "quantity": 1
        }
    }
}
`

# Update Order
Method: `PUT`
URL: `localhost:8080/api/v1/order/:id`
Content-Type: application/json

input :
`
{
    "customer_id" : 7,
    "items_code" : "tesUpdate53",
    "description" :"descUpdate53",
    "quantity" : 5
}
`

Response : 
`
{
    "meta": {
        "message": "Success to update Order",
        "code": 200,
        "status": "success"
    },
    "data": {
        "orderedAt": "0001-01-01T00:00:00Z",
        "order id": 53,
        "customer id": 7,
        "items": {
            "item code": "tesUpdateCus7",
            "description": "descUpdateCus7",
            "quantity": 5
        }
    }
}
`

# Get All Orders
Method: `GET`
URL: `localhost:8080/api/v1/orders`
Content-Type: application/json
Response :
`
{
    "meta": {
        "message": "Orders",
        "code": 200,
        "status": "success"
    },
    "data": [
        {
            "orderedAt": "0001-01-01T00:00:00Z",
            "order id": 70,
            "customer id": 7,
            "items": {
                "item code": "ip512pr7mx",
                "description": "Iphone xs Max 512 GB",
                "quantity": 1
            }
        },
        {
            "orderedAt": "0001-01-01T00:00:00Z",
            "order id": 53,
            "customer id": 7,
            "items": {
                "item code": "tesUpdateCus7",
                "description": "descUpdateCus7",
                "quantity": 5
            }
        },
        {
            "orderedAt": "0001-01-01T00:00:00Z",
            "order id": 54,
            "customer id": 2,
            "items": {
                "item code": "ip562pr2mx",
                "description": "Iphone 13 Pro Max 512 GB",
                "quantity": 1
            }
        },
        {
            "orderedAt": "0001-01-01T00:00:00Z",
            "order id": 55,
            "customer id": 2,
            "items": {
                "item code": "ip502pr2mx",
                "description": "Iphone 11 512 GB",
                "quantity": 1
            }
        },
        {
            "orderedAt": "0001-01-01T00:00:00Z",
            "order id": 56,
            "customer id": 2,
            "items": {
                "item code": "ip532pr2mx",
                "description": "Iphone xr 512 GB",
                "quantity": 1
            }
        },
        {
            "orderedAt": "0001-01-01T00:00:00Z",
            "order id": 57,
            "customer id": 3,
            "items": {
                "item code": "ip532pr3mx",
                "description": "Iphone xs Max 512 GB",
                "quantity": 1
            }
        },
        {
            "orderedAt": "0001-01-01T00:00:00Z",
            "order id": 58,
            "customer id": 4,
            "items": {
                "item code": "ip512pr4mx",
                "description": "Iphone xs Max 512 GB",
                "quantity": 1
            }
        }
    ]
}
`

# Delete Order
Method: `DELETE`
URL: `localhost:8080/api/v1/delete/:id`
Content-Type: application/json
Response : 
`
{
    "meta": {
        "message": "Success delete Order",
        "code": 200,
        "status": "success"
    },
    "data": {
        "orderedAt": "0001-01-01T00:00:00Z",
        "order id": 70,
        "customer id": 7,
        "items": {
            "item code": "ip512pr7mx",
            "description": "Iphone xs Max 512 GB",
            "quantity": 1
        }
    }
}
`
