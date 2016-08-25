package bolt

import (
	"github.com/boltdb/bolt"
	"github.com/keita0q/himatch/model"
	"encoding/json"
	"github.com/keita0q/himatch/database"
)

type BoltDB struct {
	db                  *bolt.DB
	userBucketName      []byte
	spareTimeBucketName []byte
}

func NewDatabase(aSavePath string) (*BoltDB, error) {
	tDB, tError := bolt.Open(aSavePath, 0600, nil)
	if tError != nil {
		return nil, tError
	}
	tBolt := &BoltDB{db: tDB, userBucketName:[]byte("users"), spareTimeBucketName:[]byte("spares")}
	tDB.Update(func(tx *bolt.Tx) error {
		_, tError := tx.CreateBucketIfNotExists(tBolt.userBucketName)
		if tError != nil {
			return tError
		}
		_, tError = tx.CreateBucketIfNotExists(tBolt.spareTimeBucketName)

		return tError
	})
	return tBolt, nil
}

func (aBolt *BoltDB) LoadUser(aID string) (*model.User, error) {
	tUser := &model.User{}
	tError := aBolt.db.View(func(tx *bolt.Tx) error {
		tBucket := tx.Bucket(aBolt.userBucketName)

		return findBucket(tBucket, aID, func(aValue []byte) error {
			return json.Unmarshal(aValue, tUser)
		})
	})
	return tUser, tError
}

func (aBolt *BoltDB)SaveUser(aUser *model.User) error {
	return aBolt.db.Update(func(tx *bolt.Tx) error {
		tBucket := tx.Bucket(aBolt.userBucketName)
		tValue, err := json.Marshal(aUser)
		if err != nil {
			return err
		}

		return tBucket.Put([]byte(aUser.ID), tValue)
	})
}

func (aBolt *BoltDB)DeleteUser(aID string) error {
	return aBolt.db.Update(func(tx *bolt.Tx) error {
		tBucket := tx.Bucket(aBolt.userBucketName)
		if tError := findBucket(tBucket, aID, func([]byte) error {
			return nil
		}); tError != nil {
			return tError
		}
		return tBucket.Delete([]byte(aID))
	})
}

func (aBolt *BoltDB)LoadAllSpareTimes() ([]model.SpareTime, error) {
	tSpareTimes := []model.SpareTime{}
	tError := aBolt.db.View(func(tx *bolt.Tx) error {
		tBucket := tx.Bucket(aBolt.userBucketName)
		return tBucket.ForEach(func(_, aValue []byte) error {
			tSpareTime := &model.SpareTime{}

			if tError := json.Unmarshal(aValue, tSpareTime); tError != nil {
				return tError
			}
			tSpareTimes = append(tSpareTimes, *tSpareTime)
			return nil
		})
	})

	return tSpareTimes, tError
}

func (aBolt *BoltDB)LoadSpareTime(aID string) (*model.SpareTime, error) {
	tSpareTime := &model.SpareTime{}
	tError := aBolt.db.View(func(tx *bolt.Tx) error {
		tBucket := tx.Bucket(aBolt.spareTimeBucketName)

		return findBucket(tBucket, aID, func(aValue []byte) error {
			return json.Unmarshal(aValue, tSpareTime)
		})
	})
	return tSpareTime, tError
}

func (aBolt *BoltDB)SaveSpareTime(aSpareTime *model.SpareTime) error {
	return aBolt.db.Update(func(tx *bolt.Tx) error {
		tBucket := tx.Bucket(aBolt.spareTimeBucketName)
		tValue, err := json.Marshal(aSpareTime)
		if err != nil {
			return err
		}

		return tBucket.Put([]byte(aSpareTime.ID), tValue)
	})
}

func (aBolt *BoltDB)DeleteSpareTime(aID string) error {
	return aBolt.db.Update(func(tx *bolt.Tx) error {
		tBucket := tx.Bucket(aBolt.spareTimeBucketName)
		if tError := findBucket(tBucket, aID, func([]byte) error {
			return nil
		}); tError != nil {
			return tError
		}
		return tBucket.Delete([]byte(aID))
	})
}

func findBucket(aParent *bolt.Bucket, aKey string, aFunc func([]byte) error) error {
	tExist := false
	tError := aParent.ForEach(func(k, v []byte) error {
		if string(k) == aKey {
			tExist = true
			return aFunc(v)
		}
		return nil
	})
	if !tExist {
		return database.NewNotFoundError(aKey + "が見つかりません。")
	}
	return tError
}

func (aBolt *BoltDB)Close() error {
	return aBolt.db.Close()
}