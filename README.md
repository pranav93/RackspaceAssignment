# RackspaceAssignment

Assignment solution for -> https://gist.github.com/jbartels/d75a9f5282abebe071694723a5f25f0e

How to install??
Install docker, then
```bash
git clone https://github.com/pranav93/RackspaceAssignment.git
cd RackspaceAssignment
docker build -t rack .
docker container run -p 8080:3000 rack
```

To create a cart `curl`

```curl
curl --location --request POST 'http://127.0.0.1:8080/cart/save/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cartItems": {
        "CH1": 1,
        "AP1": 3
    }
}'
```

To checkout the cart `curl` the following
```
curl --location --request POST 'http://localhost:8080/cart/checkout/35d87482-1869-4fe8-bbca-af5761c412cb/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cartItems": [
        {"productCode": "CH1"},
        {"productCode": "AP1"}
    ]
}'
```

Here replace `becece80-5653-487f-b4ae-085cb2fb210c` with your cart id.