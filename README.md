
### 作成手順

* ディレクトリ作成

```
mkdir $GOPATH/src/github.com/tosashimanto/heroku-go-matsumoto-test
cd $GOPATH/src/github.com/tosashimanto/heroku-go-matsumoto-test/

```


* main.go作成

* Procfile作成
web: heroku-go-matsumoto-test

vendor/vendor.json作成
```
{
        "rootPath": "github.com/tosashimanto/heroku-go-matsumoto-test"
}

```


```
git init
git add .
git commit -m "initial commit"
```

### Heroku側

```
heroku login
```


* Heroku側にappを作成

```
heroku apps:create heroku-go-matsumoto-test --buildpack heroku/go
heroku open --app heroku-go-matsumoto-test
```


* Herokuにデプロイ
```
git push heroku master

git remote -v
heroku	https://git.heroku.com/heroku-go-matsumoto-test.git (fetch)
heroku	https://git.heroku.com/heroku-go-matsumoto-test.git (push)


```


* ログ確認
```
heroku logs --tail
```

* アクセス例
```
https://rocky-castle-914423.herokuapp.com/gsurvey_api/v1/constructions

```
