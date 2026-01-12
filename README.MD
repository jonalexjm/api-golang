yourproject/
|---main.go
|---internal/
|   |---model/
|   |   |--book.go
|   |---store/  -> comunicarnos con base de datos
|   |   |--book_store.go
|   |---service/ -> añadir libros, remover libros
|   |   |--book_service.go
|   |---transport/ -> get, put, post
|   |   |--book_handler.go


model -> Define nuestras estructuras
store -> maneja la base de datos(repository pattern)
service -> contiene la logica de negocio
transport -> Maneja HTTP request/responses
