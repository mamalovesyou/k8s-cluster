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

const clusterTmplFile = "./templates/cluster.tf.tmpl"
const variablesFile = "./template/variables.tf.tmpl"

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

	t := template.Must(template.New(path.Base(clusterTmplFile)).Funcs(funcMap).ParseFiles(clusterTmplFile))

	// Create cluster file
	f, err := os.Create(config.ClusterFilePath())
	if err != nil {
		log.Fatalln("Fail to create terraform file: ", err)
	}

	err = t.Execute(f, config)
	if err != nil {
		log.Fatalln("Execute: ", err)
	}
	f.Close()
	log.Print("Sucessfully created terraform files")
}
