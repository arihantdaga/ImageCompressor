package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main()  {
	fmt.Println("Hello")
	loadEnv()
	processFiles(goDotEnvVariable("INPUT_PATH"), goDotEnvVariable("OUTPUT_PATH"))
}

func processFiles(input_dir string, output_dir string)  {
		start:= time.Now().Unix()
		files, err := ioutil.ReadDir(input_dir)
		count:= 0
		if err != nil {
		 log.Print(err)
		}
		for _,f := range files{
			count+=1
			fmt.Println(f.Name(), count)
			filename:= f.Name()
			extension:= strings.ToLower(filepath.Ext(filename))
			if extension == ".jpg" || extension == ".jpeg" || extension == ".png"{
				// Do something
				input_path:= filepath.Join(input_dir, f.Name())
				output_path := filepath.Join(output_dir, f.Name())
				processImage(input_path, output_path)
			}else{
				continue
			}
		}
		fmt.Println("Time taken : ", (time.Now().Unix() - start))
}

func processImage(input_path string, output_path string){
	options:= []string{"convert", "-resize", "2000x2000", "-quality", "80", input_path, output_path}
	cmd := exec.Command("magick", options...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}