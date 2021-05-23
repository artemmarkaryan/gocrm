# REST API contract

## DTO

User
```
{
    "id": 10,
    "name": "Ivan"
}
```

Order
```
{
    "id": 10,
    "userId": 1,
    "customer": Customer
    "createdTS": "2012-04-23T18:25:43.511Z", 
    "basket": Basket
    "status": "new"
}
```

Customer
```
{
    "id": 10,
    "name": "Ivan",
    "contact": map[string]string
}
```

Basket
```
{
    items: [Item]
}
```


Item
```
{
    "name": "Мясо",
    "price": 120
    "quantity": 1,
}
```
