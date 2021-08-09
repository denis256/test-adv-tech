package db

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"test-adv-tech/model"
)

func ReadData() model.AppResponse {
	// mock data generation

	var response = model.AppResponse{
		Id:      strconv.Itoa(rand.Int()),
		PayLoad: fmt.Sprintf("%x", md5.Sum([]byte{byte(rand.Int())})),
	}

	return response

}
