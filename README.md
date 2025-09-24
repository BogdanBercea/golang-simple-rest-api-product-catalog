# EASY-SHOP

## Product CatalogService Simple REST API

This REST API was developed using GO programming language and Gin Web Framework.

### Available endpoints

1. <b>Get product catalog</b> - use this endpoint to retrieve all available products.
    - Method: GET
    - Endpoint:```http://localhost:9090/products?brand=puma```

2. <b>Get product by id</b> - use this endpoint to retrieve a product by id
    - Method: GET
    - Endpoint: ```http://localhost:9090/products/20```

3. <b>Create a new product</b> - use this endpoint to create a new product and store it in product catalog
    - Method: POST
    - Endpoint: ```http://localhost:9090/products```
    - Example payload:
    ```
    {
    "id": 200,
    "name": "t-shirt-new",
    "available_sizes": [
        "XXL",
        "S",
        "M",
        "XS",
        "L",
        "XL"
    ],
    "images": [
        "https://image1.svg",
        "https://image3.svg"
    ],
    "brand": "puma"
    }
   ```