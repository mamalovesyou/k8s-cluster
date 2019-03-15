package generator

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"reflect"
	"strings"
)

const (
	templatesDir      = "../templates"
	clusterTmplFile   = "cluster.tf.tmpl"
	variablesTmplFile = "variables.tf.tmpl"
	hostsTmplFile     = "hosts.ini.tmpl"
)

func templatePath(name string) string {
	cwd, _ := os.Getwd()
	return path.Join(cwd, templatesDir, name)
}

func stringListFormat(list []string) string {
	return fmt.Sprintf(`[ "%s" ]`, strings.Join(list[:], `", "`))
}

func isLast(x int, a interface{}) bool {
	return x == reflect.ValueOf(a).Len()-1
}

func HydrateTerraformCluster(config *TerraformConfig) {

	funcMap := template.FuncMap{
		"incip":      config.IncrementIpCounter,
		"incnode":    config.IncrementNodeCounter,
		"stringlist": stringListFormat,
		"last":       isLast,
	}

	filePath := templatePath(clusterTmplFile)
	t := template.Must(template.New(path.Base(filePath)).Funcs(funcMap).ParseFiles(filePath))

	// Create cluster file
	f, err := os.Create(config.ClusterFilePath())
	if err != nil {
		log.Fatalln("Fail to create terraform cluster file: ", err)
	}

	err = t.Execute(f, config)
	if err != nil {
		log.Fatalln("Execute: ", err)
	}
	f.Close()
	log.Print("Sucessfully created terraform cluster files")
}

func HydrateTerraformVariables(config *TerraformConfig) {

	filePath := templatePath(variablesTmplFile)
	t := template.Must(template.New(path.Base(filePath)).ParseFiles(filePath))

	// Create cluster file
	f, err := os.Create(config.VariablesFilePath())
	if err != nil {
		log.Fatalln("Fail to create terraform variables file: ", err)
	}

	err = t.Execute(f, config)
	if err != nil {
		log.Fatalln("Execute: ", err)
	}
	f.Close()
	log.Print("Sucessfully created terraform variables files")
}

func HydrateHosts(config *HostsConfig) {

	filePath := templatePath(hostsTmplFile)
	t := template.Must(template.New(path.Base(filePath)).ParseFiles(filePath))

	// Create cluster file
	f, err := os.Create(config.HostsFilePath())
	if err != nil {
		log.Fatalln("Fail to create hosts file: ", err)
	}

	err = t.Execute(f, config)
	if err != nil {
		log.Fatalln("Execute: ", err)
	}
	f.Close()
	log.Print("Sucessfully created hosts file")
}
