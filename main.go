package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"bitbucket.org/e_ml_framework/yamada_system/fx_chart_evaluator/database/local"
	"bitbucket.org/e_ml_framework/yamada_system/fx_chart_evaluator/service"
	"github.com/codegangsta/cli"
	"github.com/drone/routes"
)

func main() {
	tApp := cli.NewApp()
	tApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "config, c",
		},
	}

	tApp.Action = func(aContext *cli.Context) {
		var tConfJSONPath string
		if aContext.String("config") != "" {
			tConfJSONPath = aContext.String("config")
		} else {
			tRunningProgramDirectory, tError := filepath.Abs(filepath.Dir(os.Args[0]))
			if tError != nil {
				log.Println("プログラムの走っているディレクトリの取得に失敗")
				os.Exit(1)
			}
			tConfJSONPath = path.Join(tRunningProgramDirectory, "config.json")
		}

		tJSONBytes, tError := ioutil.ReadFile(tConfJSONPath)
		if tError != nil {
			log.Println(tConfJSONPath + "の読み取りに失敗")
			os.Exit(1)
		}

		tConfig := &config{}
		if tError := json.Unmarshal(tJSONBytes, tConfig); tError != nil {
			log.Println(tError)
			log.Println(tConfJSONPath + "が不正なフォーマットです。")
			os.Exit(1)
		}

		tContextPath := "/" + tConfig.ContextPath + "/"
		tService := service.New(&service.Config{
			Database: local.NewDatabase(tConfig.SavePath),
			ContextPath: tContextPath,
			ResourcePath: tConfig.ResourcePath,
		})
		if tError != nil {
			log.Println(tError)
			os.Exit(1)
		}

		tRouter := routes.New()

		tRouter.Get(path.Join(tContextPath, "/rest/v1/users/:id"), tService.GetUser)
		// tRouter.Get(path.Join(tContextPath, "/rest/v1/users/:type/:id"), tService.GetTeacherData)
		tRouter.Post(path.Join(tContextPath, "/rest/v1/users"), tService.SaveUser)
		tRouter.Put(path.Join(tContextPath, "/rest/v1/users"), tService.SaveUser)
		tRouter.Del(path.Join(tContextPath, "/rest/v1/users/:id"), tService.DeleteUser)

		tRouter.Get(path.Join(tContextPath, "/rest/v1/spares/:id"), tService.GetSpareTime)
		tRouter.Get(path.Join(tContextPath, "/rest/v1/spares"), tService.FileterSpareTimes)
		tRouter.Post(path.Join(tContextPath, "/rest/v1/spares"), tService.SaveSpareTime)
		tRouter.Put(path.Join(tContextPath, "/rest/v1/spares"), tService.EditSpareTime)
		tRouter.Delete(path.Join(tContextPath, "/rest/v1/spares/:id"), tService.DeleteSpareTime)
		tRouter.Get(path.Join(tContextPath, "/rest/v1/spares/users/:userId"), tService.GetUserSpareTimes)
		//
		// tRouter.Get(path.Join(tContextPath, "/rest/v1/spares"), tService.GetAllProgressInfos)
		// tRouter.Get(path.Join(tContextPath, "/rest/v1/badregions/:type"),tService.GetBadRegions)
		// tRouter.Put(path.Join(tContextPath, "/rest/v1/badregions"),tService.PutBadRegion)
		// tRouter.Del(path.Join(tContextPath, "/rest/v1/badregions/:type/:id"),tService.DeleteBadRegion)

		tRouter.Get(path.Join(tContextPath, "/.*"), tService.GetFile)

		http.Handle(tContextPath, tRouter)

		http.ListenAndServe(":"+strconv.Itoa(tConfig.Port), nil)
	}

	tApp.Run(os.Args)

}

type config struct {
	ContextPath  string `json:"context_path"`
	Port         int    `json:"port"`
	SavePath     string `json:"save_path"`github.com/keita0q/himatch
	ResourcePath string `json:"resource_path"`
}
