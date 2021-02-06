package flow

type Act struct{
	Id int;
	Name int;
	InData Data;
	OutData Data;
}


type ActFunc map[int]func(inData Data)(outData Data,err error);

var ActMap = ActFunc{

}

func (act *Act)Run()(outData Data,err error){	
	return ActMap[act.Id](act.InData)
}