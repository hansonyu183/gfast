package flow

type Work struct{
	Id int;
	Name int;
	InData Data;
	Acts []Act
	OutData Data;
}

func (wk *Work)Run()(outData Data,err error){
	for i,v :=range wk.Acts{
		if i==0{
			v.InData=wk.InData
			outData,err=v.Run()
		}
		outData,err=v.Run()
	}
	return

}

