package main

import (
	"encoding/json"

	"github.com/glog"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	ldb, err := NewlevelDB("./goleveldb/db/dbfile")
	if err != nil {
		return
	}
	defer ldb.db.Close()
	//test some sample
	tom := NewStudent("tom", 10)
	mary := NewStudent("mary", 15)

	//save tom and mary info
	tombytes, _ := json.Marshal(tom)
	marybytes, _ := json.Marshal(mary)

	if err := ldb.Put([]byte("tom"), tombytes); err != nil {
		glog.Error(err)
		return
	}
	if err := ldb.Put([]byte("mary"), marybytes); err != nil {
		glog.Error(err)
		return
	}

	//get tom
	gettom, err := ldb.Get([]byte("tom"))
	if err != nil {
		glog.Error(err)
		return
	}
	glog.Info(string(gettom))
	var backTom Student
	err = json.Unmarshal(gettom, &backTom)
	glog.Infof("Get tom: %v\n", backTom)
	glog.Info(backTom)
}

//Student a student
type Student struct {
	Name string
	Age  uint
}

//NewStudent new a student object
func NewStudent(name string, age uint) *Student {
	return &Student{Name: name, Age: age}
}

//Lldb leveldb
type Lldb struct {
	db *leveldb.DB
}

//NewlevelDB new a level db
func NewlevelDB(dbpath string) (*Lldb, error) {
	db, err := leveldb.OpenFile(dbpath, nil)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return &Lldb{db: db}, err
}

//Put put method of Lldb
func (ldb *Lldb) Put(key, value []byte) error {
	return ldb.db.Put(key, value, nil)
}

//Get get method of lldb
func (ldb *Lldb) Get(key []byte) (value []byte, err error) {
	data, err := ldb.db.Get(key, nil)
	if err != nil {
		glog.Error(err)
		return nil, err
	}
	return data, nil
}

//Update update leveldb
func (ldb *Lldb) Update(key, value []byte) error {
	return nil
}
