# snowflake-go-demo

An exercise in connecting to a Snowflake DB instance using Golang.

This implementation uses the [snowflakedb/gosnowflake](https://github.com/snowflakedb/gosnowflake)
driver. Documentation is found at [godoc.org](https://godoc.org/github.com/snowflakedb/gosnowflake)

## Connection Parameters

This demo requires a `parameters.json` file in the root directory. It is excluded from source control for obvious
reasons, and should include the following values:

```json
{
    "account":   "",
    "user":      "",
    "password":  "",
    "warehouse": "",
    "database":  "",
    "schema":    ""
}
```

*Note:* In my instance, `account` was the leading portion of the web dashboard URL.
