package main

import (
	"github.com/frame-go/framego"

	"{{ cookiecutter.go_module }}/api/{{ cookiecutter.service_package_name }}"
	"{{ cookiecutter.go_module }}/internal/handlers"
)

func main() {
	app := framego.NewApp(&framego.AppInfo{
		Name:        "{{ cookiecutter.app_package_name }}",
		Version:     "0.1.0",
		Description: "{{ cookiecutter.project_title }}",
	})
	app.Init()
	{{ cookiecutter.service_package_name }}.RegisterSwag()
	{{ cookiecutter.service_package_name }}.Register{{ cookiecutter.__service_name_title }}WithGrpcAndHttp(app.GetService("{{ cookiecutter.service_name }}"), handlers.NewServer(app))
	app.RunOrExit()
}
