package modulegenerator

import "errors"

func validate(dto dtoModule) (err error) {
	if dto.services == "" {
		return errors.New("flag service cannot be empty")
	}

	if dto.name == "" {
		return errors.New("flag name cannot be empty")
	}

	if _, ok := dto.methods["custom"]; ok {
		if dto.methodName == "" {
			return errors.New("flag custom-method cannot be empty if you're using custom methods")
		}
	}

	return
}