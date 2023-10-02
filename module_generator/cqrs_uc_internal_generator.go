package modulegenerator

import (
	"os"

	"github.com/Muruyung/go-utilities/logger"
)

func internalCqrsUcGenerator(dto dtoModule, cqrsType string, isAll, isOnly bool) error {
	if !isAll && !isOnly {
		return nil
	}

	dto.path += "usecase" + dto.sep
	var (
		err         error
		pathQuery   = "query" + dto.sep + dto.name
		pathCommand = "command" + dto.sep + dto.name
		_, okGet    = dto.methods["get"]
		_, okList   = dto.methods["getList"]
		_, okCreate = dto.methods["create"]
		_, okUpdate = dto.methods["update"]
		_, okDelete = dto.methods["delete"]
		isQuery     = (okGet || okList || cqrsType == "query")
		isCommand   = (okCreate || okUpdate || okDelete || cqrsType == "command")
	)

	if isQuery {
		if _, err = os.Stat(dto.path + pathQuery); os.IsNotExist(err) {
			err = os.MkdirAll(dto.path+pathQuery, 0777)
			if err != nil {
				return err
			} else {
				logger.Logger.Info("internal directory usecase query created")
			}
		}

		if _, err = os.Stat(dto.path + pathQuery + "/init.go"); os.IsNotExist(err) {
			err = generateInitUc(dto.name, dto.path+pathQuery, dto.services, dto.fields)
			if err != nil {
				return err
			} else {
				logger.Logger.Info("internal init usecase query created")
			}
		}
	}

	if isCommand {
		if _, err = os.Stat(dto.path + pathCommand); os.IsNotExist(err) {
			err = os.MkdirAll(dto.path+pathCommand, 0777)
			if err != nil {
				return err
			} else {
				logger.Logger.Info("internal directory usecase command created")
			}
		}

		if _, err = os.Stat(dto.path + pathCommand + "/init.go"); os.IsNotExist(err) {
			err = generateInitUc(dto.name, dto.path+pathCommand, dto.services, dto.fields)
			if err != nil {
				return err
			} else {
				logger.Logger.Info("internal init usecase command created")
			}
		}
	}

	if _, ok := dto.methods["get"]; ok {
		err = generateGetUc(dto.name, dto.path+pathQuery, dto.services, dto.fields["id"])
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal get usecase created")
		}
	}

	if _, ok := dto.methods["getList"]; ok {
		err = generateGetListUc(dto.name, dto.path+pathQuery, dto.services)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal get list usecase created")
		}
	}

	if _, ok := dto.methods["create"]; ok {
		err = generateCreateUc(dto.name, dto.path+pathCommand, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal create usecase created")
		}
	}

	if _, ok := dto.methods["update"]; ok {
		err = generateUpdateUc(dto.name, dto.path+pathCommand, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal update usecase created")
		}
	}

	if _, ok := dto.methods["delete"]; ok {
		err = generateDeleteUc(dto.name, dto.path+pathCommand, dto.services, dto.fields)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal update usecase created")
		}
	}

	if _, ok := dto.methods["custom"]; ok {
		dto.path += cqrsType + dto.sep + dto.name
		err = generateCustomUc(dto)
		if err != nil {
			return err
		} else {
			logger.Logger.Info("internal custom usecase created")
		}
	}

	return nil
}
