# Setup a simple redis

Inorder to setup redis a quick and dirty way, you could use provided `Dockerfile` and also you can use [Official Image](https://hub.docker.com/_/redis) which is recomended.

### Docekr commands
```bash
sudo docker pull redis
sudo docker run --name redis1 -p 6379:6379 -d redis
sudo docker run -it --rm redis redis-cli -h 127.0.0.1 -p 6379
```

if you want to set any password into your redis instance use following command

```bash
redis-cli
> CONFIG SET requirepass "your_password_here"
```

And also if you want you could install `redis-cli` locally and connect into running instance with following command

```bash
redis-cli -h host -p port
```

### Connect into redis with `Go`
We are using [offical redis-go client](https://github.com/redis/go-redis) which is an standard and totally up-to-date with internal redis features.

```go
import (
    "context"
    "fmt"

    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ExampleClient() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
    // Output: key value
    // key2 does not exist
}
```

### Start/Stop container
```bash
sudo docker stop redis-1
sudo docker start redis-1
```


