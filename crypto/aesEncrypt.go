package crypto

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/expression/function"

)

func init() {
	_ = function.Register(&aesEncryptFn{})
}

type aesEncryptFn struct {
}

// Name returns the name of the function
func (aesEncryptFn) Name() string {
	return "aesEncrypt"
}

// Sig returns the function signature
func (aesEncryptFn) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

// Eval executes the function
func (aesEncryptFn) Eval(params ...interface{}) (interface{}, error) {

	if logger.DebugEnabled() {
		logger.Debugf("Entering function aesEncrypt ()")
	}

	plaintext, err := coerce.ToString(params[0])
	if err != nil {
		return nil, fmt.Errorf("crypto.aesEncrypt function first parameter [%+v] must be string", params[0])
	}

	key, err := coerce.ToString(params[1])
	if err != nil {
		return nil, fmt.Errorf("crypto.aesEncrypt function second parameter [%+v] must be string", params[1])
	}

	ciphertext, err := aesEncrypt([]byte(plaintext), []byte(key))
	if err != nil {
		return nil, err
	}


	if logger.DebugEnabled() {
		logger.Debugf("Exiting function aesEncrypt (eval)")
	}

	return ciphertext, nil
}