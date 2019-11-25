package tree

import "testing"

func TestIsSameTree(t *testing.T) {
	var p, q *TreeNode
	var h, r bool

	p = &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}
	q = &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}
	h = true
	r = isSameTree(p, q)
	t.Logf("%t h=%t r=%t", h == r, h, r)

	p = &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}
	q = &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		nil,
	}
	h = false
	r = isSameTree(p, q)
	t.Logf("%t h=%t r=%t", h == r, h, r)

	p = &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			nil,
			nil,
		},
	}
	q = &TreeNode{
		1,
		nil,
		&TreeNode{
			3,
			nil,
			nil,
		},
	}
	h = false
	r = isSameTree(p, q)
	t.Logf("%t h=%t r=%t", h == r, h, r)

	p = &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		nil,
	}
	q = &TreeNode{
		1,
		nil,
		&TreeNode{
			2,
			nil,
			nil,
		},
	}
	h = false
	r = isSameTree(p, q)
	t.Logf("%t h=%t r=%t", h == r, h, r)

	p = &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			1,
			nil,
			nil,
		},
	}
	q = &TreeNode{
		1,
		&TreeNode{
			1,
			nil,
			nil,
		},
		&TreeNode{
			2,
			nil,
			nil,
		},
	}
	h = false
	r = isSameTree(p, q)
	t.Logf("%t h=%t r=%t", h == r, h, r)

	//[4,5,null,1,null,2,null,3]
	//[5,4,null,1,null,2,null,3]

	p = &TreeNode{
		4,
		&TreeNode{
			5,
			&TreeNode{
				1,
				&TreeNode{
					2,
					nil,
					&TreeNode{
						3,
						nil,
						nil,
					},
				},
				nil,
			},
			nil,
		},
		nil,
	}
	q = &TreeNode{
		5,
		&TreeNode{
			4,
			&TreeNode{
				1,
				&TreeNode{
					2,
					nil,
					&TreeNode{
						3,
						nil,
						nil,
					},
				},
				nil,
			},
			nil,
		},
		nil,
	}
	h = false
	r = isSameTree(p, q)
	t.Logf("%t h=%t r=%t", h == r, h, r)
}
