# Gin APi Test


## 1. Install
Install TaskFile:
 ```bash
 sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d
 
 ## For MacOS only
 brew install go-task/tap/go-task
 ```
For more info about TaskFile, visit: https://taskfile.dev/#/installation

### Install project
```bash
task tidy # go mod tidy
```

## 2. Commands
```yaml
## Go
task tidy # go mod tidy
task run # go run . 
task test # GIN_MODE=test go test ./test
task build # go build

## App
# Générate Swagger documentation
task doc # swag init --parseDependency --parseInternal

# Add fixtures
task  fixtures # go run . fixtures

# Dump fixtures
task dump # go run . dump 

## Docker
task up # docker compose up -d
task down # docker compose down
task restart # task down + task up
```

## 3. Stack
### Packages
- gin-gonic/gin for http server
- swaggo/gin-swagger to generate documentation
- gorm.io/gorm for entity management
- golang-jwt/jwt/v5 and icahParks/keyfunc/v3 for jwt/kwks management
- go-testfixtures/testfixtures/v3 to generate fixtures
- stretchr/testify and h2non/gock for testing
<br/>

### Features
Each **$${\color{#BF4342}protected}$$** route **$${\color{#BF4342}must have middleware.SecurityMiddleware()}$$**
<br/>
On each request to API, if route is protected, request must have header <span style="color:#BF4342; font-weight:bold">Authorization</span> (Bearer jwt token) or <span style="color:#BF4342; font-weight:bold">X-API-Key</span> 

## 4. Databases and entities
#### 1. Databases
If you want to use one or multiple databases, you have to update config/database.yml file. "default" entry is for the default database :
```yaml
database:
  default:
    adapter: postgres
    encoding: utf8
    dsn: env("POSTGRES_DSN")
    pool: 5
    timeout: 5000
    migrate: true
    config:
      log_level: silent
      skip_default_transaction: true
      disable_nested_transaction: true
      create_batch_size: 5000
      full_save_associations: true
    entities:
      - "entity.Model"
      - "entity.MediaObject"
```
The options dsn and adapter are required for each connection. 
#### 2. Migrations
If you have to play migrations, register path of the entity in the database.yml file and to register init func into your struct file : 
```go

func init() {
	RegisterType("package.EntityName", func() interface{} {
		return &EntityName{}
	})
}
```


## 4. Pagination and Filters
#### 1. Pagination and items per page
You can paginate results of a query by setting the query parameter $${\color{#BF4342}itemsPerPage}$$ and the page number $${\color{#BF4342}page}$$.
Default values are 

- $${\color{#BF4342}itemsPerPage}$$ = 20
- $${\color{#BF4342}page}$$ = 1


#### 2. Filters
You can filter results of a query by setting parameters in the query. 
To filter by a field, you have to set $${\color{#BF4342}filter}$$ tag in struct and set the name of the field in $${\color{#BF4342}json}$$ tag.

Every filter value which is a string is converted to lowercase.

```go
 type Model struct {
    ID              *uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id,omitempty" filter:"id"`// in the query is id=5C6C5C6C-5C6C-5C6C-5C6C-5C6C
    Name            string        `json:"name,omitempty" filter:"name"` // in the query is name=The Name
    Title           *string       `json:"title,omitempty" filter:"title"` // in the query is title=The Title 
    SubTitle        *string       `json:"sub_title,omitempty"` // This field is not filterable because no filter tag is set
    Description     *string       `json:"description,omitempty"` // This field is not filterable because no filter tag is set
    IsNew           bool          `json:"is_new,omitempty" filter:"isNew"` // in the query is isNew=true
}
 ```

Example:
```sh
curl --location --request GET 'https://example.com/models?title=The Title&name=The Name'

 ```
## 5. Swagger documentation
Package documentation : https://github.com/swaggo/swag

Swagger documentation is generated automatically by command `task doc` based on comments in the source code.
For example :
```go
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			Model
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Model ID"
//	@Success		200	{object}	entity.Model
//	@Failure		400	{object}	ApiError
//	@Failure		404	{object}	ApiError
//	@Failure		500	{object}	ApiError
//	@Router			/models/{id} [get]
func GetModelById(c *gin.Context) {
model, _ := repository.GetModelById(handleQuery(c))
c.IndentedJSON(http.StatusOK, model)
}
```
Results by : 
![img.png](doc%2Fimg.png)

## 6. Fixtures
Package documentation : https://github.com/go-testfixtures/testfixtures


Before each test, fixtures are loaded from the directory $${\color{#BF4342}fixtures/ORM}$$ and imported in the test database.
By default each fixture file has the same name as the table it represents. One table by file.
## 6. Mock external web service
#### 1. Add new interceptor
```go
//main_test.go
func (s *WebServiceGinSuite) SetupTest() {
	....
	gock.New("https://example.com").
	Get("/endpoint").
        Persist().
        Reply(200).
        BodyString("Expected Response")	
	....
}
```
<br/>

#### 7. Add test
```go
//main_test.go
func (s *WebServiceGinSuite) TestEndpointHandler() {
    s.T().Run("Health Check 2", func(t *testing.T) {
    req, _ := http.NewRequest("GET", "https://example.com/endpoint", nil)
    res, err := (&http.Client{}).Do(req)
    body, _ := ioutil.ReadAll(res.Body)
    assert.Nil(s.T(), err)
    assert.Equal(s.T(), res.StatusCode, 200)
    assert.Equal(t, body, "Expected Response")
    })
}
```