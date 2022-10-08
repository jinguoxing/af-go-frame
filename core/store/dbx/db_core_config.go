package dbx

import (
    "sync"
    "time"
)

const (
    DefaultGroupName = "default"
)

type Config map[string]ConfigGroup

type ConfigGroup []ConfigNode


type ConfigNode struct {

    Host                 string        `json:"host"`                 // Host of server, ip or domain like: 127.0.0.1, localhost
    Port                 string        `json:"port"`                 // Port, it's commonly 3306.
    User                 string        `json:"user"`                 // Authentication username.
    Pass                 string        `json:"pass"`                 // Authentication password.
    Name                 string        `json:"name"`                 // Default used database name.
    Type                 string        `json:"type"`                 // Database type: mysql, sqlite, mssql, pgsql, oracle.
    Link                 string        `json:"link"`                 // (Optional) Custom link information for all configuration in one single string.
    Extra                string        `json:"extra"`                // (Optional) Extra configuration according the registered third-party database driver.
    Role                 string        `json:"role"`                 // (Optional, "master" in default) Node role, used for master-slave mode: master, slave.
    Debug                bool          `json:"debug"`                // (Optional) Debug mode enables debug information logging and output.
    Prefix               string        `json:"prefix"`               // (Optional) Table prefix.
    DryRun               bool          `json:"dryRun"`               // (Optional) Dry run, which does SELECT but no INSERT/UPDATE/DELETE statements.
    Weight               int           `json:"weight"`               // (Optional) Weight for load balance calculating, it's useless if there's just one node.
    Charset              string        `json:"charset"`              // (Optional, "utf8mb4" in default) Custom charset when operating on database.
    Protocol             string        `json:"protocol"`             // (Optional, "tcp" in default) See net.Dial for more information which networks are available.
    Timezone             string        `json:"timezone"`             // (Optional) Sets the time zone for displaying and interpreting time stamps.
    MaxIdleConnCount     int           `json:"maxIdle"`              // (Optional) Max idle connection configuration for underlying connection pool.
    MaxOpenConnCount     int           `json:"maxOpen"`              // (Optional) Max open connection configuration for underlying connection pool.
    MaxConnLifeTime      time.Duration `json:"maxLifeTime"`          // (Optional) Max amount of time a connection may be idle before being closed.
    QueryTimeout         time.Duration `json:"queryTimeout"`         // (Optional) Max query time for per dql.
    ExecTimeout          time.Duration `json:"execTimeout"`          // (Optional) Max exec time for dml.
    TranTimeout          time.Duration `json:"tranTimeout"`          // (Optional) Max exec time time for a transaction.
    PrepareTimeout       time.Duration `json:"prepareTimeout"`       // (Optional) Max exec time time for prepare operation.
    CreatedAt            string        `json:"createdAt"`            // (Optional) The filed name of table for automatic-filled created datetime.
    UpdatedAt            string        `json:"updatedAt"`            // (Optional) The filed name of table for automatic-filled updated datetime.
    DeletedAt            string        `json:"deletedAt"`            // (Optional) The filed name of table for automatic-filled updated datetime.
    TimeMaintainDisabled bool          `json:"timeMaintainDisabled"` // (Optional) Disable the automatic time maintaining feature.

}


var configs struct{
    sync.RWMutex
    config Config
    group string
}

func init(){
    configs.config = make(Config)
    configs.group = DefaultGroupName
}

func SetConfigGroup(group string,nodes ConfigGroup){
    defer instances.Clear()
    configs.Lock()
    defer configs.Unlock()
    for i ,node := range nodes {
        nodes[i] = parseConfigNode(node)
    }
    configs.config[group] = nodes
}

func SetConfig(config Config){
    defer instances.Clear()
    configs.Lock()
    defer configs.Unlock()
    for k, nodes := range config {
        for i , node := range nodes {
            nodes[i] = parseConfigNode(node)
        }
        config[k] = nodes
    }
    configs.config = config
}

func AddConfigNode(group string,node ConfigNode){
    defer instances.Clear()
    configs.Lock()
    defer configs.Unlock()
    configs.config[group] = append(configs.config[group],parseConfigNode(node))
}


func parseConfigNode(node ConfigNode) ConfigNode {

    //if node.Link !="" && node.Type == "" {
    //    match, _ :=gregex.MatchString(`([a-z]+):(.+)`, node.Link)
    //    if len(match) == 3 {
    //        node.Type = gstr.Trim(match[1])
    //        node.Link = gstr.Trim(match[2])
    //    }
    //}

    return node
}

func AddDefaultConfigNodeGroup(nodes ConfigGroup) {
    SetConfigGroup(DefaultGroupName, nodes)
}

func AddDefaultConfigNode(node ConfigNode) {
    AddConfigNode(DefaultGroupName, node)
}

func GetConfig(group string) ConfigGroup {
    configs.RLock()
    defer configs.RUnlock()
    return configs.config[group]
}

func SetDefaultGroup(name string){

    defer instances.Clear()
    configs.Lock()
    defer configs.Unlock()
    configs.group = name
}


func GetDefaultGroup() string {

    defer instances.Clear()
    configs.RLock()
    defer  configs.RUnlock()
    return configs.group
}






