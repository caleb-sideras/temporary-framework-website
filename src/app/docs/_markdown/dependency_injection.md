# Dependency Injection

Dependency Injection in Go with Temporary is a robust method to inject dependencies such as configuration objects, services, or variables into your handlers (`pages`, `indexes`, and `routes`). This eliminates the need for type assertions or other cumbersome workarounds, streamlining the development process.


In the initialization of your Temporary project within `main.go`, you specify the dependencies that will be accessible within your handlers.

### Folder Hierarchy

```
app
 ├─ main.go
 ├─ helpers
 |	└─ utils.go
 └─ dep
	├─ page.go
	└─ page.templ
```

```go
// main.go
func main(){
	
	config := helpers.AppConfig{
		AppName: "Temporary-Website",
		},
	}

	t := temp.NewTemp(config) 
	// invoke Build, Render, Run etc
}
```

Within your handler, you can directly utilize the passed configuration object.

```go
// page.go

func Page(conf helpers.AppConfig) templ.Component {
	return DisplayName(conf.AppName)
}

// page.templ
templ DisplayName(name string){
  <h1>{ name }</h1>
}
```

## Mutations

To modify a dependency, pass it as a pointer. This allows the dependency to be updated within the scope of a request.

```go
// main.go
func main(){
	
	config := helpers.AppConfig{
		AppName: "Temporary-Website",
		},
	}

	t := temp.NewTemp(&config) 
	// invoke Build, Render, Run etc
}
```

```go
// page.go

func Page(w http.ResponseWriter, r *http.Request, conf *helpers.AppConfig) templ.Component {
	conf.AppName = r.PostFormValue("app_name")
	return DisplayName(conf.AppName)
}

// page.templ

templ DisplayName(name string){
  <h1>{ name }</h1>
}
```

##  Database Connection

Integrating database connections as dependencies streamlines database operations within your handlers, ensuring a clean and efficient data management pattern.

### Folder Hierarchy

```
app
 ├─ main.go
 ├─ helpers
 |	├─ utils.go
 |	└─ db.go
 └─ dep
	├─ page.go
	└─ counter.templ
```

### main.go

```go
func main(){
	config := helpers.InitConfig("localhost:6379")
	t := temp.NewTemp(config) 
	// invoke Build, Render, Run etc
}
```

### utils.go

```go
// Config is a struct that holds configuration settings for the application.
type Config struct {
	Datastore Datastore
}

// initConfig initializes the configuration with a Database.
func InitConfig(dbAddr string) Config {
	rdb := NewRedisDB(dbAddr)
	return Config{
		Datastore: rdb,
	}
} 
```

### db.go

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

### page.go

```go
func Page(w http.ResponseWriter, r *http.Request, config helpers.Config) templ.Component {

	newValue, err := config.Datastore.IncrementCounter(r.Context(), "counter")
	if err != nil {
    // handle error
	}

	return server.Counter(config.AppConfig, strconv.FormatInt(newValue, 10))
}
```

### counter.templ

```go
templ Counter(appConf utils.AppConfig, number string) {
		<div>There have been: <span>{ number }</span> visitors to this page!</div>
}
```

## Implementation Details

Temporary optimizes dependency injection by resolving dependencies at build time, ensuring type safety and improving performance by eliminating runtime type assertions.

__NOTE__: The type definition of your dependency must be defined in a separate package to where you initialize your `temp.NewTemp()` 
