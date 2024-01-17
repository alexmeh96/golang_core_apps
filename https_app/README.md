1) (плохо работает, go не может распарсить закрытый ключ)генерация самоподписанного сертификата cert.pem и закрытого ключа key.pem:

openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -subj "/CN=alexmeh966 Root CA" -passout pass:"alex"

2) генерация закрытого ключа server.key:

openssl genrsa -out server.key 2048

генерация самоподписанного сертификата:

openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650 -subj "/CN=alexmeh966 Root CA"


просмотр сертификата в читаемом виде:

openssl x509 -in cert.pem -noout -text


docker: 

сборка

docker build -t test_go_https_app:v1 .

запуск

docker run --name test_go_https_app -p 8081:8081 test_go_https_app:v1

сохранение имеджа test_go_https_app:v1 в tar-файл

docker save --output test_go_https_app_v1.tar test_go_https_app:v1

загрузка имеджа из tar-файла

docker load test_go_https_app_v1.tar