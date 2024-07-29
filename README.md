# 工具

## resp
> 常用fiber返回

## ctx
> 带过期时间的content,默认为10s

## rands
> 生成指定位数的随机字符串

## security
> 生成token

## httpMiddlewares
> http token 校验中间件

# Client

## PgSql

```go
func TestPgClient(t *testing.T) {
	// 默认连接本地,默认用户名密码和连接数据库为:pgsql
	pg.MustInitPgSqlClient()

	t.Log(client.PgSqlClient().Create(&pg.PgConfig{}).Error)
}

```

## Mysql
```go
func TestMysqlClient(t *testing.T) {
	// 默认连接本地,用户名:root,密码和连接数据库为:mysql
	mysql.MustInitMysqlClient()

	client.MySqlClient().Create(&mysql.MysqlConfig{})
}

```

## Redis
```go
func TestRedisClient(t *testing.T) {
	// 默认连接本地无账号密码
	cache.MustInitRedisClient()

	client.RedisClient.Set(ctx.Timeout(), "key", "value", -1)
}

```

## ESClient

```go
func TestEsClient(t *testing.T) {
	// 默认连接本地无账号密码
	es.MustInitESClient()
	qeury := elastic.NewMatchQuery("title", "First Blog Post")
	source, _ := qeury.Source()
	marshal, _ := json.Marshal(source)
	fmt.Println(string(marshal))

	do, err := client.EsClient.Search("blogs").Query(qeury).Do(context.Background())
	if err != nil {
		panic(err)
	}
	for _, hit := range do.Hits.Hits {
		json, _ := hit.Source.MarshalJSON()
		fmt.Println(string(json))
	}
}
```
## Minio
```go
func init() {
	// 默认连接本地,账号密码为:minio-minio,不开启https
	s3.MustInitMinioClient()
}

func TestCreateBucket(t *testing.T) {
	err := s3.CreateBucket("test")
	if err != nil {
		fmt.Println(err)
	}
}

func TestMinioUpload(t *testing.T) {
	file, _ := os.OpenFile("./a.jpg", os.O_RDWR, 0777)
	stat, _ := file.Stat()
	// 如果数据为byte数组可以用: s3.UploadDataByte()
	err := s3.UploadDataReader(file, stat.Size(), "test", stat.Name(), "image/jpeg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stat.Name())
}

func TestDownload(t *testing.T) {
	data, err := s3.DownloadData("test", "a.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data.Name)
	fmt.Println(data.Size)
	fmt.Println(data.ContextType)
}

func TestRun(t *testing.T) {
	app := fiber.New()
	group := app.Group("/s3")
	group.Get("/*", s3.StartS3FiberHandler("/s3/"))
	app.Listen(":9999")
}
```