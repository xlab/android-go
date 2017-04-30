package main

import (
	"io"
	"text/template"
)

type TemplateContext struct {
	ProjectName   string
	ProjectTarget string
	SDKDir        string
}

type Template interface {
	Name() string
	TargetName() string
	Render(ctx *TemplateContext, wr io.Writer) error
}

var templates = []Template{
	&tpl{
		name:   "build.xml.tpl",
		target: "build.xml",
		Template: template.Must(template.New("build.xml.tpl").
			Parse(string(MustAsset("templates/build.xml.tpl")))),
	},
	&tpl{
		name:   "local.properties.tpl",
		target: "local.properties",
		Template: template.Must(template.New("local.properties.tpl").
			Parse(string(MustAsset("templates/local.properties.tpl")))),
	},
	&tpl{
		name:   "project.properties.tpl",
		target: "project.properties",
		Template: template.Must(template.New("project.properties.tpl").
			Parse(string(MustAsset("templates/project.properties.tpl")))),
	},
	&tpl{
		name:   "proguard-project.txt.tpl",
		target: "proguard-project.txt",
		Template: template.Must(template.New("proguard-project.txt.tpl").
			Parse(string(MustAsset("templates/proguard-project.txt.tpl")))),
	},
}

type tpl struct {
	*template.Template

	name   string
	target string
}

func (t *tpl) Name() string {
	return t.name
}

func (t *tpl) TargetName() string {
	return t.target
}

func (t *tpl) Render(ctx *TemplateContext, wr io.Writer) error {
	return t.Template.Execute(wr, ctx)
}
