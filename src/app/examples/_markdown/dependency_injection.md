## Folder Hierarchy

```
app
 ├─ main.go
 └─ dep
	├─ utils.go
	├─ db.go
	├─ page.go
	└─ counter.templ
```

## main.go

```go

appConfig := utils.AppConfig{
	AppName: "Temporary-Website",
	// examples
	APIKey: "s3cr3t-k3y",
	OtherSettings: map[string]string{
		"featureFlag": "true",
		"maxUsers":    "1000",
	},
}

config := InitConfig(appConfig, "localhost:6379")
t := temp.NewTemp(config) 
```

## utils.go

```go
// Config is a struct that holds configuration settings for the application.
type Config struct {
	AppConfig
	Datastore Datastore
}

type AppConfig struct {
	AppName       string
	APIKey        string
	OtherSettings map[string]string // Additional settings as a map for flexibility
}

// initConfig initializes the configuration with a Database.
func InitConfig(appConfig AppConfig, dbAddr string) Config {
	rdb := NewRedisDB(dbAddr)
	return Config{
		AppConfig: appConfig,
		Datastore: rdb,
	}
} 
```

## db.go

```go
// Datastore is an interface that our Database will implement.
type Datastore interface {
	IncrementCounter(ctx context.Context, key string) (int64, error)
}

// RedisDB wraps the Redis client.
type RedisDB struct {
	client *redis.Client
}

// NewRedisDB creates a new Database instance with a Redis client.
func NewRedisDB(addr string) *RedisDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr, // Redis server address
		Password: "",   // no password set
		DB:       0,    // use default DB
	})
	return &RedisDB{client: rdb}
}

// IncrementCounter increments a number in Redis and returns the updated value.
func (db *RedisDB) IncrementCounter(ctx context.Context, key string) (int64, error) {
	return db.client.Incr(ctx, key).Result()
}
```

## page.go

```go
func Page(w http.ResponseWriter, r *http.Request, config utils.Config) templ.Component {

	newValue, err := config.Datastore.IncrementCounter(r.Context(), "counter")
	if err != nil {
    // handle error
	}

	return server.Counter(config.AppConfig, strconv.FormatInt(newValue, 10))
}
```

## counter.templ

```go
templ Counter(appConf utils.AppConfig, number string) {
	<div class="counter-example">
		<div>There have been: <span>{ number }</span> visitors to this page!</div>
		<h3>AppConfig</h3>
		<div>AppName: <span>{ appConf.AppName }</span></div>
		<div>APIKey: <span>{ appConf.APIKey }</span> (never do this!!!)</div>
		<div>OtherSettings:</div>
		<div>featureFlag: <span>{ appConf.OtherSettings["featureFlag"] }</span></div>
		<div>maxUsers: <span>{ appConf.OtherSettings["maxUsers"] }</span></div>
	</div>
}
```
