package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

var baseDir = "../"
var outputDir = fmt.Sprintf("./output-%d-%d-%d_%d-%d", time.Now().Day(), time.Now().Month(), time.Now().Year(), time.Now().Hour(), time.Now().Minute())
var ignoreDirs = map[string]bool{"node_modules": true, "file_inspector_go": true}
var ignoreFiles = []*regexp.Regexp{regexp.MustCompile(`.+\.js`), regexp.MustCompile(`.+\.ts`)}
var accumulator int

func main() {
	recordFilePath := filepath.Join(".", "fileHashes.json")
	exist := false

	if _, err := os.Stat(recordFilePath); err == nil {
		exist = true
	}

	var fileHashes map[string]string
	if exist {
		oldFileHashes, err := ioutil.ReadFile(recordFilePath)
		if err != nil {
			fmt.Println("Error al leer el archivo de registro:", err)
			return
		}
		fileHashes = scanDir(baseDir, parseFileHashes(oldFileHashes))
	} else {
		fileHashes = scanDir(baseDir, make(map[string]string))
	}

	if err := writeToFile(recordFilePath, fileHashes); err != nil {
		fmt.Println("Error al escribir en el archivo de registro:", err)
		return
	}

	if accumulator > 0 {
		fmt.Printf("Se han agregado %d archivos\n", accumulator)
	}

	if exist {
		fmt.Println("Archivo de registro actualizado")
	} else {
		fmt.Println("Archivo de registro generado")
	}
}

func scanDir(dir string, fileHashes map[string]string) map[string]string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error al leer el directorio %s: %v\n", dir, err)
		return fileHashes
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			if _, ignore := ignoreDirs[file.Name()]; !ignore {
				fileHashes = scanDir(filePath, fileHashes)
			}
		} else {
			if !shouldIgnore(file.Name()) {
				fileData, err := ioutil.ReadFile(filePath)
				if err != nil {
					fmt.Printf("Error al leer el archivo %s: %v\n", filePath, err)
					continue
				}

				fileHash := hashFile(fileData)
				relativePath, err := filepath.Rel(baseDir, filePath)
				if err != nil {
					fmt.Printf("Error al obtener la ruta relativa de %s: %v\n", filePath, err)
					continue
				}

				if fileHash != fileHashes[relativePath] {
					outputFilePath := filepath.Join(outputDir, relativePath)
					outputDirPath := filepath.Dir(outputFilePath)
					if err := os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
						fmt.Printf("Error al crear el directorio %s: %v\n", outputDirPath, err)
						continue
					}
					if err := ioutil.WriteFile(outputFilePath, fileData, os.ModePerm); err != nil {
						fmt.Printf("Error al escribir en el archivo %s: %v\n", outputFilePath, err)
						continue
					}
					accumulator++
				}
				fileHashes[relativePath] = fileHash
			}
		}
	}

	return fileHashes
}

func shouldIgnore(fileName string) bool {
	for _, regex := range ignoreFiles {
		if regex.MatchString(fileName) {
			return true
		}
	}
	return false
}

func hashFile(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

func parseFileHashes(data []byte) map[string]string {
	var fileHashes map[string]string
	if err := json.Unmarshal(data, &fileHashes); err != nil {
		fmt.Println("Error al analizar los datos del archivo de registro:", err)
	}
	return fileHashes
}

func writeToFile(filePath string, data interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}
