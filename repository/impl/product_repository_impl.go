package impl

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gin-clean-architecture/entity"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/RizkiMufrizal/gin-clean-architecture/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewProductRepositoryImpl(DB *sqlx.DB) repository.ProductRepository {
	return &productRepositoryImpl{DB: DB}
}

type productRepositoryImpl struct {
	*sqlx.DB
}

func (repository *productRepositoryImpl) Insert(ctx context.Context, product entity.Product) entity.Product {
	product.Id = uuid.New()

	//begin transaction
	tx, err := repository.DB.Beginx()
	exception.PanicLogging(err)

	_, err = tx.NamedExecContext(ctx, "INSERT INTO tb_product (product_id, name, price, quantity) VALUES(:product_id, :name, :price, :quantity)", product)
	if err != nil {
		err := tx.Rollback()
		exception.PanicLogging(err)
	}
	//commit
	err = tx.Commit()
	exception.PanicLogging(err)

	return product
}

func (repository *productRepositoryImpl) Update(ctx context.Context, product entity.Product) entity.Product {
	//begin transaction
	tx, err := repository.DB.Beginx()
	exception.PanicLogging(err)

	_, err = tx.NamedExecContext(ctx, "UPDATE tb_product SET name=:name, price=:price, quantity=:quantity WHERE product_id=:product_id", product)
	if err != nil {
		err := tx.Rollback()
		exception.PanicLogging(err)
	}
	//commit
	err = tx.Commit()
	exception.PanicLogging(err)

	return product
}

func (repository *productRepositoryImpl) Delete(ctx context.Context, product entity.Product) {
	//begin transaction
	tx, err := repository.DB.Beginx()
	exception.PanicLogging(err)

	_, err = tx.NamedExecContext(ctx, "DELETE FROM tb_product WHERE product_id=:product_id", product)
	if err != nil {
		err := tx.Rollback()
		exception.PanicLogging(err)
	}
	//commit
	err = tx.Commit()
	exception.PanicLogging(err)
}

func (repository *productRepositoryImpl) FindById(ctx context.Context, id string) (entity.Product, error) {
	var product entity.Product
	err := repository.DB.GetContext(ctx, &product, "SELECT * FROM tb_product WHERE product_id = $1", id)
	if err != nil {
		return entity.Product{}, errors.New("product Not Found")
	}
	return product, nil
}

func (repository *productRepositoryImpl) FindAl(ctx context.Context) []entity.Product {
	var products []entity.Product
	err := repository.DB.SelectContext(ctx, &products, "SELECT * FROM tb_product")
	if err != nil {
		exception.PanicLogging(err)
	}
	return products
}
