# go-todo-list-app

適当に table を作る

## Todo 一覧の取得

```shell script
$ curl localhost:8080/tasks
```

## Todo 作成

```shell script
$ curl -X POST -H "Content-Type: application/json" -d '{"title": "todo1", "done": false}' localhost:8080/tasks
```
