package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const DEFAULT_CONFIG_FILE = "application.xml"

func LoadConfigFile() *Application {
	// load config file
	xmlFile, err := os.Open(DEFAULT_CONFIG_FILE)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// create object from XML
	var application Application
	xml.Unmarshal(byteValue, &application)

	return &application
}

func getJvmName(application *Application) string {
	var builder strings.Builder

	builder.WriteString("jdk-")
	builder.WriteString(fmt.Sprintf("%d", application.Jvm.Version))

	return builder.String()
}

func FindJvmCommand(application *Application) string {
	var builder strings.Builder

	homeDir := os.Getenv("HOME")
	builder.WriteString(homeDir + "/")
	builder.WriteString("AppBox/JVM/")
	builder.WriteString(getJvmName(application))
	builder.WriteString("/bin/java")

	return filepath.FromSlash(builder.String())
}

func GetCmdLineOptions(application *Application) []string {
	options := make([]string, 0)
	options = append(options, setJvmOptions(application)...)
	options = append(options, setJvmProperties(application)...)
	options = append(options, setModulePath(application)...)
	options = append(options, setModule(application)...)
	options = append(options, setClasspath(application)...)
	options = append(options, setMainClass(application)...)
	options = append(options, setJar(application)...)
	options = append(options, setArguments(application)...)

	return options
}

func setModulePath(application *Application) []string {
	if len(application.Jvm.ModulePath) > 0 {
		return []string{"--module-path", strings.TrimSpace(application.Jvm.ModulePath)}
	}
	return []string{}

}

func setModule(application *Application) []string {
	if len(application.Jvm.Module) > 0 {
		return []string{"--module", strings.TrimSpace(application.Jvm.Module)}
	}
	return []string{}

}

func setClasspath(application *Application) []string {
	if len(application.Jvm.Classpath) > 0 {
		return []string{"--class-path", strings.TrimSpace(application.Jvm.Classpath)}
	}
	return []string{}
}

func setJvmOptions(application *Application) []string {
	return strings.Fields(application.Jvm.JvmOptions)
}

func setJvmProperties(application *Application) []string {
	return strings.Fields(application.Jvm.JvmProperties)
}

func setArguments(application *Application) []string {
	return strings.Fields(application.Jvm.Arguments)
}

func setMainClass(application *Application) []string {
	return []string{strings.TrimSpace(application.Jvm.MainClass)}
}

func setJar(application *Application) []string {
	if len(application.Jvm.Jar) > 0 {
		return []string{"-jar", strings.TrimSpace(application.Jvm.Jar)}
	}
	return []string{}
}