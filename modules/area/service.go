package area

import (
	"math"
)

type Service interface {
	InsertArea(param1 int32, param2 int64, area_type string) (Area, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) InsertArea(param1 int32, param2 int64, area_type string) (Area, error) {
	ar := Area{}
	var area int64

	switch area_type {
	case "persegi panjang":
		area = int64(math.Ceil(float64(param1) * float64(param2)))
		ar.AreaValue = area
		ar.AreaType = "persegi panjang"
	case "persegi":
		area = int64(math.Ceil(float64(param1) * float64(param2)))
		ar.AreaValue = area
		ar.AreaType = "persegi"
	case "segitiga":
		area = int64(math.Ceil(0.5 * (float64(param1) * float64(param2))))
		ar.AreaValue = area
		ar.AreaType = "segitiga"
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data"
		// _, err = s.repository.Save(ar)
		// if err != nil {
		// 	return err
		// }
	}
	newArea, err := s.repository.Save(ar)
	if err != nil {
		return newArea, err
	}
	return newArea, nil
	// return nil

}
