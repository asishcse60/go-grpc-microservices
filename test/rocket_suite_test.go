//+acceptance
package test

import (
	"context"
	rkt "github.com/asishcse60/go-grpc-microservices/grpc-protos/rocket/v1"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestRocket() {
	client := GetClient()

	u := (uuid.NewV4()).String()

	s.T().Run("add a new rocket", func(t *testing.T) {
		resp, err := client.AddRocket(context.Background(), &rkt.AddRocketRequest{
			Rocket: &rkt.Rocket{
				Id: u,
				Name: "V1",
				Type: "Falcon Heavy",
			},
		})
		if err != nil{
			assert.Fail(t, "Failed add rocket")
		}
		//assert.NoError(s.T(), err)
		assert.Equal(s.T(),u, resp.Rocket.Id)
	})

	s.T().Run("get a rocket by id", func(t *testing.T) {
		resp,err := client.GetRocket(context.Background(),&rkt.GetRocketRequest{
			Id: u,
		})
		//assert.NoError(s.T(), err)
		if err != nil{
			assert.Fail(t, "Failed add rocket")
		}
		assert.Equal(s.T(), u, resp.Rocket.Id)
	})
	s.T().Run("delete a rocket successfully", func(t *testing.T) {
		resp,err := client.DeleteRocket(context.Background(), &rkt.DeleteRocketRequest{
			Rocket: &rkt.Rocket{
				Id: u,
				Type: "Falcon Heavy",
				Name: "V1",
			},
		})
		//assert.NoError(s.T(), err)
		if err != nil{
			assert.Fail(t, "Failed delete rocket")
		}
		assert.Equal(s.T(), "Successfully delete rocket", resp.Status)
	})
}

func TestRocketService(t *testing.T){
	suite.Run(t, new(RocketTestSuite))
}