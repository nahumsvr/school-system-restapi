# Sistema de Control Escolar

**Proyecto final** del taller **Fundamentos de Go: Desarrollo de aplicaciones de alto rendimiento**.

Este sistema escolar incluye tres endpoints que permiten realizar operaciones **CRUD** (crear, leer, actualizar y eliminar) sobre los siguientes elementos:

* Estudiantes
* Materias
* Calificaciones

---

## Requisitos

Antes de comenzar, asegúrate de tener lo siguiente instalado en tu sistema:

* [Go](https://golang.org/doc/install)
* [Docker](https://www.docker.com/products/docker-desktop)

---

## Instrucciones de ejecución

1. **Ejecuta el siguiente comando en tu terminal para crear el contenedor de PostgreSQL:**

   ```bash
   docker run --name restapi-postgres \
     -e POSTGRES_USER=nahum \
     -e POSTGRES_PASSWORD=nahumpassword \
     -p 5432:5432 \
     -d postgres
   ```

2. **Accede al contenedor:**

   ```bash
   docker exec -it restapi-postgres bash
   ```

3. **Ingresa a la consola de PostgreSQL:**

   ```bash
   psql -U nahum --password
   ```

   La contraseña es:

   ```
   nahumpassword
   ```

4. **Crea la base de datos:**

   ```sql
   CREATE DATABASE gorm;
   ```

5. **En una nueva terminal, navega a la carpeta del repositorio del proyecto y ejecuta:**

   ```bash
   go run main.go
   ```

---

