решение проблемы "data race" с использованием каналов

unsafe_server - пример небезопасного добавления пользователей из разнх горутин
safe_server - пример безопасного добавления пользователей из разнх горутин с использованием каналов
safe_server2 - пример безопасного добавления пользователей и установкой общего сообщения из разных горутин с использованием каналов

протестировать все тесты, с включением обнаружения условия гонги
```shell
go test ./... --race  
```