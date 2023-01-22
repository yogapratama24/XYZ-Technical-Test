package service

import (
	"xyz_test/model"
	"xyz_test/repository"
)

type TransactionService interface {
	CreateTransaction(list model.CreateTransactionRequest) error
	ReadTransaction(userId int) (response []model.ReadTransactionResponse, err error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) *transactionService {
	return &transactionService{repository}
}

func (s *transactionService) CreateTransaction(request model.CreateTransactionRequest) error {
	transactionCreate := model.Transaction{
		NomorKontrak:  request.NomorKontrak,
		Otr:           request.Otr,
		AdminFee:      request.AdminFee,
		JumlahCicilan: request.JumlahCicilan,
		JumlahBunga:   request.JumlahBunga,
		NamaAsset:     request.NamaAsset,
		UserId:        request.UserId,
	}
	if err := s.transactionRepository.CreateTransaction(transactionCreate); err != nil {
		return err
	}

	return nil
}

func (s *transactionService) ReadTransaction(userId int) (response []model.ReadTransactionResponse, err error) {
	data, err := s.transactionRepository.ReadTransaction(userId)
	if err != nil {
		return response, err
	}

	for i := range data {
		var respRow model.ReadTransactionResponse
		respRow.Id = data[i].Id
		respRow.NomorKontrak = data[i].NomorKontrak
		respRow.Otr = data[i].Otr
		respRow.AdminFee = data[i].AdminFee
		respRow.JumlahCicilan = data[i].JumlahCicilan
		respRow.JumlahBunga = data[i].JumlahBunga
		respRow.NamaAsset = data[i].NamaAsset
		respRow.UserId = data[i].UserId
		respRow.User.Id = data[i].User.Id
		respRow.User.FullName = data[i].User.FullName
		response = append(response, respRow)
	}
	return response, nil
}
