<p align="center">
  <a href="https://temporary-framework.org">
    <picture>
      <img src="/static/assets/temporary.png" height="128">
    </picture>
    <h1 align="center">Temporary Website</h1>
  </a>
</p>

<p align="center">
  <a aria-label="Temporary logo" href="https://temporary-framework.org">
    <img src="https://img.shields.io/badge/MADE%20BY%20Caleb%20Sideras-000000.svg?style=for-the-badge&logo=Go&labelColor=000">
  </a>
  <a aria-label="License">
    <img alt="" src="https://img.shields.io/npm/l/next.svg?style=for-the-badge&labelColor=000000">
  </a>
</p>

## What is Temporary?

Temporary is a framework designed to make working with HTMX and Templ easier

## Getting Started

Visit <a href="https://temporary-framework.org">https://temporary-framework.org</a> to get started with Temporary.

## Documentation

Visit [https://temporary-framework.org/docs](https://temporary-framework.org/docs) to view the full documentation.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- Bun: [Download and Install Bun](https://bun.sh/)
- Go Programming Language: [Download and Install Go](https://golang.org/doc/install)
- Templ: [Download and Install Templ](https://templ.guide/quick-start/installation)


## Installing

1. Clone the repository to your local machine.

   ```bash
   git clone --recursive https://github.com/caleb-sideras/temporary-framework-website.git
   ```

2. Navigate to the project's directory.

   ```bash
   cd temporary-framework-website
   ```

3. Install the project dependencies.

   ```bash
   bun install
   ```

## Building the Application

Run the Temporary build process through Bun:

  ```bash
  bun run build
  ```

## Running the Application

Run the Temporary start process through Bun:

   ```bash
   bun run start
   ```

## License

This project is licensed under the MIT License.

## TODO

2. Delete/create default files on BUILD - > html/css/js delete 
    - so we dont want to just delete all the files in this dir - simply because users might add shit to this?
    - store files we saved so we can remove them?
    - give a warning to the user about storing files here? 

3. Create code for the handlers instead of the definitions.go??
    - could still have the same process but the routes would be stored in separate package?

4. Fix highlighting of focused color

5. Fix mobile titles length

6. docs -> example dependancy, explain build process

7. remove mux handler

9. metadata api

10. more route level ui

11. nested indexs

12. chunking of js

## REGEX double example

```regex
regex={`^\/projects\/aproject(?:$|\/.*)`}
```
