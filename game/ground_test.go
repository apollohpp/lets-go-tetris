package game

import (
	"github.com/go-test/deep"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"lets-go-tetris/interfaces/renderer"
	"lets-go-tetris/render"
)

var _ = Describe("ground reset 테스트", func() {
	type testData struct {
		x, y		int32
		expected 	[]cell
	}

	DescribeTable("테스트 케이스", func(d testData) {
		g := ground{x: d.x, y: d.y}
		g.reset()
		actual := g.cells
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("2x3", testData{2, 3, []cell{
			false,	false,
			false,	false,
			false,	false,
		}}),
		Entry("4x3", testData{4, 3, []cell{
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
	)
})

var _ = Describe("ground merge 테스트", func() {
	type testData struct {
		s			Shape
		x, y		int32
		expected	[]cell
	}

	g := ground{x: 4, y:10}
	g.reset()

	DescribeTable("테스트 케이스", func(d testData) {
		m := mino{x: d.x, y:d.y}
		m.init(shapes[d.s])
		g.merge(&m)

		actual := g.cells
		diff := deep.Equal(actual, d.expected)
		Expect(diff).Should(BeNil())
	},
		Entry("L", testData{L, 0,0, []cell{
			false,	false,	true,	false,
			true,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("I", testData{I, 2,3, []cell{
			false,	false,	true,	false,
			true,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	true,	true,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
		Entry("O", testData{O, 3,5, []cell{
			false,	false,	true,	false,
			true,	true,	true,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	true,	true,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
			false,	false,	false,	false,
		}}),
	)
})

var _ = Describe("ground collide 테스트", func() {
	It("정상적인 상황에서는 충돌하지 않는다.", func() {
		g := ground{x: 4, y: 10}
		g.reset()

		m := newMino(S)
		g.merge(m)

		m = newMino(L)
		m.y = 4

		actual := g.collide(m)

		expected := false

		Expect(actual).Should(Equal(expected))
	})

	It("이미 ground에 머지되어있는 블럭과 잘 충돌한다.", func() {
		g := ground{x: 4, y: 10}
		g.reset()

		m := newMino(S)
		g.merge(m)

		m = newMino(L)
		actual := g.collide(m)

		expected := true

		Expect(actual).Should(Equal(expected))
	})

	It("ground 밖으로 나갔는지 체크한다.", func() {
		g := ground{x: 4, y: 10}
		g.reset()

		m := newMino(I)
		m.x = 100
		m.y = 100

		actual := g.collide(m)

		expected := true

		Expect(actual).Should(Equal(expected))
	})
})

var _ = XDescribe("ground step 테스트", func() {
})

var _ = XDescribe("mino wall kick 테스트", func() {
})

var _ = XDescribe("ground tetris check 테스트", func() {
})

var _ = Describe("ground.RenderInfo() 함수가", func() {
	It("렌더링 정보를 제대로 반환한다.", func() {
		g := ground{
			x: 2,
			y: 3,
			cells: []cell{ true, true, false, true, true, true},
			colors: []uint32{123, 123, 123, 123, 123, 123},
		}

		expected := []renderer.Info{
			&render.InfoImpl{PosX: 0, PosY: 0, Color: 123},
			&render.InfoImpl{PosX: 1, PosY: 0, Color: 123},
			&render.InfoImpl{PosX: 0, PosY: 1, Color: 0xff353535},
			&render.InfoImpl{PosX: 1, PosY: 1, Color: 123},
			&render.InfoImpl{PosX: 0, PosY: 2, Color: 123},
			&render.InfoImpl{PosX: 1, PosY: 2, Color: 123},
		}
		actual := g.RenderInfo()
		Expect(actual).Should(Equal(expected))
	})
})
