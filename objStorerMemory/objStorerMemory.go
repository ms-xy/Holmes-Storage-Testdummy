package objStorerMemory

import (
	"errors"
	"github.com/ms-xy/Holmes-Storage-Testdummy/objStorerGeneric"
)

type ObjStorerMemory struct {
	samples map[string]*objStorerGeneric.Sample
}

func (s ObjStorerMemory) Initialize(configs []*objStorerGeneric.ObjDBConnector) (objStorerGeneric.ObjStorer, error) {
	s.samples = make(map[string]*objStorerGeneric.Sample)
	return s, nil
}

func (s ObjStorerMemory) Setup() error {
	return nil
}

func (s ObjStorerMemory) StoreSample(sample *objStorerGeneric.Sample) error {
	if _, exists := s.samples[sample.SHA256]; !exists {
		s.samples[sample.SHA256] = sample
	}
	return nil
}

func (s ObjStorerMemory) GetSample(sha256 string) (*objStorerGeneric.Sample, error) {
	if sample, exists := s.samples[sha256]; exists {
		return sample, nil
	}
	return nil, errors.New("Sample not found")
}

// TODO: Support MultipleObjects retrieval and getting. Useful when using something over 100megs
