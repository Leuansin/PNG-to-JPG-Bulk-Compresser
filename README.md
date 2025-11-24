# 🖼️ PNG to JPG Converter

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.16+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey?style=for-the-badge)

**A fast, simple, and efficient command-line tool to batch convert PNG images to JPG format**

[English](#english) • [Español](#español)

</div>

---

<a name="english"></a>

## 🌟 Features

✨ **Batch Processing** - Convert entire directories of PNG files in one go  
🎨 **Smart Transparency Handling** - Automatically applies white background to transparent areas  
⚙️ **Customizable Quality** - Control JPG compression from 1 to 100  
📁 **Preserves Directory Structure** - Maintains your original folder organization  
🚀 **Zero Dependencies** - Pure Go implementation, no external libraries needed  
💻 **Cross-Platform** - Works seamlessly on Linux, macOS, and Windows

## 🚀 Quick Start

### Prerequisites

- Go 1.16 or higher installed on your system

### Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/yourusername/png2jpg-converter.git
cd png2jpg-converter
go build -o png2jpg
```

### Basic Usage

Place your PNG files in an `inputs` folder and run:

```bash
./png2jpg
```

Your converted JPG files will appear in the `outputs` folder.

## ⚙️ Configuration Options

| Flag | Default | Description |
|------|---------|-------------|
| `-in` | `./inputs` | Input directory containing PNG files |
| `-out` | `./outputs` | Output directory for JPG files |
| `-q` | `80` | JPG quality (1-100, recommended: 70-85) |

### Examples

**Convert with custom quality:**
```bash
./png2jpg -q 90
```

**Specify custom directories:**
```bash
./png2jpg -in /path/to/pngs -out /path/to/jpgs
```

**Low quality for web optimization:**
```bash
./png2jpg -in ./images -out ./compressed -q 65
```

## 🎯 Use Cases

- 📸 **Photography Workflows** - Batch process RAW to JPG conversions
- 🌐 **Web Optimization** - Reduce file sizes for faster loading
- 📦 **Asset Management** - Standardize image formats across projects
- 💾 **Storage Reduction** - Save disk space without quality loss

## 🛠️ How It Works

The converter uses Go's native image processing libraries to:

1. 📂 Recursively scan the input directory for PNG files
2. 🎨 Decode each PNG image and handle transparency
3. ⚪ Apply a white background to transparent regions
4. 💾 Encode to JPG format with your specified quality
5. 📁 Save to the output directory maintaining folder structure

## 📋 Requirements

- **Operating System:** Linux, macOS, Windows
- **Go Version:** 1.16+
- **Memory:** Minimal (scales with image size)
- **Disk Space:** Sufficient for output files

## 🤝 Contributing

Contributions are welcome! Feel free to:

- 🐛 Report bugs
- 💡 Suggest new features
- 🔧 Submit pull requests
- 📖 Improve documentation

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

Built with ❤️ using Go's powerful standard library

---

<a name="español"></a>

## 🌟 Características

✨ **Procesamiento por Lotes** - Convierte directorios completos de archivos PNG de una vez  
🎨 **Manejo Inteligente de Transparencias** - Aplica automáticamente fondo blanco a áreas transparentes  
⚙️ **Calidad Personalizable** - Controla la compresión JPG de 1 a 100  
📁 **Preserva Estructura de Directorios** - Mantiene tu organización de carpetas original  
🚀 **Cero Dependencias** - Implementación pura en Go, sin librerías externas  
💻 **Multiplataforma** - Funciona perfectamente en Linux, macOS y Windows

## 🚀 Inicio Rápido

### Requisitos Previos

- Go 1.16 o superior instalado en tu sistema

### Instalación

Clona el repositorio y compila el binario:

```bash
git clone https://github.com/yourusername/png2jpg-converter.git
cd png2jpg-converter
go build -o png2jpg
```

### Uso Básico

Coloca tus archivos PNG en una carpeta `inputs` y ejecuta:

```bash
./png2jpg
```

Tus archivos JPG convertidos aparecerán en la carpeta `outputs`.

## ⚙️ Opciones de Configuración

| Flag | Por Defecto | Descripción |
|------|-------------|-------------|
| `-in` | `./inputs` | Directorio de entrada con archivos PNG |
| `-out` | `./outputs` | Directorio de salida para archivos JPG |
| `-q` | `80` | Calidad JPG (1-100, recomendado: 70-85) |

### Ejemplos

**Convertir con calidad personalizada:**
```bash
./png2jpg -q 90
```

**Especificar directorios personalizados:**
```bash
./png2jpg -in /ruta/a/pngs -out /ruta/a/jpgs
```

**Baja calidad para optimización web:**
```bash
./png2jpg -in ./imagenes -out ./comprimidas -q 65
```

## 🎯 Casos de Uso

- 📸 **Flujos de Fotografía** - Procesa conversiones por lotes de RAW a JPG
- 🌐 **Optimización Web** - Reduce tamaños de archivo para carga más rápida
- 📦 **Gestión de Assets** - Estandariza formatos de imagen en proyectos
- 💾 **Reducción de Almacenamiento** - Ahorra espacio en disco sin perder calidad

## 🛠️ Cómo Funciona

El conversor utiliza las librerías nativas de procesamiento de imágenes de Go para:

1. 📂 Escanear recursivamente el directorio de entrada buscando archivos PNG
2. 🎨 Decodificar cada imagen PNG y manejar transparencias
3. ⚪ Aplicar un fondo blanco a las regiones transparentes
4. 💾 Codificar al formato JPG con la calidad especificada
5. 📁 Guardar en el directorio de salida manteniendo la estructura de carpetas

## 📋 Requisitos

- **Sistema Operativo:** Linux, macOS, Windows
- **Versión de Go:** 1.16+
- **Memoria:** Mínima (escala con el tamaño de imagen)
- **Espacio en Disco:** Suficiente para los archivos de salida

## 🤝 Contribuciones

¡Las contribuciones son bienvenidas! Siéntete libre de:

- 🐛 Reportar bugs
- 💡 Sugerir nuevas características
- 🔧 Enviar pull requests
- 📖 Mejorar la documentación

## 📄 Licencia

Este proyecto está licenciado bajo la Licencia MIT - consulta el archivo LICENSE para más detalles.

## 🙏 Agradecimientos

Construido con ❤️ usando la poderosa librería estándar de Go

---

<div align="center">

**Made with 💙 by [Your Name]**

⭐ Star this repo if you find it useful!

</div>
