# CRUD Server

Este proyecto es una API básica con arquitectura hexagonal para gestionar usuarios utilizando Go y GORM con PostgreSQL.

## Requisitos Previos

Antes de comenzar, asegúrate de tener instalados:

- [Go](https://golang.org/) (versión 1.23 o superior)
- [PostgreSQL](https://www.postgresql.org/)

## Instalación

1. Clona el repositorio:

   ```bash
   git clone https://github.com/eduardoquispe/crud_server.git

   cd crud_server
   ```

2. Instala las dependencias del proyecto:

   ```bash
   go mod tidy
   ```

3. Crea una base de datos en PostgreSQL, el nombre de la base de datos puede ser `crud_db`.

4. Configura las variables de entorno:
   Crea un archivo `.env` en la raíz del proyecto y define las variables necesarias. Por ejemplo:

   ```
   APP_PORT=8080

   DATABASE_HOST=localhost
   DATABASE_PORT=5432
   DATABASE_USER=postgres
   DATABASE_PASSWORD=123456
   DATABASE_NAME=crud_db
   ```

## Ejecución

1. Asegúrate de que la base de datos esté en ejecución.

2. Ejecuta el proyecto:

   ```bash
   go run main.go
   ```

3. La aplicación estará disponible por defecto en:

   ```
   http://localhost:8080
   ```

4. Para acceder a la API, puedes utilizar herramientas como [Postman](https://www.postman.com/)

5. La API consta de las siguientes rutas:

   - GET `/api/v1/user`: Devuelve una lista de usuarios.

     ```bash
     **Request**

      No hay parámetros de solicitud.

     **Response**

     {
       "success": true,
       "message": "Users listed successfully",
       "data": [
         {
           "id": 1,
           "firstName": "John",
           "lastName": "Doe",
           "email": "john.doe@example.com",
           "age": 30
         },
         {
           "id": 2,
           "firstName": "Jane",
           "lastName": "Smith",
           "email": "jane.smith@example.com",
           "age": 25
         }
       ]
     }
     ```

   - GET `/api/v1/user/{id}`: Devuelve un usuario específico.

     ```bash
     **Request**

     No hay parámetros de solicitud.

     **Response**

     {
     "success": true,
     "message": "User found successfully",
     "data": {
        "id": 1,
        "firstName": "John",
        "lastName": "Doe",
        "email": "john.doe@example.com",
        "age": 30
     }
     }
     ```

   - POST `/api/v1/user`: Crea un nuevo usuario.

     ```bash
     **Request**

     {
        "firstName": "John",
        "lastName": "Doe",
        "email": "john.doe@example.com",
        "age": 30
     }

     **Response**

     {
        "success": true,
        "message": "User created successfully",
        "data": {
           "id": 1,
           "firstName": "John",
           "lastName": "Doe",
           "email": "john.doe@example.com",
           "age": 30
        }
     }
     ```

   - PUT `/api/v1/user/{id}`: Actualiza un usuario.

     ```bash
     **Request**

     {
        "firstName": "John",
        "lastName": "Doe",
        "email": "john.doe@example.com",
        "age": 30
     }

     **Response**

     {
        "success": true,
        "message": "User updated successfully",
        "data": {
           "id": 1,
           "firstName": "John",
           "lastName": "Doe",
           "email": "john.doe@example.com",
           "age": 30
        }
     }
     ```

   - DELETE `/api/v1/user/{id}`: Elimina un usuario.

     ```bash
     **Request**

     No hay parámetros de solicitud.

     **Response**

     {
        "success": true,
        "message": "User deleted successfully",
        "data": null
     }
     ```
