package region

import (
	"fmt"

	"github.com/sobo-gen3v1/soboweb-tegola-cr/container/singlelist/point/list"
	"github.com/sobo-gen3v1/soboweb-tegola-cr/maths"
)

func (r *Region) GoString() string {
	str := fmt.Sprintf("   Region:(%v)", r.winding)
	r.ForEachIdx(func(idx int, p list.ElementerPointer) bool {
		str += fmt.Sprintf("[%v](%#v)", idx, p)
		return true
	})
	return str
}

func (r *Region) DebugStringAugmented(augmenter func(idx int, e maths.Pt) string) string {
	str := fmt.Sprintf("  Region:(%v)", r.winding)
	r.ForEachIdx(func(idx int, pt list.ElementerPointer) bool {
		str += augmenter(idx, pt.Point())
		return true
	})
	return str
}
