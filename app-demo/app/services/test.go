package services

type MathService struct {
	
}

type Args struct {
	Arg1, Arg2 int
}


func (that *MathService) Add(args Args, reply *int) error {
	*reply = args.Arg1 + args.Arg2
	return nil
}