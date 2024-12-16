package service

import (
	"github.com/GOAT-prod/goatcontext"
	"github.com/samber/lo"
	"time"
	"warehouse-service/cluster/clientservice"
	"warehouse-service/database/broker"
	"warehouse-service/database/models"
	"warehouse-service/domain"
	"warehouse-service/domain/mappings"
	"warehouse-service/repository"
)

type WareHouse interface {
	GetProducts(ctx goatcontext.Context) ([]domain.Product, error)
	GetDetailedProductsInfo(ctx goatcontext.Context, productId int) (domain.Product, error)
	AddProducts(ctx goatcontext.Context, products []domain.Product) error
	UpdateProducts(ctx goatcontext.Context, products []domain.Product) error
	DeleteProducts(ctx goatcontext.Context, productIds []int) error
	GetMaterialsList(ctx goatcontext.Context) ([]domain.ProductMaterial, error)
	GetProductItemsInfo(ctx goatcontext.Context, ids []int) ([]domain.ProductItemInfo, error)
}

type Impl struct {
	warehouseRepository repository.Warehouse
	clientServiceClient *clientservice.Client
	kafkaProducer       *broker.Producer
}

func NewWarehouseService(warehouseRepository repository.Warehouse, clientServiceClient *clientservice.Client, kafkaProducer *broker.Producer) WareHouse {
	return &Impl{
		warehouseRepository: warehouseRepository,
		clientServiceClient: clientServiceClient,
		kafkaProducer:       kafkaProducer,
	}
}

func (s *Impl) GetProducts(ctx goatcontext.Context) ([]domain.Product, error) {
	dbProducts, err := s.warehouseRepository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	factoryIds := lo.Map(dbProducts, func(item models.Product, _ int) int {
		return item.FactoryId
	})

	shortFactoryInfos, err := s.clientServiceClient.GetClientInfoShort(ctx, factoryIds)
	if err != nil {
		return nil, err
	}

	shortFactoryInfosMap := lo.Associate(shortFactoryInfos, func(item clientservice.ClientInfoShort) (int, clientservice.ClientInfoShort) {
		return item.Id, item
	})

	products := make([]domain.Product, 0, len(dbProducts))
	for _, dbProduct := range dbProducts {
		products = append(products, mappings.ToDomainProduct(dbProduct, shortFactoryInfosMap[dbProduct.FactoryId]))
	}

	return products, nil
}

func (s *Impl) GetDetailedProductsInfo(ctx goatcontext.Context, productId int) (domain.Product, error) {
	dbProduct, err := s.warehouseRepository.GetProductById(ctx, productId)
	if err != nil {
		return domain.Product{}, err
	}

	shortFactoryInfo, err := s.clientServiceClient.GetClientInfoShort(ctx, []int{dbProduct.FactoryId})
	if err != nil {
		return domain.Product{}, err
	}

	return mappings.ToDomainProduct(dbProduct, shortFactoryInfo[0]), nil
}

func (s *Impl) AddProducts(ctx goatcontext.Context, products []domain.Product) error {
	addingProducts := make([]models.Product, 0, len(products))
	for _, product := range products {
		addingProducts = append(addingProducts, mappings.ToDatabaseProduct(product))
	}

	addedProducts, err := s.warehouseRepository.AddProducts(ctx, addingProducts)
	if err != nil {
		return err
	}

	return s.produceRequest(addedProducts)
}

func (s *Impl) UpdateProducts(ctx goatcontext.Context, products []domain.Product) error {
	updatedItems := make([]models.Product, 0, len(products))
	for _, product := range products {
		updatedItems = append(updatedItems, mappings.ToDatabaseProduct(product))
	}

	return s.warehouseRepository.UpdateProducts(ctx, updatedItems)
}

func (s *Impl) DeleteProducts(ctx goatcontext.Context, productIds []int) error {
	return s.warehouseRepository.DeleteProducts(ctx, productIds)
}

func (s *Impl) GetMaterialsList(ctx goatcontext.Context) ([]domain.ProductMaterial, error) {
	dbMaterials, err := s.warehouseRepository.GetAllMaterials(ctx)
	if err != nil {
		return nil, err
	}

	return mappings.ToDomainProductMaterials(dbMaterials), nil
}

func (s *Impl) GetProductItemsInfo(ctx goatcontext.Context, ids []int) ([]domain.ProductItemInfo, error) {
	dbInfos, err := s.warehouseRepository.GetProductItemsInfo(ctx, ids)
	if err != nil {
		return nil, err
	}

	infos := lo.Map(dbInfos, func(item models.ProductItemInfo, _ int) domain.ProductItemInfo {
		return domain.ProductItemInfo{
			Id:    item.Id,
			Name:  item.Name,
			Price: item.Price,
			Color: item.Color,
			Size:  item.Size,
		}
	})

	return infos, nil
}

func (s *Impl) produceRequest(products []models.Product) error {
	for _, product := range products {
		request := broker.Request{
			Status:      1,
			Type:        1,
			UpdatedDate: time.Now(),
			Summary:     "новый продукт на аппрув",
			Item:        product,
		}

		if err := s.kafkaProducer.Produce(request); err != nil {
			return err
		}
	}

	return nil
}
