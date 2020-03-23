# go-todo-list-app

~適当に table を作る~

マイグレーションしてくれるからOK

## precommit

diffのあるファイルのフォーマットとテスト

```sh
ln -s $(pwd)/pre-commit.sh $(pwd)/.git/hooks/pre-commit
```


## Todo 一覧の取得

```shell script
$ curl localhost:8080/tasks
```

## Todo 作成

```shell script
$ curl -X POST -H "Content-Type: application/json" -d '{"title": "todo1", "done": false}' localhost:8080/tasks
```

## Todo 確認

```shell script
$ curl localhost:8080/tasks/:id
```

## Todo 更新

```shell script
$ curl -X PUT -H "Content-Type: application/json" -d '{"title": "todo1 yeah!", "done": true}' localhost:8080/tasks/1
```
