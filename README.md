# RackspaceAssignment

Assignment solution for -> https://gist.github.com/jbartels/d75a9f5282abebe071694723a5f25f0e

## How to install?
First of all Install docker, then run the following commands
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

Response should be
```json
{
   "data":{
      "cart":{
         "ID":"96c9e508-c244-4f55-8381-af954711a3ee",
         "Items":[
            {
               "ID":"2e5b815e-15ac-4f2e-bccd-54ee0f2644bb",
               "Name":"Chai",
               "Code":"CH1",
               "Price":3.11,
               "Discount":null
            },
            {
               "ID":"830703cb-8130-47e6-b409-478bd0b30fe9",
               "Name":"Apples",
               "Code":"AP1",
               "Price":6,
               "Discount":null
            },
            {
               "ID":"e856c9c8-63f7-4cf0-bb3e-804fb3b3a01c",
               "Name":"Apples",
               "Code":"AP1",
               "Price":6,
               "Discount":null
            },
            {
               "ID":"0d3c492e-5922-48fd-b558-7f3c4ff06756",
               "Name":"Apples",
               "Code":"AP1",
               "Price":6,
               "Discount":null
            }
         ],
         "Total":21.11,
         "CartMap":{
            "AP1":3,
            "CH1":1
         }
      }
   }
```

***

To checkout the cart `curl` the following, this will apply the discount rules
```
curl --location --request POST 'http://localhost:8080/cart/checkout/96c9e508-c244-4f55-8381-af954711a3ee/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cartItems": [
        {"productCode": "CH1"},
        {"productCode": "AP1"}
    ]
}'
```

The result should be,
```json
{
   "data":{
      "cart":{
         "ID":"96c9e508-c244-4f55-8381-af954711a3ee",
         "Items":[
            {
               "ID":"2e5b815e-15ac-4f2e-bccd-54ee0f2644bb",
               "Name":"Chai",
               "Code":"CH1",
               "Price":3.11,
               "Discount":null
            },
            {
               "ID":"830703cb-8130-47e6-b409-478bd0b30fe9",
               "Name":"Apples",
               "Code":"AP1",
               "Price":4.5,
               "Discount":{
                  "Name":"APPL",
                  "Tag":"",
                  "Price":-1.5
               }
            },
            {
               "ID":"e856c9c8-63f7-4cf0-bb3e-804fb3b3a01c",
               "Name":"Apples",
               "Code":"AP1",
               "Price":4.5,
               "Discount":{
                  "Name":"APPL",
                  "Tag":"",
                  "Price":-1.5
               }
            },
            {
               "ID":"0d3c492e-5922-48fd-b558-7f3c4ff06756",
               "Name":"Apples",
               "Code":"AP1",
               "Price":4.5,
               "Discount":{
                  "Name":"APPL",
                  "Tag":"",
                  "Price":-1.5
               }
            },
            {
               "ID":"49d2b474-35dc-4010-9f70-b19248a7a96d",
               "Name":"Milk",
               "Code":"MK1",
               "Price":0,
               "Discount":{
                  "Name":"CHMK",
                  "Tag":"",
                  "Price":-4.75
               }
            }
         ],
         "Total":16.61,
         "CartMap":{
            "AP1":3,
            "CH1":1,
            "MK1":1
         }
      }
   }
}
```

Here replace `becece80-5653-487f-b4ae-085cb2fb210c` with your cart id.

***

## How to test?
To test this project run the command

```bash
docker build --file DockerfileTestLocal -t racktest .
```
Or
```bash
docker build --file DockerfileTestMaster -t racktest .
```

It will provide the test information with the coverage stats for each function. It should look like this,

```
.
.
github.com/pranav93/RackspaceAssignment/models/cart.go:43:		ApplyDiscount	100.0%
github.com/pranav93/RackspaceAssignment/models/cart.go:53:		ApplyPriceItem	100.0%
github.com/pranav93/RackspaceAssignment/models/cart.go:71:		AddFreeItem	76.9%
.
.
.
```

***
To apply discounts, we use a set of rules. The structs for it are,
```go
// Rule is a struct for storing rules data
type Rule struct {
	ID     int
	Rtype  string  // Rule can be either `get a free item` or `apply discount on item` i.e. 'applyPrice'/ 'getFree'
	Name   string
	Action Action  // Each rule has an action struct to check the condition
	Result Result  // Each rule has an result struct to apply needed result
}

// Action is a struct for storing actions for a rule
type Action struct {
	ID          int
	RuleID      int
	ProductCode string
	Operator    string  // types are "ge", "le", "g", "l", "eq". Practically "ge" is used
	Qty         int  // quantity to be matched
}

// Result is a struct for storing results for a rule
type Result struct {
	ID               int
	RuleID           int
	ProductCode      string
	Qty              int  // Amount of items that should be added or discounted
	AppliedPrice     float64  // flat or percent price to be applied
	AppliedPriceType string  // either `flat` or `percent`
	ResultType       string  // Rule can be either `get a free item` or `apply discount on item` i.e. 'applyPrice'/ 'getFree'
}

```