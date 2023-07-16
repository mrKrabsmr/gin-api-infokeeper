package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	"github.com/mrKrabsmr/infokeeper/internal/app/dao"
	"github.com/mrKrabsmr/infokeeper/internal/app/dto"
	"github.com/mrKrabsmr/infokeeper/internal/app/models"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"os"
	"strconv"
)

type InfoService struct {
	DAO *dao.InfoDAO
}

func NewInfoService() (*InfoService, error) {
	dao, err := dao.NewInfoDAO()
	if err != nil {
		return nil, err
	}
	return &InfoService{DAO: dao}, nil
}

func (s *InfoService) GetCountInfo(ip string) (int, error) {
	countInfo, err := s.DAO.GetCountValuesByIPAddress(ip)
	if err != nil {
		return -1, err
	}

	return countInfo, nil
}

func (s *InfoService) GetValue(info *dto.GetInfoDTO) (string, error) {
	id, err := uuid.Parse(info.ID)
	if err != nil {
		return "", err
	}

	infoObj, err := s.CheckKeyAndGetObject(id, info.Key)
	if err != nil {
		return "", err
	}

	plaintext, err := s.Decrypt(infoObj.Key[32:], infoObj.Value)
	if err != nil {
		return "", err
	}

	return plaintext, nil
}

func (s *InfoService) Save(info *dto.CreateInfoDTO, ip string) (uuid.UUID, error) {
	key := s.HashKey([]byte(info.Key), nil)

	value, err := s.Encrypt(key[32:], []byte(info.Value))
	if err != nil {
		return uuid.UUID{}, err
	}

	clientService, err := NewClientService()
	if err != nil {
		return uuid.UUID{}, err
	}

	cId, err := clientService.Save(ip)
	if err != nil {
		return uuid.UUID{}, err
	}

	infoObj := &models.Info{
		ID:       uuid.New(),
		Key:      key,
		Value:    value,
		ClientID: cId,
		ReadOnly: info.ReadOnly,
	}

	infoID, err := s.DAO.Create(*infoObj)
	if err != nil {
		return uuid.UUID{}, err
	}

	return infoID, nil
}

func (s *InfoService) PartialUpdate(info *dto.UpdateInfoDTO, ip string) error {
	infoObj, err := s.CheckKeyAndGetObject(info.ID, info.Key)
	if err != nil {
		return err
	}

	ipObj, err := s.DAO.GetIPAddressByKey(infoObj.Key)
	if err != nil {
		return err
	}

	if infoObj.ReadOnly {
		if ip != ipObj {
			return errors.New("this information is read-only")
		}
	}

	value, err := s.Encrypt(infoObj.Key[32:], []byte(info.Value))
	if err != nil {
		return err
	}

	updInfo := &models.Info{
		ID:       infoObj.ID,
		Key:      infoObj.Key,
		Value:    value,
		ReadOnly: info.ReadOnly,
	}

	if err := s.DAO.Update(*updInfo); err != nil {
		return err
	}

	return nil
}

func (s *InfoService) Delete(info *dto.DeleteInfoDTO) error {
	infoObj, err := s.CheckKeyAndGetObject(info.ID, info.Key)
	if err != nil {
		return err
	}

	if err := s.DAO.Delete(infoObj.Key); err != nil {
		return err
	}

	return nil
}

func (s *InfoService) CheckKeyAndGetObject(id uuid.UUID, key string) (models.Info, error) {
	notFoundError := errors.New("invalid id - key pair")

	infoObj, err := s.DAO.GetObjByID(id)
	if err != nil {
		return infoObj, notFoundError
	}

	hash := s.HashKey([]byte(key), infoObj.Key[:32])

	if hex.EncodeToString(hash) != hex.EncodeToString(infoObj.Key) {
		return infoObj, notFoundError
	}

	return infoObj, nil
}

func (s *InfoService) HashKey(key, salt []byte) []byte {
	if len(salt) == 0 {
		salt = make([]byte, 32)
		rand.Read(salt)
	}
	numberOfIterations, _ := strconv.Atoi(os.Getenv("NUMBER_OF_ITERATIONS"))
	hash := pbkdf2.Key(key, salt, numberOfIterations, 32, sha512.New)

	saltAndHash := append(salt, hash...)

	log.Println(len(saltAndHash))

	return saltAndHash
}

func (s *InfoService) Encrypt(key, plaintext []byte) ([]byte, error) {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	rand.Read(nonce)
	ciphertext := gcmInstance.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

func (s *InfoService) Decrypt(key, ciphertext []byte) (string, error) {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return "", err
	}

	nonceSize := gcmInstance.NonceSize()

	plaintext, err := gcmInstance.Open(
		nil, ciphertext[:nonceSize], ciphertext[nonceSize:], nil,
	)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
