package repository

import (
	"apotek-pos-go/internal/modules/inventory/entity"

	"gorm.io/gorm"
)

type UnitRepository interface {
	Create(unit *entity.Unit) error
	Update(unit *entity.Unit) error
	Delete(id uint) error
	FindAll() ([]entity.Unit,error)
	FindById(id uint) (entity.Unit,error)
}

type unitRepositoryImp struct{
	db *gorm.DB
}

func NewUnitRepository(db *gorm.DB) UnitRepository{
	return &unitRepositoryImp{db: db}
}

func (r *unitRepositoryImp) Create(unit *entity.Unit) error{
	err:=r.db.Create(unit).Error
	return err
}

func (r *unitRepositoryImp) Update(unit *entity.Unit) error{
	err:=r.db.Save(unit).Error //pake save bukan update
	return err
}

func (r *unitRepositoryImp) Delete(id uint) error{
	err:=r.db.Delete(&entity.Unit{},id).Error
	return  err
}

func (r *unitRepositoryImp) FindAll() ([]entity.Unit,error){
	var unit  []entity.Unit
	err:=r.db.Find(&unit).Error
	return unit,err
}

func (r *unitRepositoryImp) FindById(id uint) (entity.Unit,error){
	var unit entity.Unit
	err:=r.db.First(&unit,id).Error
	return unit,err
}