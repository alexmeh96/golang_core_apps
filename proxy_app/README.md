### Ssl

генерация закрытого ключа server.key:
```shell
openssl genrsa -out server.key 2048
```
генерация самоподписанного сертификата:
```shell
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650 -subj "/CN=alexmeh96 Root CA"
```
просмотр сертификата в читаемом виде:
```shell
openssl x509 -in cert.pem -noout -text
```

### Docker 

сборка
```shell
docker build -t test_go_https_app:v1 .
```
запуск
```shell
docker run --name test_go_https_app -p 8081:8081 test_go_https_app:v1
```
сохранение имеджа test_go_https_app:v1 в tar-файл
```shell
docker save --output test_go_https_app_v1.tar test_go_https_app:v1
```
загрузка имеджа из tar-файла
```shell
docker load test_go_https_app_v1.tar
```