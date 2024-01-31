console.time('Execution time')

const fs = require('fs');
const path = require('path');
const crypto = require('crypto');

const baseDir = '../SitioPublicoRedisenioSitefinity13';
const date = new Date()
const outputDir = `./output-${date.getDate()}-${date.getUTCMonth() + 1}-${date.getFullYear()}(${date.getHours()}-${date.getMinutes()})`;
const ignoreDirs = ['node_modules', '.git', '.vs', 'packages'];
const ignoreFile = [''];

const scanDir = (dir, fileHashes = {}) => {
    const files = fs.readdirSync(dir);
    files.forEach(file => {
        const filePath = path.join(dir, file);
        const fileStat = fs.statSync(filePath);
        if (fileStat.isDirectory()) {
            if (ignoreDirs.includes(file)) return;
            scanDir(filePath, fileHashes);
        } else {
            if (ignoreFile.includes(file)) return;
            const fileData = fs.readFileSync(filePath);
            const fileHash = crypto.createHash('sha256').update(fileData).digest('hex');
            const relativePath = path.relative(baseDir, filePath);
            if ((fileHash !== fileHashes[relativePath]) && exist) {
                const outputFilePath = path.join(outputDir, relativePath);
                const outputDirPath = path.dirname(outputFilePath);
                if (!fs.existsSync(outputDirPath)) {
                    fs.mkdirSync(outputDirPath, { recursive: true });
                }
                fs.copyFileSync(filePath, outputFilePath);
                //console.log(`Copiado ${filePath} a ${outputFilePath}`);
                accumulator++
            }
            fileHashes[relativePath] = fileHash;
        }
    });
    return fileHashes;
};

const recordFilePath = path.join(__dirname, 'fileHashes.json');
let exist = false
let accumulator = 0

if (fs.existsSync(recordFilePath)) {
    exist = true
    const oldFileHashes = JSON.parse(fs.readFileSync(recordFilePath));
    const newFileHashes = scanDir(baseDir, oldFileHashes);
    fs.writeFileSync(recordFilePath, JSON.stringify(newFileHashes, null, 2));
    console.log('Archivo de registro actualizado')
    console.log(`Se han agregado ${accumulator} archivos`)
} else {
    const fileHashes = scanDir(baseDir);
    fs.writeFileSync(recordFilePath, JSON.stringify(fileHashes, null, 2));
    console.log('Archivo de registro generado')
}

console.timeEnd('Execution time')