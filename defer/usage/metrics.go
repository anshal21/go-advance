package usage

func TaskOrchestrator() error {

	err := FunctionA()
	if err != nil {
		log.Error(err)
		metrics.Count("error")
		return err
	}

	err = FunctionB() 
	if err != nil {
		log.Error(err)
		metrics.Count("error")
		return err
	}

	err = FunctionC()
	if err != nil {
		log.Error(err)
		metrics.Count("error")
		return err
	}


	err = Function D()
	if err != nil {
		log.Error(err)
		metrics.Count("error")
		return err
	}

	return nil
}


func TaskOrchestratorWithDefer() (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("panic happened")
		}
		
		if err != nil {
			log.Error(err)
			metrics.Count("error")
		}
	} ()

	err = FunctionA()
	if err != nil {
		return err
	}

	err = FunctionB() 
	if err != nil {
		return err
	}

	err = FunctionC()
	if err != nil {
		return err
	}

	err = Function D()
	if err != nil {
		return err
	}

	return nil
}