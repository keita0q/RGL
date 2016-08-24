package service

import (
	"net/http"
	"net/url"

	"encoding/json"

	"io/ioutil"

	"log"

	"strconv"

	"fmt"
	"github.com/keita0q/himatch/core"
	"github.com/keita0q/himatch/database"
	"github.com/keita0q/himatch/model"
	"path/filepath"
	"strings"
	"time"
)

type Service struct {
	himatch      core.Himatch
	contextPath  string
	resourcePath string
}

type Config struct {
	Database     database.Database
	ContextPath  string
	ResourcePath string
}

func New(aConfig *Config) *Service {
	return &Service{
		himatch:      core.New(aConfig.Database),
		contextPath:  aConfig.ContextPath,
		resourcePath: aConfig.ResourcePath,
	}
}

func (aService *Service) GetFile(aWriter http.ResponseWriter, aRequest *http.Request) {
	tPath := strings.TrimPrefix(aRequest.RequestURI, aService.contextPath)
	fmt.Println(tPath)
	if i := strings.Index(tPath, "?"); i > 0 {
		tPath = tPath[:i]
	}
	http.ServeFile(aWriter, aRequest, filepath.Join(aService.resourcePath, tPath))
}

type user struct {
	ID   string     `json:"id"`
	Name string     `json:"name"`
	Age  int        `json:"age"`
	Sns  *model.Sns `json:"sns"`
	Tags []string   `json:"tags"`
}

func (aService *Service) GetUser(aWriter http.ResponseWriter, aRequest *http.Request) {
	aService.handler(func(aQueries url.Values, aRequestBody []byte) (int, interface{}, error) {
		tID := aQueries.Get(":id")

		tUser, tError := aService.himatch.GetUser(tID)
		if tError != nil {
			return handleError(tError), nil, tError
		}

		tReturnUser := &user{
			ID:   tUser.ID,
			Name: tUser.Name,
			Age:  tUser.Age,
			Sns:  tUser.Sns,
		}
		return http.StatusOK, tReturnUser, nil
	})(aWriter, aRequest)
}

func (aService *Service) SaveUser(aWriter http.ResponseWriter, aRequest *http.Request) {
	aService.handler(func(aQuerys url.Values, aRequestBody []byte) (int, interface{}, error) {
		tUser := &model.User{}

		if tError := json.Unmarshal(aRequestBody, tUser); tError != nil {
			return http.StatusBadRequest, nil, tError
		}

		if tError := aService.himatch.SaveUser(tUser); tError != nil {
			return http.StatusInternalServerError, nil, tError
		}
		return http.StatusNoContent, nil, nil
	})(aWriter, aRequest)
}

func (aService *Service) EditUser(aWriter http.ResponseWriter, aRequest *http.Request) {
	aService.handler(func(aQuerys url.Values, aRequestBody []byte) (int, interface{}, error) {
		tUser := &user{}

		if tError := json.Unmarshal(aRequestBody, tUser); tError != nil {
			return http.StatusBadRequest, nil, tError
		}

		tOld, tError := aService.himatch.GetUser(tUser.ID)
		if tError != nil {
			return http.StatusBadRequest, nil, tError
		}

		if tError := aService.himatch.SaveUser(userUpdate(tOld, tUser)); tError != nil {
			return http.StatusInternalServerError, nil, tError
		}
		return http.StatusNoContent, nil, nil
	})(aWriter, aRequest)
}

func userUpdate(tOld *model.User, tUser *user) *model.User {
	return &model.User{ID: tUser.ID, Password: tOld.Password, Name: tUser.Name, Age: tUser.Age, Sns: tUser.Sns, Tags: tUser.Tags}
}

func (aService *Service) DeleteUser(aWriter http.ResponseWriter, aRequest *http.Request) {
	aService.handler(func(aQueries url.Values, aRequestBody []byte) (int, interface{}, error) {
		tID := aQueries.Get(":id")
		if tError := aService.himatch.GetUser(tID); tError != nil {
			return http.StatusInternalServerError, nil, tError
		}
		return http.StatusNoContent, nil, nil
	})(aWriter, aRequest)
}

func (aService *Service) GetSpareTime(aWriter http.ResponseWriter, aRequest *http.Request) {
	aService.handler(func(aQueries url.Values, aRequestBody []byte) (int, interface{}, error) {
		tID := aQueries.Get(":id")
		tSpareTime, tError := aService.himatch.GetSpareTime(tID)
		if tError != nil {
			return handleError(tError), nil, tError
		}

		return http.StatusOK, tSpareTime, nil
	})(aWriter, aRequest)
}

func (aService *Service) GetUserSpareTime(aWriter http.ResponseWriter, aRequest *http.Request) {
	aService.handler(func(aQueries url.Values, aRequestBody []byte) (int, interface{}, error) {
		tID := aQueries.Get(":userId")
		tSpareTimes, tError := aService.himatch.FilterSpareTimesByUserID(tID)
		if tError != nil {
			return handleError(tError), nil, tError
		}

		return http.StatusOK, tSpareTimes, nil
	})(aWriter, aRequest)
}

