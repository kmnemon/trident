package analysis

// func DiffObjects(base, target *object.Project) *object.Project {
// 	if reflect.DeepEqual(base, target) {
// 		log.Println("deep equal")
// 		return nil
// 	}

// 	dp := object.NewProject("diff")

// 	for tname, tpack := range target.GetPackages() {
// 		_, ok := base.GetPackages()[tname]
// 		if !ok {
// 			dp.GetPackages()[tname] = tpack
// 		} else {
// 			for tname1, tclass := range tpack.GetClasses() {
// 				basePack := base.GetPackages()[tname]
// 				_, ok = basePack.GetClasses()[tname1]
// 				if !ok {

// 				}
// 			}
// 		}

// 	}

// 	return dp
// }
