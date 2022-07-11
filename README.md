### Config example

```
{
    "proxy_server_addr":"127.0.0.1:8080",
    "iots_devices":[
        {
        "id":0,
        "name":"iot_0",
        "base_number":0,
        "step":2,
        "formulas": {
        "formula_1":"2*base_number"
        }

        }
    ]
}
```

Projects use lib https://github.com/Knetic/govaluate for formulas. U can add 5 formulas: formula_1-formula_5. The
base_number is sinonym for x. Step use for increse the base_number(x) when u call func(formula). Id of iot device must
be unic. 
