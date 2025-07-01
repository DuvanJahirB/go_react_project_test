# Proyecto Test

Este proyecto es una aplicación web completa que consta de un frontend desarrollado con React, un backend en Go y una base de datos MongoDB, todo orquestado para su despliegue con Docker Compose.

## Tabla de Contenidos

- [Frontend](#frontend)
- [Backend](#backend)
- [Despliegue con Docker](#despliegue-con-docker)
  - [Prerrequisitos](#prerrequisitos)
  - [Cómo Desplegar](#cómo-desplegar)
  - [Acceso a los Servicios](#acceso-a-los-servicios)
- [Estructura del Proyecto](#estructura-del-proyecto)
- [Variables de Entorno](#variables-de-entorno)

## Frontend

El frontend de la aplicación es una interfaz de usuario interactiva construida con:

-   **React**: Una biblioteca de JavaScript para construir interfaces de usuario.
-   **Vite**: Un bundler de próxima generación para el desarrollo web.
-   **Tailwind CSS**: Un framework CSS de utilidad para un diseño rápido y personalizado.

### Cómo Ejecutar el Frontend Localmente (sin Docker)

1.  Navega al directorio `frontend`:
    ```bash
    cd frontend
    ```
2.  Instala las dependencias:
    ```bash
    npm install
    ```
3.  Inicia el servidor de desarrollo:
    ```bash
    npm run dev
    ```
    El frontend estará disponible en `http://localhost:3020` (o el puerto que Vite asigne).

## Backend

El backend de la aplicación es una API RESTful desarrollada en Go, encargada de la lógica de negocio y la interacción con la base de datos.

### Cómo Ejecutar el Backend Localmente (sin Docker)

1.  Asegúrate de tener Go instalado en tu sistema.
2.  Navega al directorio `backend`:
    ```bash
    cd backend
    ```
3.  Instala las dependencias de Go:
    ```bash
    go mod tidy
    ```
4.  Ejecuta la aplicación:
    ```bash
    go run cmd/main.go
    ```
    El backend estará disponible en `http://localhost:8088`. Asegúrate de que una instancia de MongoDB esté ejecutándose y accesible en `mongodb://localhost:27017` si lo ejecutas de esta manera.

## Despliegue con Docker

Este proyecto está configurado para un despliegue sencillo utilizando Docker Compose, lo que permite levantar todos los servicios (frontend, backend, MongoDB) con un solo comando.

### Prerrequisitos

Asegúrate de tener Docker y Docker Compose instalados en tu sistema. Puedes seguir las instrucciones en la [documentación oficial de Docker](https://docs.docker.com/get-docker/).

### Cómo Desplegar

1.  Abre una terminal en el directorio raíz del proyecto (`Proyecto Test`).
2.  Para construir las imágenes de Docker y levantar todos los servicios en segundo plano, ejecuta:
    ```bash
    docker-compose up -d --build
    ```
    El flag `--build` asegura que las imágenes se reconstruyan si hay cambios en el código fuente del frontend o backend.

3.  Para verificar el estado de los servicios:
    ```bash
    docker-compose ps
    ```

4.  Para detener y eliminar todos los contenedores, redes y volúmenes creados por `docker-compose`:
    ```bash
    docker-compose down
    ```

### Acceso a los Servicios

Una vez que los servicios estén en ejecución con Docker Compose:

-   **Frontend**: Accede a la aplicación web en tu navegador en `http://localhost:3020`.
-   **Backend**: La API del backend estará disponible en `http://localhost:8088`.
-   **MongoDB**: La base de datos MongoDB estará accesible internamente por el backend y externamente en `mongodb://localhost:27017`.

## Estructura del Proyecto

```
.
├── .gitignore
├── docker-compose.yml         # Configuración de Docker Compose
├── backend/                   # Código fuente del backend (Go)
│   ├── dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── cmd/
│   │   └── main.go
│   └── internal/
│       ├── handlers/
│       ├── middleware/
│       ├── models/
│       └── utils/
├── frontend/                  # Código fuente del frontend (React, Vite, Tailwind CSS)
│   ├── .env
│   ├── Dockerfile
│   ├── package.json
│   ├── src/
│   │   ├── components/
│   │   ├── context/
│   │   └── pages/
│   └── ...
└── mongo-data/                # Volumen persistente para los datos de MongoDB
```

## Variables de Entorno

Este proyecto utiliza variables de entorno para la configuración de los servicios. Estas variables se definen en el archivo `docker-compose.yml` y son inyectadas en los contenedores correspondientes.

### Backend

Las siguientes variables de entorno son utilizadas por el servicio `backend`:

-   `MONGO_URI`: URI de conexión a la base de datos MongoDB. Por defecto, `mongodb://mongo:27017` cuando se ejecuta con Docker Compose.
-   `PORT`: Puerto en el que el servidor backend escuchará las peticiones. Por defecto, `8080`.
-   `JWT_SECRET_KEY`: Clave secreta utilizada para firmar y verificar los tokens JWT. **¡Importante: Cambia este valor por una clave segura en un entorno de producción!**

### Frontend

Las siguientes variables de entorno son utilizadas por el servicio `frontend`:

-   `VITE_BACKEND_URL`: URL base del servicio backend al que el frontend realizará las peticiones. Por defecto, `http://backend:8080` cuando se ejecuta con Docker Compose.
