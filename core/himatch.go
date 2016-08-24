package core

import (
	"github.com/keita0q/himatch/database"
	"github.com/keita0q/himatch/model"
	"time"
)

type Himatch struct {
	database database.Database
}

func New(aDatabase database.Database) *Himatch {
	return &Himatch{database: aDatabase}
}

func (aHimatch *Himatch) GetUser(aID string) (*model.User, error) {
	return aHimatch.database.LoadUser(aID)
}

func (aHimatch *Himatch) SaveUser(aUser *model.User) error {
	return aHimatch.database.SaveUser(aUser)
}

func (aHimatch *Himatch) DeleteUser(aID string) error {
	return aHimatch.database.DeleteUser(aID)
}

func (aHimatch *Himatch) GetSpareTime(aID string) (*model.SpareTime, error) {
	return aHimatch.database.LoadSpareTime(aID)
}

func (aHimatch *Himatch) SaveSpareTime(aSpareTime *model.SpareTime) error {
	return aHimatch.database.SaveSpareTime(aSpareTime)
}

func (aHimatch *Himatch) DeleteSpareTime(aID string) error {
	return aHimatch.database.DeleteSpareTime(aID)
}

func (aHimatch *Himatch) FilterSpareTimesByTime(aTime time.Time) ([]model.SpareTime, error) {
	return aHimatch.filterSpareTime(func(aSpareTime *model.SpareTime) bool {
		return aSpareTime.Start.Equal(aTime) || aSpareTime.Start.Before(aTime) && aSpareTime.End.After(aTime)
	})
}

func (aHimatch *Himatch) FilterSpareTimesByTags(aTags []string) ([]model.SpareTime, error) {
	return aHimatch.filterSpareTime(func(aSpareTime *model.SpareTime) bool {
		tIsContains := true

		for _, tTag := range aTags {
			for _, t := range aSpareTime.Tags {
				tIsContains = false
				if t == tTag {
					tIsContains = true
					break
				}
			}
			if !tIsContains {
				return false
			}
		}
		return true
	})
}

func (aHimatch *Himatch) FilterSpareTimesByUserID(aUserID string) ([]model.SpareTime, error) {
	return aHimatch.filterSpareTime(func(aSpareTime *model.SpareTime) bool {
		return aSpareTime.UserID == aUserID
	})
}

func (aHimatch *Himatch) FilterSpareTimesByTagsAndTime(aTime time.Time, aTags []string) ([]model.SpareTime, error) {
	return aHimatch.filterSpareTime(func(aSpareTime *model.SpareTime) bool {
		if !(aSpareTime.Start.Equal(aTime) || aSpareTime.Start.Before(aTime) && aSpareTime.End.After(aTime)) {
			return false
		}
		tIsContains := true

		for _, tTag := range aTags {
			for _, t := range aSpareTime.Tags {
				tIsContains = false
				if t == tTag {
					tIsContains = true
					break
				}
			}
			if !tIsContains {
				return false
			}
		}
		return true
	})
}

func (aHimatch *Himatch) filterSpareTime(aFilter func(aSpareTime *model.SpareTime) bool) ([]model.SpareTime, error) {
	tSpareTimes, tError := aHimatch.database.LoadAllSpareTimes()
	if tError != nil {
		return nil, tError
	}

	tNewSpareTimes := []model.SpareTime{}
	for _, tSpareTime := range tSpareTimes {
		if aFilter(&tSpareTime) {
			tSpareTimes = append(tNewSpareTimes, tSpareTime)
		}
	}
	return tNewSpareTimes, nil
}
