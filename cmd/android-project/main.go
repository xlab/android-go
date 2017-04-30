package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jawher/mow.cli"
	"github.com/xlab/closer"
	"github.com/xlab/treeprint"
)

var app = cli.App("android-project", `Android project bootstrap tool for use instead of version that was removed from Android SDK (since Revision 25.3.0).
	See https://developer.android.com/studio/releases/sdk-tools.html`)

func init() {
	log.SetFlags(0)
}

func main() {
	app.Command("update", "Updates an Android project (must already have an AndroidManifest.xml).", cmdUpdateFunc)
	app.Command("platforms", "Lists Android platforms supported within SDK.", cmdPlatformsFunc)

	if err := app.Run(os.Args); err != nil {
		closer.Fatalln(err)
	}
}

func cmdPlatformsFunc(c *cli.Cmd) {
	sdkDir := c.String(cli.StringOpt{
		Name:      "sdk",
		Desc:      "Android SDK location, can also be set using ANDROID_HOME env variable.",
		EnvVar:    "ANDROID_HOME",
		HideValue: true,
	})
	c.Spec = "[--sdk]"
	c.Action = func() {
		defer closer.Close()

		if len(*sdkDir) == 0 {
			closer.Fatalln("ANDROID_HOME environment variable not set (or --sdk flag not provided).")
		}
		_, platformsDir := learnPaths(*sdkDir)
		log.Println("Available Android platforms:")
		platforms := learnPlatforms(platformsDir)
		for _, platform := range platforms {
			log.Println("*", platform)
		}
	}
}

func cmdUpdateFunc(c *cli.Cmd) {
	pathOpt := c.StringOpt("p path", "", "The project's directory.")
	nameOpt := c.StringOpt("n name", "", "Project name.")
	forceOpt := c.BoolOpt("f force", false, "Force generation, even if the target platform is not supported.")
	targetOpt := c.StringOpt("t target", "android-23", "Target ID to set for the project.")
	sdkDir := c.String(cli.StringOpt{
		Name:      "sdk",
		Desc:      "Android SDK location, can also be set using ANDROID_HOME env variable.",
		EnvVar:    "ANDROID_HOME",
		HideValue: true,
	})
	c.Spec = "[--sdk] [--target] --name --path"
	c.Action = func() {
		defer closer.Close()

		if len(*sdkDir) == 0 {
			closer.Fatalln("ANDROID_HOME environment variable not set (or --sdk flag not provided).")
		}
		tree := treeprint.New()
		closer.Bind(func() {
			log.Println(tree)
		})
		ndkDir, platformsDir := learnPaths(*sdkDir)

		env := tree.AddBranch("Environment").
			AddMetaNode(*sdkDir, "Android SDK location").
			AddMetaNode(ndkDir, "Android NDK location").
			AddMetaNode(platformsDir, "Android Platforms location")
		platforms := learnPlatforms(platformsDir)
		env.AddMetaNode(platforms, "Android Platforms available")

		ctx := &TemplateContext{
			ProjectName:   *nameOpt,
			ProjectTarget: *targetOpt,

			SDKDir: *sdkDir,
		}
		tree.AddBranch("Project").
			AddMetaNode(ctx.ProjectName, "Project name").
			AddMetaNode(ctx.ProjectTarget, "Project target").
			AddMetaNode(*pathOpt, "Project location")

		manifestFileName := filepath.Join(*pathOpt, "AndroidManifest.xml")
		if _, err := os.Stat(manifestFileName); os.IsNotExist(err) {
			closer.Fatalln("AndroidManifest.xml not found in project location")
		}
		var platformSupported bool
		for _, platform := range platforms {
			if *targetOpt == platform {
				platformSupported = true
			}
		}
		if !platformSupported && !(*forceOpt) {
			closer.Fatalln("Platform", *targetOpt, "not supported within this SDK, build may fail. Use -f flag to override.")
		}

		files := tree.AddBranch("Files updated")
		for _, tpl := range templates {
			fileName := filepath.Join(*pathOpt, tpl.TargetName())
			f, err := os.Create(fileName)
			if err != nil {
				closer.Fatalf("Failed to create file %s: %v", fileName, err)
			}
			if err := tpl.Render(ctx, f); err != nil {
				closer.Fatalf("Failed to render template %s: %v", tpl.Name(), err)
			}
			f.Close()
			files.AddNode(fileName)
		}
	}
}

func learnPlatforms(platformsDir string) []string {
	var platforms []string
	filepath.Walk(platformsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == platformsDir {
			return nil
		}
		platforms = append(platforms, filepath.Base(path))
		return filepath.SkipDir
	})
	return platforms
}

func learnPaths(sdkDir string) (ndkDir, platformsDir string) {
	if _, err := os.Stat(sdkDir); err != nil {
		closer.Fatalln("Android SDK location is provided but not valid:", sdkDir)
	}
	ndkDir = filepath.Join(sdkDir, "ndk-bundle")
	if _, err := os.Stat(sdkDir); err != nil {
		log.Println("WARN: Android SDK location is provided but has no ndk-bundle:", ndkDir)
	}
	platformsDir = filepath.Join(sdkDir, "platforms")
	if _, err := os.Stat(sdkDir); err != nil {
		closer.Fatalln("Android SDK location is provided but has no platforms:", platformsDir)
	}
	return
}
