package storerMemory

import (
	"errors"
	"strings"

	"github.com/ms-xy/Holmes-Storage-Testdummy/storerGeneric"
	//TODO: Take a look at gocassa, gocqltable, cqlc, cqlr
	//      and check if these packages would be a good addition.
)

type StorerMemory struct {
	submissions map[string][]*storerGeneric.Submission
	objects     map[string]*storerGeneric.Object
	results     map[string][]*storerGeneric.Result
	configs     map[string]*storerGeneric.Config
}

func (s StorerMemory) CreateDB(c []*storerGeneric.DBConnector) error {
	return nil
}

func (s StorerMemory) Initialize(c []*storerGeneric.DBConnector) (storerGeneric.Storer, error) {
	s.submissions = make(map[string][]*storerGeneric.Submission)
	s.objects = make(map[string]*storerGeneric.Object)
	s.results = make(map[string][]*storerGeneric.Result)
	return s, nil
}

func (s StorerMemory) Setup() error {
	return nil
}

func (s StorerMemory) ParseUUID(id string) (string, int, error) {
	parts := strings.Split(id, "---")
	if len(parts) != 2 {
		return "", -1, errors.New("Invalid ID supplied")
	}
	sha256 := parts[0]
	index := parts[1]
	return sha256, index, nil
}

func (s StorerMemory) StoreObject(object *storerGeneric.Object) error {
	submissions, _ := s.GetSubmissionsByObject(object.SHA256)

	l := len(submissions)
	if l == 0 {
		return errors.New("Tried to store an object which was never submitted!")
	}

	source := make([]string, l)
	obj_name := make([]string, l)
	submission_ids := make([]string, l)
	for k, v := range submissions {
		source[k] = v.Source
		obj_name[k] = v.ObjName
		submission_ids[k] = v.Id
	}

	object.Source = source
	object.ObjName = obj_name
	object.Submissions = submission_ids

	s.objects[object.SHA256] = object

	return nil
}

func (s StorerMemory) GetObject(id string) (*storerGeneric.Object, error) {
	object := &storerGeneric.Object{}

	sha256, _, err := s.ParseUUID(id)
	if err != nil {
		return object, err
	}

	if object, ok := s.objects[sha256]; ok {
		return object, nil
	}

	return nil, errors.New("Object not found")
}

func (s StorerMemory) StoreSubmission(submission *storerGeneric.Submission) error {
	if submissions, ok := s.submissions[submission.SHA256]; ok {
		submission.Id = submission.SHA256 + "---" + string(len(submissions))
		s.submissions[submission.SHA256] = append(s.submissions[submission.SHA256], submission)
	} else {
		submission.Id = submission.SHA256 + "---" + string(0)
		s.submissions[submission.SHA256] = []*storerGeneric.Submission{submission}
	}
	return nil
}

func (s StorerMemory) GetSubmission(id string) (*storerGeneric.Submission, error) {
	sha256, index, err := s.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	if submissions, ok := s.submissions[sha256]; ok && len(submissions) >= index-1 {
		return submissions[index], nil
	}
	return nil, errors.New("Submission not found")
}

func (s StorerMemory) GetSubmissionsByObject(sha256 string) ([]*storerGeneric.Submission, error) {
	if submissions, ok := s.submissions[sha256]; ok {
		return submissions, nil
	}
	return []*storerGeneric.Submission{}, nil
}

func (s StorerMemory) StoreResult(result *storerGeneric.Result) error {
	if results, ok := s.results[result.SHA256]; ok {
		result.Id = result.SHA256 + "---" + string(len(results))
		s.results[result.SHA256] = append(s.results[result.SHA256], result)
	} else {
		result.Id = result.SHA256 + "---" + string(0)
		s.results[result.SHA256] = []*storerGeneric.Result{result}
	}
	return nil
}

func (s StorerMemory) GetResult(id string) (*storerGeneric.Result, error) {
	sha256, index, err := s.ParseUUID(id)

	if results, ok := s.results[sha256]; ok && len(results) >= index-1 {
		return results[index], nil
	}
	return nil, errors.New("Result not found")
}

func (s StorerMemory) StoreConfig(config *storerGeneric.Config) error {
	s.configs[config.Path] = config
	return nil
}

func (s StorerMemory) GetConfig(path string) (*storerGeneric.Config, error) {
	if config, ok := s.configs[path]; ok {
		return config, nil
	}
	return nil, errors.New("Config not found")
}
