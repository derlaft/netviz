package main
import (
	πg "grumpy"
	π_grumpyΓlibΓnumpy "grumpy/lib/numpy"
	π_grumpyΓlibΓsys "grumpy/lib/sys"
	π_os "os"
)
func initModule(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
	ß__main__ := πg.InternStr("__main__")
	ß__name__ := πg.InternStr("__name__")
	ßadd := πg.InternStr("add")
	ßappend := πg.InternStr("append")
	ßarray := πg.InternStr("array")
	ßcrossProd := πg.InternStr("crossProd")
	ßdet := πg.InternStr("det")
	ßdivide := πg.InternStr("divide")
	ßdo := πg.InternStr("do")
	ßdot := πg.InternStr("dot")
	ßfloat := πg.InternStr("float")
	ßinput := πg.InternStr("input")
	ßlen := πg.InternStr("len")
	ßlinalg := πg.InternStr("linalg")
	ßlist := πg.InternStr("list")
	ßmultiply := πg.InternStr("multiply")
	ßnegative := πg.InternStr("negative")
	ßnumpy := πg.InternStr("numpy")
	ßrange := πg.InternStr("range")
	ßreversed := πg.InternStr("reversed")
	ßsplit := πg.InternStr("split")
	ßstdin := πg.InternStr("stdin")
	ßstrip := πg.InternStr("strip")
	ßsys := πg.InternStr("sys")
	var πTemp001 *πg.Object
	_ = πTemp001
	var πTemp002 []*πg.Object
	_ = πTemp002
	var πTemp003 []πg.Param
	_ = πTemp003
	var πTemp004 *πg.Object
	_ = πTemp004
	var πTemp005 *πg.Object
	_ = πTemp005
	var πTemp006 *πg.Object
	_ = πTemp006
	var πTemp007 *πg.Object
	_ = πTemp007
	var πTemp008 bool
	_ = πTemp008
	var πTemp009 *πg.Object
	_ = πTemp009
	var πTemp010 *πg.Object
	_ = πTemp010
	var πTemp011 *πg.Object
	_ = πTemp011
	var πE *πg.BaseException; _ = πE
	for ; πF.State() >= 0; πF.PopCheckpoint() {
		switch πF.State() {
		case 0:
		default: panic("unexpected function state")
		}
		// line 1: import numpy
		πF.SetLineno(1)
		if πTemp002, πE = πg.ImportModule(πF, "numpy", []*πg.Code{π_grumpyΓlibΓnumpy.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßnumpy.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 2: import sys
		πF.SetLineno(2)
		if πTemp002, πE = πg.ImportModule(πF, "sys", []*πg.Code{π_grumpyΓlibΓsys.Code}); πE != nil {
			continue
		}
		πTemp001 = πTemp002[0]
		if πE = πF.Globals().SetItem(πF, ßsys.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 8: def input():
		πF.SetLineno(8)
		πTemp003 = make([]πg.Param, 0)
		πTemp001 = πg.NewFunction(πg.NewCode("input", "../tr2/step1.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µs *πg.Object = πg.UnboundLocal; _ = µs
			var µline *πg.Object = πg.UnboundLocal; _ = µline
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πTemp006 []*πg.Object
			_ = πTemp006
			var πTemp007 *πg.Object
			_ = πTemp007
			var πTemp008 bool
			_ = πTemp008
			var πTemp009 *πg.Object
			_ = πTemp009
			var πTemp010 []πg.Param
			_ = πTemp010
			var πTemp011 *πg.Object
			_ = πTemp011
			return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 5: goto Label5
					default: panic("unexpected function state")
					}
					// line 9: for s in sys.stdin:
					πF.SetLineno(9)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßsys); πE != nil {
						continue
					}
					if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßstdin, nil); πE != nil {
						continue
					}
					if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
						continue
					}
				Label1:
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
							continue
						}
						if !isStop {
							continue
						}
						πE = nil
						πF.RestoreExc(nil, nil)
						goto Label2
					}
					µs = πTemp003
					// line 10: line = s.strip().split()
					πF.SetLineno(10)
					if πE = πg.CheckLocal(πF, µs, "s"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µs, ßstrip, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, πTemp005, ßsplit, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					µline = πTemp005
					πTemp006 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					πTemp006[0] = µline
					if πTemp005, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
						continue
					}
					if πTemp007, πE = πTemp005.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					if πTemp004, πE = πg.LT(πF, πTemp007, πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					if πTemp008, πE = πg.IsTrue(πF, πTemp004); πE != nil {
						continue
					}
					if πTemp008 {
						goto Label3
					}
					goto Label4
					// line 12: if len(line) < 3:
					πF.SetLineno(12)
				Label3:
					// line 13: continue
					πF.SetLineno(13)
					goto Label1
					goto Label4
				Label4:
					// line 14: yield (line[0], numpy.array([float(x) for x in line[1:]]))
					πF.SetLineno(14)
					πTemp005 = πg.NewInt(0).ToObject()
					if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
						continue
					}
					if πTemp007, πE = πg.GetItem(πF, µline, πTemp005); πE != nil {
						continue
					}
					πTemp006 = πF.MakeArgs(1)
					πTemp010 = make([]πg.Param, 0)
					πTemp009 = πg.NewFunction(πg.NewCode("<generator>", "../tr2/step1.py", πTemp010, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µx *πg.Object = πg.UnboundLocal; _ = µx
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 *πg.Object
						_ = πTemp002
						var πTemp003 *πg.Object
						_ = πTemp003
						var πTemp004 []*πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πTemp006 *πg.Object
						_ = πTemp006
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							var πE *πg.BaseException; _ = πE
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 3: goto Label3
								default: panic("unexpected function state")
								}
								// line 14: yield (line[0], numpy.array([float(x) for x in line[1:]]))
								πF.SetLineno(14)
								if πTemp001, πE = πg.SliceType.Call(πF, πg.Args{πg.NewInt(1).ToObject(), πg.None, πg.None}, nil); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µline, "line"); πE != nil {
									continue
								}
								if πTemp002, πE = πg.GetItem(πF, µline, πTemp001); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, πTemp002); πE != nil {
									continue
								}
							Label1:
								if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
										continue
									}
									if !isStop {
										continue
									}
									πE = nil
									πF.RestoreExc(nil, nil)
									goto Label2
								}
								µx = πTemp003
								// line 14: yield (line[0], numpy.array([float(x) for x in line[1:]]))
								πF.SetLineno(14)
								πTemp004 = πF.MakeArgs(1)
								if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
									continue
								}
								πTemp004[0] = µx
								if πTemp005, πE = πg.ResolveGlobal(πF, ßfloat); πE != nil {
									continue
								}
								if πTemp006, πE = πTemp005.Call(πF, πTemp004, nil); πE != nil {
									continue
								}
								πF.FreeArgs(πTemp004)
								πF.PushCheckpoint(3)
								return πTemp006, nil
							Label3:
								πTemp005 = πSent
								goto Label1
								goto Label2
							Label2:
								return nil, nil
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp011, πE = πTemp009.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πg.ListType.Call(πF, πg.Args{πTemp011}, nil); πE != nil {
						continue
					}
					πTemp006[0] = πTemp005
					if πTemp005, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
						continue
					}
					if πTemp011, πE = πg.GetAttr(πF, πTemp005, ßarray, nil); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp011.Call(πF, πTemp006, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp006)
					πTemp004 = πg.NewTuple2(πTemp007, πTemp005).ToObject()
					πF.PushCheckpoint(5)
					return πTemp004, nil
				Label5:
					πTemp005 = πSent
					goto Label1
					goto Label2
				Label2:
					return nil, nil
				}
				return nil, πE
			}).ToObject(), nil
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßinput.ToObject(), πTemp001); πE != nil {
			continue
		}
		// line 16: def crossProd(a,b):
		πF.SetLineno(16)
		πTemp003 = make([]πg.Param, 2)
		πTemp003[0] = πg.Param{Name: "a", Def: nil}
		πTemp003[1] = πg.Param{Name: "b", Def: nil}
		πTemp004 = πg.NewFunction(πg.NewCode("crossProd", "../tr2/step1.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µa *πg.Object = πArgs[0]; _ = µa
			var µb *πg.Object = πArgs[1]; _ = µb
			var µdimension *πg.Object = πg.UnboundLocal; _ = µdimension
			var µc *πg.Object = πg.UnboundLocal; _ = µc
			var µi *πg.Object = πg.UnboundLocal; _ = µi
			var µj *πg.Object = πg.UnboundLocal; _ = µj
			var µk *πg.Object = πg.UnboundLocal; _ = µk
			var πTemp001 []*πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 *πg.Object
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 *πg.Object
			_ = πTemp005
			var πTemp006 *πg.Object
			_ = πTemp006
			var πTemp007 *πg.Object
			_ = πTemp007
			var πTemp008 *πg.Object
			_ = πTemp008
			var πTemp009 bool
			_ = πTemp009
			var πTemp010 *πg.Object
			_ = πTemp010
			var πTemp011 *πg.Object
			_ = πTemp011
			var πTemp012 *πg.Object
			_ = πTemp012
			var πTemp013 *πg.Object
			_ = πTemp013
			var πTemp014 *πg.Object
			_ = πTemp014
			var πTemp015 *πg.Object
			_ = πTemp015
			var πTemp016 *πg.Object
			_ = πTemp016
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 17: dimension = len(a)
				πF.SetLineno(17)
				πTemp001 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp001[0] = µa
				if πTemp002, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				µdimension = πTemp003
				// line 18: c = []
				πF.SetLineno(18)
				πTemp001 = make([]*πg.Object, 0)
				πTemp002 = πg.NewList(πTemp001...).ToObject()
				µc = πTemp002
				// line 19: for i in range(dimension):
				πF.SetLineno(19)
				πTemp001 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µdimension, "dimension"); πE != nil {
					continue
				}
				πTemp001[0] = µdimension
				if πTemp002, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
					continue
				}
				if πTemp003, πE = πTemp002.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				if πTemp002, πE = πg.Iter(πF, πTemp003); πE != nil {
					continue
				}
			Label1:
				if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
					isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
					if exc != nil {
						πE = exc
						continue
					}
					if !isStop {
						continue
					}
					πE = nil
					πF.RestoreExc(nil, nil)
					goto Label2
				}
				µi = πTemp004
				// line 20: c.append(0)
				πF.SetLineno(20)
				πTemp001 = πF.MakeArgs(1)
				πTemp001[0] = πg.NewInt(0).ToObject()
				if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
					continue
				}
				if πTemp005, πE = πg.GetAttr(πF, µc, ßappend, nil); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				// line 21: for j in range(dimension):
				πF.SetLineno(21)
				πTemp001 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µdimension, "dimension"); πE != nil {
					continue
				}
				πTemp001[0] = µdimension
				if πTemp005, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
					continue
				}
				if πTemp006, πE = πTemp005.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				if πTemp005, πE = πg.Iter(πF, πTemp006); πE != nil {
					continue
				}
			Label3:
				if πTemp007, πE = πg.Next(πF, πTemp005); πE != nil {
					isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
					if exc != nil {
						πE = exc
						continue
					}
					if !isStop {
						continue
					}
					πE = nil
					πF.RestoreExc(nil, nil)
					goto Label4
				}
				µj = πTemp007
				if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
					continue
				}
				if πTemp008, πE = πg.NE(πF, µj, µi); πE != nil {
					continue
				}
				if πTemp009, πE = πg.IsTrue(πF, πTemp008); πE != nil {
					continue
				}
				if πTemp009 {
					goto Label5
				}
				goto Label6
				// line 22: if j != i:
				πF.SetLineno(22)
			Label5:
				// line 23: for k in range(dimension):
				πF.SetLineno(23)
				πTemp001 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µdimension, "dimension"); πE != nil {
					continue
				}
				πTemp001[0] = µdimension
				if πTemp008, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
					continue
				}
				if πTemp010, πE = πTemp008.Call(πF, πTemp001, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp001)
				if πTemp008, πE = πg.Iter(πF, πTemp010); πE != nil {
					continue
				}
			Label7:
				if πTemp011, πE = πg.Next(πF, πTemp008); πE != nil {
					isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
					if exc != nil {
						πE = exc
						continue
					}
					if !isStop {
						continue
					}
					πE = nil
					πF.RestoreExc(nil, nil)
					goto Label8
				}
				µk = πTemp011
				if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
					continue
				}
				if πTemp012, πE = πg.NE(πF, µk, µi); πE != nil {
					continue
				}
				if πTemp009, πE = πg.IsTrue(πF, πTemp012); πE != nil {
					continue
				}
				if πTemp009 {
					goto Label9
				}
				goto Label10
				// line 24: if k != i:
				πF.SetLineno(24)
			Label9:
				if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
					continue
				}
				if πTemp012, πE = πg.GT(πF, µk, µj); πE != nil {
					continue
				}
				if πTemp009, πE = πg.IsTrue(πF, πTemp012); πE != nil {
					continue
				}
				if πTemp009 {
					goto Label11
				}
				if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
					continue
				}
				if πTemp012, πE = πg.LT(πF, µk, µj); πE != nil {
					continue
				}
				if πTemp009, πE = πg.IsTrue(πF, πTemp012); πE != nil {
					continue
				}
				if πTemp009 {
					goto Label12
				}
				goto Label13
				// line 25: if k > j:
				πF.SetLineno(25)
			Label11:
				// line 26: c[i] += a[j]*b[k]
				πF.SetLineno(26)
				if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
					continue
				}
				πTemp012 = µi
				if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
					continue
				}
				if πTemp013, πE = πg.GetItem(πF, µc, πTemp012); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
					continue
				}
				πTemp014 = µj
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				if πTemp015, πE = πg.GetItem(πF, µa, πTemp014); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
					continue
				}
				πTemp014 = µk
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				if πTemp016, πE = πg.GetItem(πF, µb, πTemp014); πE != nil {
					continue
				}
				if πTemp012, πE = πg.Mul(πF, πTemp015, πTemp016); πE != nil {
					continue
				}
				if πTemp014, πE = πg.IAdd(πF, πTemp013, πTemp012); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
					continue
				}
				πTemp015 = µi
				if πE = πg.SetItem(πF, µc, πTemp015, πTemp014); πE != nil {
					continue
				}
				goto Label13
				// line 27: elif k < j:
				πF.SetLineno(27)
			Label12:
				// line 28: c[i] -= a[j]*b[k]
				πF.SetLineno(28)
				if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
					continue
				}
				πTemp012 = µi
				if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
					continue
				}
				if πTemp013, πE = πg.GetItem(πF, µc, πTemp012); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µj, "j"); πE != nil {
					continue
				}
				πTemp014 = µj
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				if πTemp015, πE = πg.GetItem(πF, µa, πTemp014); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µk, "k"); πE != nil {
					continue
				}
				πTemp014 = µk
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				if πTemp016, πE = πg.GetItem(πF, µb, πTemp014); πE != nil {
					continue
				}
				if πTemp012, πE = πg.Mul(πF, πTemp015, πTemp016); πE != nil {
					continue
				}
				if πTemp014, πE = πg.ISub(πF, πTemp013, πTemp012); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
					continue
				}
				if πE = πg.CheckLocal(πF, µi, "i"); πE != nil {
					continue
				}
				πTemp015 = µi
				if πE = πg.SetItem(πF, µc, πTemp015, πTemp014); πE != nil {
					continue
				}
				goto Label13
			Label13:
				goto Label10
			Label10:
				goto Label7
				goto Label8
			Label8:
				goto Label6
			Label6:
				goto Label3
				goto Label4
			Label4:
				goto Label1
				goto Label2
			Label2:
				// line 29: return c
				πF.SetLineno(29)
				if πE = πg.CheckLocal(πF, µc, "c"); πE != nil {
					continue
				}
				return µc, nil
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßcrossProd.ToObject(), πTemp004); πE != nil {
			continue
		}
		// line 31: def do(dot):
		πF.SetLineno(31)
		πTemp003 = make([]πg.Param, 1)
		πTemp003[0] = πg.Param{Name: "dot", Def: nil}
		πTemp005 = πg.NewFunction(πg.NewCode("do", "../tr2/step1.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µdot *πg.Object = πArgs[0]; _ = µdot
			var µname *πg.Object = πg.UnboundLocal; _ = µname
			var µx *πg.Object = πg.UnboundLocal; _ = µx
			var µbasis *πg.Object = πg.UnboundLocal; _ = µbasis
			var µa *πg.Object = πg.UnboundLocal; _ = µa
			var µb *πg.Object = πg.UnboundLocal; _ = µb
			var µnorm *πg.Object = πg.UnboundLocal; _ = µnorm
			var µmn1 *πg.Object = πg.UnboundLocal; _ = µmn1
			var µres *πg.Object = πg.UnboundLocal; _ = µres
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 []*πg.Object
			_ = πTemp003
			var πTemp004 []*πg.Object
			_ = πTemp004
			var πTemp005 []*πg.Object
			_ = πTemp005
			var πTemp006 []*πg.Object
			_ = πTemp006
			var πTemp007 []*πg.Object
			_ = πTemp007
			var πTemp008 []*πg.Object
			_ = πTemp008
			var πTemp009 *πg.Object
			_ = πTemp009
			var πTemp010 *πg.Object
			_ = πTemp010
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				// line 33: name = dot[0]
				πF.SetLineno(33)
				πTemp001 = πg.NewInt(0).ToObject()
				if πE = πg.CheckLocal(πF, µdot, "dot"); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetItem(πF, µdot, πTemp001); πE != nil {
					continue
				}
				µname = πTemp002
				// line 34: x = dot[1]
				πF.SetLineno(34)
				πTemp001 = πg.NewInt(1).ToObject()
				if πE = πg.CheckLocal(πF, µdot, "dot"); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetItem(πF, µdot, πTemp001); πE != nil {
					continue
				}
				µx = πTemp002
				// line 36: basis = range(len(x))
				πF.SetLineno(36)
				πTemp003 = πF.MakeArgs(1)
				πTemp004 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
					continue
				}
				πTemp004[0] = µx
				if πTemp001, πE = πg.ResolveGlobal(πF, ßlen); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003[0] = πTemp002
				if πTemp001, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				µbasis = πTemp002
				// line 38: a = numpy.array(basis)
				πF.SetLineno(38)
				πTemp003 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µbasis, "basis"); πE != nil {
					continue
				}
				πTemp003[0] = µbasis
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßarray, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				µa = πTemp001
				// line 39: b = numpy.array(list(reversed(basis)))
				πF.SetLineno(39)
				πTemp003 = πF.MakeArgs(1)
				πTemp004 = πF.MakeArgs(1)
				πTemp005 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µbasis, "basis"); πE != nil {
					continue
				}
				πTemp005[0] = µbasis
				if πTemp001, πE = πg.ResolveGlobal(πF, ßreversed); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp005)
				πTemp004[0] = πTemp002
				if πTemp001, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003[0] = πTemp002
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßarray, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				µb = πTemp001
				// line 41: norm = numpy.linalg.det(numpy.array([[numpy.dot(a, a), numpy.dot(a, b)], [numpy.dot(b, a), numpy.dot(b, b)]]))
				πF.SetLineno(41)
				πTemp003 = πF.MakeArgs(1)
				πTemp004 = πF.MakeArgs(1)
				πTemp005 = make([]*πg.Object, 2)
				πTemp006 = make([]*πg.Object, 2)
				πTemp007 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp007[0] = µa
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp007[1] = µa
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdot, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp007)
				πTemp006[0] = πTemp001
				πTemp007 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp007[0] = µa
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp007[1] = µb
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdot, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp007)
				πTemp006[1] = πTemp001
				πTemp001 = πg.NewList(πTemp006...).ToObject()
				πTemp005[0] = πTemp001
				πTemp006 = make([]*πg.Object, 2)
				πTemp007 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp007[0] = µb
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp007[1] = µa
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdot, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp007)
				πTemp006[0] = πTemp001
				πTemp007 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp007[0] = µb
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp007[1] = µb
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdot, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp007)
				πTemp006[1] = πTemp001
				πTemp001 = πg.NewList(πTemp006...).ToObject()
				πTemp005[1] = πTemp001
				πTemp001 = πg.NewList(πTemp005...).ToObject()
				πTemp004[0] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßarray, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003[0] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßlinalg, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πg.GetAttr(πF, πTemp002, ßdet, nil); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				µnorm = πTemp002
				// line 43: mn1 = numpy.add(
				πF.SetLineno(43)
				πTemp003 = πF.MakeArgs(2)
				πTemp004 = πF.MakeArgs(2)
				πTemp005 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp005[0] = µa
				πTemp006 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp006[0] = µa
				πTemp007 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp007[0] = µb
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp007[1] = µb
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdot, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp007)
				πTemp006[1] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdot, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp006)
				πTemp005[1] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßcrossProd); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp005, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp005)
				πTemp004[0] = πTemp002
				πTemp005 = πF.MakeArgs(2)
				πTemp005[0] = πg.NewInt(2).ToObject()
				πTemp006 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp006[0] = µa
				πTemp007 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp007[0] = µb
				πTemp008 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp008[0] = µa
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp008[1] = µb
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdot, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp008, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp008)
				πTemp007[1] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdot, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp007, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp007)
				πTemp006[1] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßcrossProd); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp006, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp006)
				πTemp005[1] = πTemp002
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßmultiply, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp005)
				πTemp004[1] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßadd, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003[0] = πTemp001
				πTemp004 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp004[0] = µb
				πTemp005 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
					continue
				}
				πTemp005[0] = µb
				πTemp006 = πF.MakeArgs(2)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp006[0] = µa
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp006[1] = µa
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdot, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp006, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp006)
				πTemp005[1] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdot, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp005)
				πTemp004[1] = πTemp001
				if πTemp001, πE = πg.ResolveGlobal(πF, ßcrossProd); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003[1] = πTemp002
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßadd, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				µmn1 = πTemp001
				// line 50: res = numpy.divide(crossProd(numpy.negative(x), mn1), norm)
				πF.SetLineno(50)
				πTemp003 = πF.MakeArgs(2)
				πTemp004 = πF.MakeArgs(2)
				πTemp005 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µx, "x"); πE != nil {
					continue
				}
				πTemp005[0] = µx
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßnegative, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp005, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp005)
				πTemp004[0] = πTemp001
				if πE = πg.CheckLocal(πF, µmn1, "mn1"); πE != nil {
					continue
				}
				πTemp004[1] = µmn1
				if πTemp001, πE = πg.ResolveGlobal(πF, ßcrossProd); πE != nil {
					continue
				}
				if πTemp002, πE = πTemp001.Call(πF, πTemp004, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp004)
				πTemp003[0] = πTemp002
				if πE = πg.CheckLocal(πF, µnorm, "norm"); πE != nil {
					continue
				}
				πTemp003[1] = µnorm
				if πTemp001, πE = πg.ResolveGlobal(πF, ßnumpy); πE != nil {
					continue
				}
				if πTemp002, πE = πg.GetAttr(πF, πTemp001, ßdivide, nil); πE != nil {
					continue
				}
				if πTemp001, πE = πTemp002.Call(πF, πTemp003, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp003)
				µres = πTemp001
				// line 52: print (name, res[0], res[1])
				πF.SetLineno(52)
				πTemp003 = make([]*πg.Object, 1)
				if πE = πg.CheckLocal(πF, µname, "name"); πE != nil {
					continue
				}
				πTemp002 = πg.NewInt(0).ToObject()
				if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
					continue
				}
				if πTemp009, πE = πg.GetItem(πF, µres, πTemp002); πE != nil {
					continue
				}
				πTemp002 = πg.NewInt(1).ToObject()
				if πE = πg.CheckLocal(πF, µres, "res"); πE != nil {
					continue
				}
				if πTemp010, πE = πg.GetItem(πF, µres, πTemp002); πE != nil {
					continue
				}
				πTemp001 = πg.NewTuple3(µname, πTemp009, πTemp010).ToObject()
				πTemp003[0] = πTemp001
				if πE = πg.Print(πF, πTemp003, true); πE != nil {
					continue
				}
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßdo.ToObject(), πTemp005); πE != nil {
			continue
		}
		if πTemp007, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
			continue
		}
		if πTemp006, πE = πg.Eq(πF, πTemp007, ß__main__.ToObject()); πE != nil {
			continue
		}
		if πTemp008, πE = πg.IsTrue(πF, πTemp006); πE != nil {
			continue
		}
		if πTemp008 {
			goto Label1
		}
		goto Label2
		// line 54: if __name__ == "__main__":
		πF.SetLineno(54)
	Label1:
		// line 55: for dot in input():
		πF.SetLineno(55)
		if πTemp006, πE = πg.ResolveGlobal(πF, ßinput); πE != nil {
			continue
		}
		if πTemp007, πE = πTemp006.Call(πF, nil, nil); πE != nil {
			continue
		}
		if πTemp006, πE = πg.Iter(πF, πTemp007); πE != nil {
			continue
		}
	Label3:
		if πTemp009, πE = πg.Next(πF, πTemp006); πE != nil {
			isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
			if exc != nil {
				πE = exc
				continue
			}
			if !isStop {
				continue
			}
			πE = nil
			πF.RestoreExc(nil, nil)
			goto Label4
		}
		if πE = πF.Globals().SetItem(πF, ßdot.ToObject(), πTemp009); πE != nil {
			continue
		}
		// line 56: do(dot)
		πF.SetLineno(56)
		πTemp002 = πF.MakeArgs(1)
		if πTemp010, πE = πg.ResolveGlobal(πF, ßdot); πE != nil {
			continue
		}
		πTemp002[0] = πTemp010
		if πTemp010, πE = πg.ResolveGlobal(πF, ßdo); πE != nil {
			continue
		}
		if πTemp011, πE = πTemp010.Call(πF, πTemp002, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp002)
		goto Label3
		goto Label4
	Label4:
		goto Label2
	Label2:
		return nil, nil
	}
	return nil, πE
}
var Code *πg.Code
func main() {
	Code = πg.NewCode("<module>", "../tr2/step1.py", nil, 0, initModule)
	π_os.Exit(πg.RunMain(Code))
}
