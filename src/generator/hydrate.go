package generator

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	templatesDir             = "templates"
	clusterTmplFile          = "cluster.tf.tmpl"
	variablesTmplFile        = "variables.tf.tmpl"
	hostsTmplFile            = "hosts.ini.tmpl"
	initialTmplFile          = "initial.yml.tmpl"
	kubeDependenciesTmplFile = "kube-dependencies.yml.tmpl"
	kubeClusterTmplFile      = "kube-cluster.yml.tmpl"
)

type Hydrator struct {
	TemplatesDir string
	Templates    map[string]*template.Template
}

func (h *Hydrator) LoadTemplates() {

	if h.Templates == nil {
		h.Templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(h.TemplatesDir + "*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	includeFiles, err := filepath.Glob(h.TemplatesDir + "*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}
	log.Println("Templates successfully loaded")
}

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
	t := template.Must(templates.Funcs(funcMap).ParseFiles(filePath))

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

func HydratePlaybooks(config *PlaybooksConfig) {
	playbooks := map[string]string{
		initialTmplFile:          "initial.yml",
		kubeClusterTmplFile:      "kube-cluster.yml",
		kubeDependenciesTmplFile: "kube-dependencies.yml",
	}

	for tmpl, fileName := range playbooks {

		filePath := templatePath(tmpl)
		t := template.Must(template.New(path.Base(filePath)).ParseFiles(filePath))

		f, err := os.Create(path.Join(config.DeploymentsPath, fileName))
		if err != nil {
			log.Fatalln("Fail to create playbook: ", err)
		}

		err = t.Execute(f, config)
		if err != nil {
			log.Fatalln("Execute: ", err)
		}
		f.Close()
		log.Printf("Sucessfully created playbook %s", fileName)
	}
}