func (aService *Service) FilterSpareTimes(aWriter http.ResponseWriter, aRequest *http.Request) {
	aService.handler(func(aQueries url.Values, aRequestBody []byte) (int, interface{}, error) {

		type request struct {
			Tags []string  `json:"tags,omitempty"`
			Time time.Time `json:"time, omitempty"`
		}
		tRequest := &request{}

		if tError := json.Unmarshal(aRequestBody, tRequest); tError != nil {
			return http.StatusBadRequest, nil, tError
		}

		if tRequest.Tags == nil && tRequest.Time == nil {
			return http.StatusBadRequest, nil, nil
		}

		tSpareTimes := []model.SpareTime{}
		if tRequest.Tags != nil && tRequest.Time != nil {
			tSpareTimes, tError := aService.himatch.FilterSpareTimesByTagsAndTime(tRequest.Time, tRequest.Tags)
			if tError != nil {
				return http.StatusInternalServerError, nil, tError
			}
			return http.StatusOK, tSpareTimes, nil
		}

		if tRequest.Tags != nil {
			tSpareTimes, tError := aService.himatch.FilterSpareTimesByTags(tRequest.Tags)
			if tError != nil {
				return http.StatusInternalServerError, nil, tError
			}
			return http.StatusOK, tSpareTimes, nil
		}

		tSpareTimes, tError := aService.himatch.FilterSpareTimesByTime(tRequest.Time)
		if tError != nil {
			return http.StatusInternalServerError, nil, tError
		}

		return http.StatusOK, tSpareTimes, nil
	})(aWriter, aRequest)
}

func (aService *Service) SaveSpareTime(aWriter http.ResponseWriter, aRequest *http.Request) {
	aService.handler(func(aQuerys url.Values, aRequestBody []byte) (int, interface{}, error) {
		tSpareTime := &model.SpareTime{}
		if tError := json.Unmarshal(aRequestBody, tSpareTime); tError != nil {
			return http.StatusBadRequest, nil, tError
		}

		if tError := aService.himatch.SaveSpareTime(tSpareTime); tError != nil {
			return http.StatusInternalServerError, nil, tError
		}
		return http.StatusNoContent, nil, nil
	})(aWriter, aRequest)
}

func (aService *Service) DeleteSpareTime(aWriter http.ResponseWriter, aRequest *http.Request) {
	aService.handler(func(aQueries url.Values, aRequestBody []byte) (int, interface{}, error) {
		tID := aQueries.Get(":id")
		if tError := aService.himatch.DeleteSpareTime(tID); tError != nil {
			return http.StatusInternalServerError, nil, tError
		}
		return http.StatusNoContent, nil, nil
	})(aWriter, aRequest)
}

func (aService *Service) GetUserSpareTimes(aWriter http.ResponseWriter, aRequest *http.Request) {
	aService.handler(func(aQueries url.Values, aRequestBody []byte) (int, interface{}, error) {
		tID := aQueries.Get(":userId")
		tSpareTimes, tError := aService.himatch.FilterSpareTimesByUserID(tID)
		if tError != nil {
			return handleError(tError), nil, tError
		}

		return http.StatusOK, tSpareTimes, nil
	})(aWriter, aRequest)
}

func handleError(aError error) int {
	if _, ok := aError.(*database.NotFoundError); ok {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}

func (aService *Service) handler(aAPI func(url.Values, []byte) (int, interface{}, error)) func(http.ResponseWriter, *http.Request) {
	return func(aWriter http.ResponseWriter, aRequest *http.Request) {
		log.Printf("[INFO] access:%s", aRequest.RequestURI)
		defer aRequest.Body.Close()

		tResponseBody, tError := ioutil.ReadAll(aRequest.Body)
		if tError != nil {
			http.Error(aWriter, tError.Error(), http.StatusBadRequest)
		}
		tStatusCode, tResult, tError := aAPI(aRequest.URL.Query(), tResponseBody)
		if tError != nil {
			http.Error(aWriter, tError.Error(), tStatusCode)
			return
		}

		if tStatusCode == http.StatusNoContent {
			aWriter.WriteHeader(http.StatusNoContent)
			return
		}

		tBytes, tError := json.MarshalIndent(tResult, "", "  ")
		if tError != nil {
			http.Error(aWriter, tError.Error(), tStatusCode)
			return
		}

		aWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		aWriter.Header().Set("Content-Length", strconv.Itoa(len(tBytes)))
		aWriter.Header().Set("Access-Control-Allow-Origin", "*")
		aWriter.WriteHeader(tStatusCode)
		aWriter.Write(tBytes)
	}
}
