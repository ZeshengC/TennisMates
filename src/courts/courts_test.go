package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct{
		request Request
		expect Response
		err error
	}{
		{
			request:Request{},
			expect: Response{
				StatusCode: 200,
				Headers: Headers{ContentType: "application/json"},
				Body: "[{\"id\":1,\"coordinates\":{\"x\":38.93485,\"y\":-77.065859}},{\"id\":2,\"coordinates\":{\"x\":38.94393,\"y\":-77.05041}},{\"id\":3,\"coordinates\":{\"x\":38.936987,\"y\":-77.084742}}]",
				IsBase64Encoded:false,
			},
			err:nil,
		},
	}

	for _,test := range tests {
		response,err := Handler(test.request)
		assert.IsType(t,test.err,err)
		assert.Equal(t,test.expect.Body,response.Body)
	}
}