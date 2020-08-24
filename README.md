# RackspaceAssignment

Assignment solution for -> https://gist.github.com/jbartels/d75a9f5282abebe071694723a5f25f0e

To create a cart `curl`

```curl
curl --location --request POST 'http://ambhorepranav1c.mylabserver.com/cart/save/' \
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
curl --location --request POST 'http://ambhorepranav1c.mylabserver.com/cart/checkout/becece80-5653-487f-b4ae-085cb2fb210c/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cartItems": [
        {"productCode": "CH1"},
        {"productCode": "AP1"}
    ]
}'
```

Here replace `becece80-5653-487f-b4ae-085cb2fb210c` with your cart id.