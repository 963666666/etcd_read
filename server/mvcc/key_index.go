package mvcc

type revision struct {
	main int64 //当事务发生时自动生成的主版本号
	sub  int64 //事务内的子版本号
}

type generation struct {
	ver     int64      //表示当前generation包含的修改次数
	created revision   //generation创建时的版本
	revs    []revision //存储所有的版本信息
}
