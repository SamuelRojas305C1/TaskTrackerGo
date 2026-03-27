# TaskTrackerGo

TaskTrackerGo es una aplicación de línea de comandos (CLI) escrita en Go para gestionar tareas personales. Incluye un sistema de autenticación de usuarios que permite a cada usuario tener su propia lista de tareas separada.

## Características

- **Autenticación de usuarios:** Registro, inicio de sesión y cierre de sesión.
- **Gestión individual:** Cada usuario tiene su propio registro y almacenamiento de tareas.
- **Gestión de tareas:** Crear, listar, actualizar y eliminar tareas.
- **Estados de tareas:** Las tareas se gestionan con los estados `Pendiente`, `En Curso` y `Hecho`.
- **Almacenamiento local:** Los datos se guardan en formato JSON localmente (`users.json`, `session.json`, y archivos separados por usuario en la carpeta `data/`).
- **Seguridad:** Las contraseñas de los usuarios se almacenan cifradas utilizando bcrypt.

## Requisitos

- [Go](https://go.dev/dl/) instalado.

## Instalación y Compilación

1. Abre una terminal y navega hasta el directorio del proyecto `TaskTrackerGo`.
2. Asegúrate de instalar las dependencias necesarias:
   ```bash
   go mod tidy
   ```
3. Compila el código para generar el ejecutable:
   ```bash
   go build -o tasktrackergo.exe
   ```
   *(En Linux o macOS usa `go build -o tasktrackergo` y ejecuta con `./tasktrackergo`)*

## Uso

El formato general para llamar a la aplicación es:

```bash
tasktrackergo <comando> [argumentos]
```
*(Si no lo has compilado y usas Go directamente, puedes ejecutar `go run . <comando> [argumentos]`)*

### Autenticación

Para usar la gestión de tareas, primero necesitas tener una cuenta e iniciar sesión.

- **Registrar un usuario:**
  ```bash
  tasktrackergo register <email> <password> <nombre completo>
  ```
  Ejemplo: `tasktrackergo register juan@email.com 123456 Juan Perez`

- **Iniciar sesión:**
  ```bash
  tasktrackergo login <email> <password>
  ```

- **Cerrar sesión:**
  ```bash
  tasktrackergo logout
  ```

### Tareas

*Nota: Estos comandos requieren haber iniciado sesión previamente.*

- **Agregar tarea:**
  ```bash
  tasktrackergo add <descripción>
  ```
  Ejemplo: `tasktrackergo add Comprar manzanas`

- **Listar tareas:**
  Puedes listar todas tus tareas, o filtrarlas por su estado (`Pendiente`, `En Curso`, `Hecho`).
  ```bash
  tasktrackergo list
  tasktrackergo list "Pendiente"
  ```

- **Actualizar descripción:**
  ```bash
  tasktrackergo update <id_tarea> <nueva descripción>
  ```
  Ejemplo: `tasktrackergo update 1 Comprar manzanas y peras`

- **Eliminar tarea:**
  ```bash
  tasktrackergo delete <id_tarea>
  ```

- **Marcar como "En Curso":**
  ```bash
  tasktrackergo mark-in-progress <id_tarea>
  ```

- **Marcar como "Hecho":**
  ```bash
  tasktrackergo mark-done <id_tarea>
  ```

## Estructura del Código

- `main.go`: Punto de entrada de la aplicación y enrutamiento principal.
- `commands.go`: Lógica detallada para cada uno de los comandos (add, delete, register, etc.) y parseo de argumentos (flags).
- `auth.go`: Funciones de registro, login, logout usando validaciones y hashing (`bcrypt`).
- `task.go / user.go`: Definiciones del modelo de datos (`Task`, `User`, `Session`, y estados).
- `taskservice.go`: Clase (`TaskManager`) para la lógica de administrar las tareas en memoria antes de guardado.
- `storage.go`: Manejo del almacenamiento en el sistema de archivos (cargar y guardar los archivos `.json`).
