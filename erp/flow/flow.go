package flow

type Data struct{
	id int;
}
// Flow 工作流
type Flow struct{
	Id int;
	Name int;
	InData Data;
	Works []Work;
	OutData Data;
} 



func StartFlow(flowId int,inData Data)(outData Data,err error){
	fl:=&Flow{
		Id: flowId,
		InData: inData,
	}
	fl.Load()
	fl.Run()
	return
}

func (fl *Flow)Load(){
	
}

func (fl *Flow)Run()(outData Data,err error){
	for i,v :=range fl.Works{
		if i==0{
			v.InData=fl.InData
			outData,err=v.Run()
		}
		outData,err=v.Run()
	}
	return
}

