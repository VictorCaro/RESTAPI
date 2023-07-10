# RESTAPI

El proyecto es una API REST desarrollada en Go que utiliza el enrutador Gorilla Mux y la biblioteca ORM Gorm para interactuar con una base de datos PostgreSQL. Su propósito principal es proporcionar endpoints para la gestión de usuarios y tareas.

La API permite realizar operaciones CRUD (Crear, Leer, Actualizar, Eliminar) en la base de datos para los usuarios y las tareas. Los endpoints permiten obtener la lista de usuarios, obtener información detallada de un usuario específico, crear nuevos usuarios, eliminar usuarios y realizar acciones similares para las tareas.

# Requisitos previos.

Go (versión X.X.X)
PostgreSQL (versión X.X.X)
Instalación

Clona este repositorio: git clone https://github.com/VictorCaro/RESTAPI
Accede a la carpeta del proyecto: cd Proyecto
Ejecuta el comando para instalar las dependencias: go mod tidy

# Configuracion

Abre el archivo db/connection.go.
Modifica la variable DSN con la configuración de tu base de datos PostgreSQL (host, usuario, contraseña, nombre de la base de datos, puerto).
Guarda los cambios.
Uso

Ejecuta el comando para iniciar el servidor: go run main.go.
La API estará disponible en http://localhost:3000.

# Endpoints

GET /users: Obtiene la lista de usuarios.
GET /users/{id}: Obtiene un usuario específico.
POST /users: Crea un nuevo usuario.
DELETE /users/{id}: Elimina un usuario.
GET /tasks: Obtiene la lista de tareas.
GET /tasks/{id}: Obtiene una tarea específica.
POST /tasks: Crea una nueva tarea.
DELETE /tasks/{id}: Elimina una tarea.
Contribución

Este proyecto es una muestra de trabajo personal y no acepta contribuciones en este momento.

# Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo LICENSE para más detalles.

# Contacto

Si tienes alguna pregunta o comentario, puedes contactarme en victor.caro.rainao@gmail.com.
