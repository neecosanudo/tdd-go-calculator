# TDD Go Calculator

El objetivo de este proyecto es desarrollar una calculadora en lenguaje Go (Golang) siguiendo los principios de TDD (Test-Driven Development). El proyecto se desarrolla de manera incremental, agregando funcionalidades de cálculo una a una y utilizando pruebas unitarias para asegurar que el código está funcionando correctamente en cada etapa.

## Índice

La documentación consta de dos partes:

[**Parte I - Bitácora**](#bitácora)

1. [Desarrollo guiado por pruebas](#desarrollo-guiado-por-pruebas)
2. [Sobre Git y el flujo de trabajo que probé](#sobre-git-y-el-flujo-de-trabajo-que-probé)
    1. [Mensajes de commit](#mensajes-de-commit)
    2. [Comienzo, creación de ramas y merge de ramas](#comienzo-creación-de-ramas-y-merge-de-ramas)
3. [ChatGPT como herramienta de apoyo](#chatgpt-como-herramienta-de-apoyo)

[**Parte II - Documentación**](#documentación)

1. [¿Qué es el desarrollo guiado por pruebas?](#qué-es-el-desarrollo-guiado-por-pruebas)
2. [Archivos](#archivos)
3. [Funciones](#funciones)
    1. [Suma](#suma)
    2. [Resta](#resta)
4. [Como usar](#como-usar)
    1. [Contenido de main.go](#contenido-de-maingo)
    2. [Compilación y ejecución](#compilación-y-ejecución)

## Bitácora

Mi idea con este proyecto es poner en práctica el flujo de trabajo aprendido en [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/) y algunas conceptos de *workflow* con Git.

Mi primer decisión fue no aplicar conceptos de Golang que no necesite, por eso el código es muy simple. Este enfoque me permite concentrarme en adquirir hábitos de programación más que en reforzar el aprendizaje de conceptos propios de Go.

### Desarrollo guiado por pruebas

El desarrollo guiado por pruebas que explica [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests/) recomienda seguir el siguiente ciclo:

1. Escribir el test.
2. Hacer que el compilador pase el test: escribir lo mínimo y necesario.
3. Correr el test, ver que falló y verificar si el mensaje del error es significativo.
4. Escribir el código para que el test pase.
5. Refactor.

### Sobre Git y el flujo de trabajo que probé

Luego de leer la documentación oficial y algunos blogs sobre Git, armé un flujo de trabajo que me resulta cómodo y quería probar su productividad.

#### Mensajes de commit

Los **mensajes de commit** siguen el patrón recomendado por [*Conventionals Commits*](https://www.conventionalcommits.org/es/v1.0.0-beta.3/), agregando los emojis sugeridos por [*EmojiDev*](https://gitmoji.dev/) al inicio.

```bash
  🚧 FEAT: refactoring Resta() function to accept multiple numbers
```

#### Comienzo, creación de ramas y *merge* de ramas

El primer commit del repositorio contiene los archivos:

- `README.md`: bitácora y documentación del código.
- `.gitignore`: lo que Git no va a llevar registro en el historial. Creado utilizando [toptal.com/developers/gitignore](https://www.toptal.com/developers/gitignore).
- `LICENSE`: licencia del repositorio. Creado utilizando [*Choose a License*](https://choosealicense.com/).
- `go.mod`: especificación del módulo del proyecto.

Las ramas del repositorio son:

- `main`: rama principal.
- `features`: nace del primer commit de `main` y de ella se desprenden las ramas de caracteristicas.
- `documentation`: se desprende de `main` y en ella trabajo la documentación del proyecto.
- Ramas de características como `suma`, `resta`, `execute-operations` y `refactor`. Se desprenden de `features`.
- `temporal`: es una rama que utilizo para hacer un *merge* y luego la elimino.

El flujo de trabajo que puse en práctica consiste en:

1. Creo el repositorio con algunos archivos esenciales.
2. Creo la rama `features` para luego desprender futuras ramas en base a la característica a trabajar.
3. Desarrollo el código de la característica actual siguiendo el desarrollo guiado por pruebas (*Test Driven Development*).
4. Por cada prueba aprobada, realizo un commit siguiendo el [patrón de mensajes de commit](#mensajes-de-commit) que especifiqué antes.
5. Realizo el proceso de *refactor* y *agregar comentarios en código*, haciendo luego su correspondiente commit.

Para explicar el proceso de *merge* de ramas voy a utilizar de ejemplo un *merge* entre `features` y `suma`:

1. Creo una nueva rama a partir repositorio de `suma` utilizando el comando `git checkout -b temporal`.
2. Desde `features` hago un `git merge --squash temporal`. Esto lleva todos los commits de `temporal` como 1 solo cambio nuevo en la rama `features` sin perder el historial de commits en `suma`.
3. Realizo el commit en `features` con los cambios agregados.
4. Elimino la rama `temporal`.

**Nota**: el procedimiento está en un ciclo de análisis y desarrollo. No están implementados en esta explicación comandos como `git reset` y el uso de ramas como `hotfix`.

### ChatGPT como herramienta de apoyo

[ChatGPT](https://openai.com/blog/chatgpt/) fue una herramienta importante para mejorar varios aspectos de mi código porque me ayudó en puntos donde estoy mejorando. Dentro de lo que fue "*conversado*" con ChatGPT está:

1. **Nombres**: el nombre del repositorio, nombre de las ramas de Git y el nombre de la función `assertMessage()`.
2. **Commits**: los mensajes de los commits tenían que ser más cortos y concisos. Además de incluir un *scope* y un emoji al comienzo.
3. Aspectos técnicos de la documentación como la definición de TDD en las sección [¿Qué es el desarrollo guiado por pruebas?](#¿qué-es-el-desarrollo-guiado-por-pruebas)

## Documentación

Este es un proyecto en Golang que implementa funciones básicas de una calculadora como suma y resta, siguiendo el desarrollo guiado por pruebas.

### Qué es el desarrollo guiado por pruebas

También conocido como Test Driven Development (TDD), es un enfoque de desarrollo de software en el cual las pruebas son escritas antes de escribir el código de la funcionalidad.

En lugar de escribir primero el código y luego escribir las pruebas, en TDD se escriben las pruebas primero y luego se escribe el código para hacerlas pasar. El proceso se repite cada vez que se agrega nueva funcionalidad.

Esto ayuda a garantizar que el código cumpla con los requisitos y sea de alta calidad, ya que se prueba continuamente mientras se desarrolla.

### Archivos

- `operaciones.go`: contiene las funciones de `Suma()` y `Resta()`.
- `operaciones_test.go`: contiene las pruebas unitarias para las funciones de suma y resta.
- `main.go`: Contiene el código para ejecutar las funciones de suma y resta en un archivo principal.

### Funciones

#### Suma

`Suma(numeros ...int) int` toma una cantidad variable de números y regresa la suma de todos ellos.

```go
func Suma(numeros ...int) int {
 resultado := 0

 for _, num := range numeros {
  resultado += num
 }

 return resultado
}
```

#### Resta

`Resta(numeros ...int) int` toma una cantidad variable de números y regresa el resultado de restar los elementos en orden consecutivo.

```go
func Resta(numeros ...int) int {
 resultado := numeros[0]

 for _, num := range numeros {
  if num != numeros[0] {
   resultado -= num
  }
 }

 return resultado
}
```

### Como usar

#### Contenido de main.go

```go
func main() {

 const (
  a = 20
  b = -7
  c = 15
  d = 89
 )

 resultadoSuma := Suma(a, b, c, d)
 resultadoResta := Resta(a, b, c, d)

 fmt.Printf("Dado los números %d, %d, %d y %d:\nLa suma de todos es: %d.\nLa resta de todos es: %d.\n", a, b, c, d, resultadoSuma, resultadoResta)
}
```

El código es un programa escrito en Go que se ejecuta en la función `main()`.

Dentro de `main()` se declaran 4 variables constantes: a, b, c y d.

Luego se asigna el resultado de la función `Suma()` a la variable `resultadoSuma`, utilizando las variables a, b, c y d como parámetros.

Se asigna el resultado de la función `Resta()` a la variable `resultadoResta`, utilizando las mismas variables a, b, c y d como parámetros.

Finalmente, se imprime en consola el resultado de las operaciones de suma y resta utilizando las variables a, b, c y d como parámetros.

#### Compilación y ejecución

Para compilar y ejecutar el código en Linux, Windows y MacOS, debes seguir los siguientes pasos:

1. Asegúrate de tener instalada la última versión de Go en tu computadora. Puedes descargarla desde el [sitio oficial](<https://golang.org/dl/>).

2. Abre una terminal en el directorio donde guardas tus proyectos y ejecuta el comando `git clone https://github.com/neecosanudo/tdd-go-calculator`.

3. Ingresa al directorio del proyecto con el comando `cd tdd-go-calculator`.

4. Ejecuta el comando `go build main.go operaciones.go` para compilar el código. El comando anterior generará un archivo ejecutable con el nombre del paquete (main en este caso) en la carpeta actual.

5. Para ejecutar el programa, escribe el nombre del archivo ejecutable generado en la consola y presiona enter. En este caso sería `./main`.

6. Si todo sale bien, se ejecutará el programa y verás el resultado en la terminal:

```bash
# Salida en consola
Dado los números 20, -7, 15 y 89:
La suma de todos es: 117.
La resta de todos es: -77.
```

**Nota**: si estas en Windows, el comando para ejecutar el archivo ejecutable es `main.exe`.
