package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

// The structure of the Terraform .tfstate file
type TerraformState struct {
	Version int
	Serial  int
	Backend *TerraformBackend
	Modules []TerraformStateModule
}

// The structure of the "backend" section of the Terraform .tfstate file
type TerraformBackend struct {
	Type   string
	Config map[string]interface{}
}

// The structure of a "module" section of the Terraform .tfstate file
type TerraformStateModule struct {
	Path      []string
	Outputs   map[string]interface{}
	Resources map[string]ScalewayResource
}

type ScalewayResource struct {
	Type      string
	DependsOn []string
	Primary   ScalewayPrimary
	Provider  string
}

type ScalewayPrimary struct {
	ID         string
	Attributes map[string]interface{}
	Meta       interface{}
	Tainted    bool
}

type ScalewayServer struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	PrivateIp string   `json:"private_ip"`
	PublicIp  string   `json:"public_ip"`
	Tags      []string `json:"tags"`
}

func NewScalewayServer(data map[string]interface{}) *ScalewayServer {
	a := &ScalewayServer{}
	l := data["tags.#"]
	fmt.Printf("%v %T", l, l)
	tagsLen, _ := strconv.Atoi(l.(string))
	tags := make([]string, tagsLen)
	for t := 0; t < tagsLen; t++ {
		name := fmt.Sprintf("tags.%d", t)
		tags[t] = data[name].(string)
	}
	a.ID = data["id"].(string)
	a.Name = data["name"].(string)
	a.PublicIp = data["public_ip"].(string)
	a.PrivateIp = data["private_ip"].(string)
	a.Tags = tags
	return a
}

func ParseTerraformStateFile(filePath string) (*TerraformState, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshall the data
	terraformState := &TerraformState{}
	if err := json.Unmarshal(bytes, terraformState); err != nil {
		return nil, err
	}

	return terraformState, nil
}

func ServerFromTerraformState(terraformState *TerraformState) []ScalewayServer {
	server := []ScalewayServer{}
	for _, m := range terraformState.Modules {
		for _, resource := range m.Resources {
			if resource.Type == "scaleway_server" {
				spew.Dump(resource.Primary.Attributes)
				sv := NewScalewayServer(resource.Primary.Attributes)
				server = append(server, *sv)
			}
		}
	}
	return server
}

func ServerListToHostsFile(servers []ScalewayServer) map[string][]string {
	hosts := map[string][]string{}
	for _, h := range servers {
		for _, t := range h.Tags {
			if hList, ok := hosts[t]; ok {
				hosts[t] = append(hList, h.Name)
			} else {
				hosts[t] = []string{h.Name}
			}
		}
	}
	return hosts
}
