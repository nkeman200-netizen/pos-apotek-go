package service

import (
	"apotek-pos-go/internal/modules/inventory/entity"
	"apotek-pos-go/internal/modules/inventory/repository"
	"errors"
)

type UnitService interface {
	CreateUnit(unit *entity.Unit) error
	UpdateUnit(unit *entity.Unit) error
	DeleteUnit(id uint) error
	GetAllUnit() ([]entity.Unit,error)
	GetUnitById(id uint) (entity.Unit,error)
}

type unitServiceImpl struct{
	repo repository.UnitRepository
}

func NewUnitService(repo repository.UnitRepository) UnitService{
	return &unitServiceImpl{repo: repo}
}

func (s *unitServiceImpl) CreateUnit(unit *entity.Unit) error{
	if unit.Name=="" {
		return errors.New("GAGAL: Nama unit tidak boleh kosong")
	}
	if len(unit.Name)>50 {
		return errors.New("GAGAL: Nama unit tidak boleh lebih dari 50 huruf")
	}
	if unit.ShortName=="" {
		return errors.New("GAGAL: Short name tidak boleh kosong")
	}
	if len(unit.ShortName)>20 {
		return errors.New("GAGAL: Short name unit tidak boleh lebih dari 20 huruf")
	}
	
	return s.repo.Create(unit)
}

func (s *unitServiceImpl) UpdateUnit(unit *entity.Unit) error{
	if unit.Name=="" {
		return errors.New("GAGAL: Nama unit tidak boleh kosong")
	}
	if len(unit.Name)>50 {
		return errors.New("GAGAL: Nama unit tidak boleh lebih dari 50 huruf")
	}
	if unit.ShortName=="" {
		return errors.New("GAGAL: Short name tidak boleh kosong")
	}
	if len(unit.ShortName)>20 {
		return errors.New("GAGAL: Short name unit tidak boleh lebih dari 20 huruf")
	}
	
	return s.repo.Update(unit)
}

func (s *unitServiceImpl) DeleteUnit(id uint) error{
	return s.repo.Delete(id)
}

func (s *unitServiceImpl) GetAllUnit() ([]entity.Unit,error){
	return s.repo.FindAll()
}

func(s *unitServiceImpl) GetUnitById(id uint) (entity.Unit,error){
	return s.repo.FindById(id)
}