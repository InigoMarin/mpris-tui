# MPRIS TUI

Una aplicación de terminal (TUI) simple escrita en Go para controlar reproductores de medios compatibles con MPRIS en Linux a través de `playerctl`.

## Características

- **Detección automática de reproductores:** Detecta y lista todos los reproductores de medios disponibles al iniciarse.
- **Selección de reproductor:** Permite seleccionar un reproductor de la lista para controlarlo.
- **Controles de reproducción:** Ofrece controles básicos como Play/Pause, Stop, Siguiente y Anterior.
- **Interfaz intuitiva:** Diseñada para ser fácil de usar, incluso en pantallas pequeñas o a través de una conexión SSH desde un dispositivo móvil.
- **Ligera y rápida:** Al estar escrita en Go, la aplicación es un único binario sin apenas dependencias externas, más allá de `playerctl`.

## Requisitos

- **Go:** Necesario para compilar la aplicación.
- **playerctl:** Necesario para la detección y el control de los reproductores. Asegúrate de que esté instalado en tu sistema.

## Instalación y Uso

1.  **Clona o descarga el repositorio.**
2.  **Navega al directorio del proyecto:**
    ```bash
    cd mpris-tui
    ```
3.  **Ejecuta la aplicación directamente:**
    ```bash
    go run .
    ```
4.  **(Opcional) Compila el binario:**
    Para crear un archivo ejecutable único, puedes compilar la aplicación:
    ```bash
    go build .
    ```
    Esto generará un binario llamado `mpris-tui` que puedes mover a cualquier directorio en tu `$PATH` para un acceso global (por ejemplo, `/usr/local/bin`).

## Controles

### Vista de Selección de Reproductor

- **Flechas arriba/abajo:** Navegar por la lista de reproductores.
- **Enter:** Seleccionar un reproductor y pasar a la vista de control.
- **q / Ctrl+c:** Salir de la aplicación.

### Vista de Control

- **p:** Play / Pause
- **s:** Stop
- **n:** Siguiente
- **v:** Anterior
- **b:** Volver a la lista de selección de reproductores.
- **q / Ctrl+c:** Salir de la aplicación.
