# File Inspector - README

## Descripción

File Inspector es un script escrito en JavaScript y diseñado para ejecutarse en un entorno de Node.js. Su propósito principal es verificar los cambios en una carpeta específica y extraer únicamente los archivos que han cambiado. Esto puede facilitar los despliegues a producción al identificar rápidamente los archivos modificados que necesitan ser actualizados.

## Requisitos previos

Antes de utilizar este script, asegúrate de tener instalado lo siguiente:

- Node.js: Puedes descargar e instalar Node.js desde su sitio web oficial: [https://nodejs.org](https://nodejs.org)

## Uso

Sigue estos pasos para utilizar el script File Inspector:

1. Descarga el archivo `file-inspector.js` en tu entorno de desarrollo.

2. En tu terminal, navega hasta el directorio donde se encuentra el archivo `fileInspector.js`.

3. Modificar la ruta que se encuentra en el archivp y ejecuta el siguiente comando en la terminal:

```bash
node fileInspector.js
```

4. El script comenzará a ejecutarse y analizará la carpeta especificada en busca de cambios. Una vez que finalice la ejecución, se mostrará una lista de los archivos modificados.

## Personalización

El script File Inspector se puede personalizar para adaptarse a tus necesidades específicas. En el archivo `file-inspector.js`, encontrarás las siguientes variables que puedes modificar:

```javascript
const baseDir = '/ruta/a/carpeta'; // Ruta de la carpeta a inspeccionar
const outputDir = `./output-${date.getTime()}`; // Carpeta de salida de archivos
const ignoreDirs = ['.gitignore']; // Carpetas que se ignorarán durante la inspección
const ignoreFile = ['.gitignore']; // Archivos que se ignorarán durante la inspección
```

Ajusta la ruta de la carpeta (`baseDir`) y la lista de archivos ignorados (`gitignore`) según tus requisitos.

## Implementacion en GO (Experimental)

El concepot de programacion es similar, pero se hace una implementacion e go para evaluar el rendimineto y ver si exixte mejora con respecto a un lenguaje interpretado, acutualmente se encuentra en fase experimental

## Contribuciones

Las contribuciones son bienvenidas. Si encuentras algún error, tienes alguna idea para mejorar el script o deseas agregar nuevas características, no dudes en abrir un problema o enviar una solicitud de extracción.

## Licencia

Este proyecto está licenciado bajo la [Licencia GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html). Si utilizas este script en tu propio proyecto, agradecemos un enlace o atribución al repositorio original.

---

¡Disfruta utilizando File Inspector! Si tienes alguna pregunta o necesitas ayuda, no dudes en comunicarte con nosotros.