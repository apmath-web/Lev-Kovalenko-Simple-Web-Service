# Lev Kovalenko - Simple Web Service

веб сервис позволяющий обрабаьывать запросы следующего вида:
- "/"
- "/date"
- "/wait?delay=XXXX"

для запроса типа "/" возвращается тсрока "hello world"
для запрсоа типа "/date" возвращается дата в формате ISO 8601
для запроса типа "/wait?delay=XXXX", где XXXX - uint64,
создается задержка и возвращется строка: "Sleep for XXXXms completed"