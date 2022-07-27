package iot

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	log "github.com/sirupsen/logrus"
	"strconv"
	"stressClient/config"
)

type CommonIotDevice struct {
	id   int
	name string

	baseNumber int
	step       int

	formulas config.Formulas
}

func (c *CommonIotDevice) GetId() int {
	return c.id
}

func (c *CommonIotDevice) GetName() string {
	return c.name
}

func (c *CommonIotDevice) Init(config config.IotConfig) {
	c.baseNumber = config.BaseNumber
	c.step = config.Step
	c.formulas = config.Formulas
	c.id = config.ID
	c.name = config.Name
}

func (c *CommonIotDevice) GetCurrentStateFunc1() (string, error) {
	log.Infoln("GetCurrentStateFunc1")
	return c.processFormula(c.formulas.Formula1)
}

func (c *CommonIotDevice) GetCurrentStateFunc2() (string, error) {
	log.Infoln("GetCurrentStateFunc2")
	return c.processFormula(c.formulas.Formula2)
}

func (c *CommonIotDevice) GetCurrentStateFunc3() (string, error) {
	log.Infoln("GetCurrentStateFunc3")
	return c.processFormula(c.formulas.Formula3)
}

func (c *CommonIotDevice) GetCurrentStateFunc4() (string, error) {
	log.Infoln("GetCurrentStateFunc4")
	return c.processFormula(c.formulas.Formula4)
}

func (c *CommonIotDevice) GetCurrentStateFunc5() (string, error) {
	log.Infoln("GetCurrentStateFunc5")
	return c.processFormula(c.formulas.Formula5)
}

func (c *CommonIotDevice) processFormula(formula string) (string, error) {
	if formula == "" {
		return "", errors.New("empty formula")
	}

	expression, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		return "", fmt.Errorf("govaluate err:%v", err)
	}

	c.baseNumber = c.baseNumber + c.step
	parameters := make(map[string]interface{}, 8)
	parameters["base_number"] = c.baseNumber

	result, err := expression.Evaluate(parameters)
	if err != nil {
		return "", fmt.Errorf("expressioan evaluate err:%v", err)
	}

	resultFloat, ok := result.(float64)
	if !ok {
		return "", fmt.Errorf("processFormula: formula returned not float")
	}

	resultStr := strconv.FormatFloat(resultFloat, 'f', -1, 64)
	return resultStr, nil
}
